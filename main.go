package main

import (
	"errors"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/go-chef/kasoku/kasoku"
	"github.com/golang/groupcache"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

var (
	me         = flag.String("me", "127.0.0.1:710", "The bind address for the groupcache endpoint")
	peerlist   = flag.String("peerlist", "", "The remote peers to attempt to form Voltron and activate incredible speed power")
	mountpoint = flag.String("mountpoint", "/kasoku", "Where to mount the FUSE groupcache storage")
	groupName  = flag.String("groupName", "kasoku", "The name of the groupcache cluster")
	cacheLimit = flag.Int64("cacheLimit", int64(1e6<<20), "Default maximum key size limit, in bytes")
	target     = flag.String("target", "/var/cache/cookbooks", "The source from which groupcache fetnches from")
)

// TODO(fujin): Refactor and extract.
func main() {
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())

	peers := strings.Split(*peerlist, ",")

	NewHTTPPool(
		*me,
		peers,
	)

	group := NewGroup(
		*groupName,
		*cacheLimit,
	)

	if err := os.MkdirAll(*mountpoint, 0777); err != nil {
		log.Fatal("could not create directory: ", *mountpoint, err)
	}

	conn, err := NewConn(*mountpoint)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// TODO(fujin): Consider a better way to give the groupcache to FUSE.
	td := kasoku.NewTargetDir(
		*target,
		group,
	)
	fs.Serve(conn, td)
}

// NewHTTPPool sets up a new grooupcache pool (maybe with friends)
func NewHTTPPool(me string, peerlist []string) *groupcache.HTTPPool {
	pool := groupcache.NewHTTPPool(me)

	// Voltron mode activate
	if peerlist != nil && len(peerlist) > 0 {
		pool.Set(peerlist...)
	}

	return pool
}

// TODO(fujin): This will likely need to understand/use go-chef/chef.
func getter(ctx groupcache.Context, key string, dest groupcache.Sink) error {
	contents, err := ioutil.ReadFile(key)
	dest.SetBytes(contents)
	return err
}

// NewGroup compositely constructs a groupcache Group
func NewGroup(name string, cacheBytes int64) *groupcache.Group {
	return groupcache.NewGroup(
		name,
		cacheBytes,
		groupcache.GetterFunc(getter),
	)
}

// NewConn returns a connection to a (named) FUSE mount.
// It's not safe to use immediately; it must be checked if it is ready
// Via the conn's Ready chan.
func NewConn(mountpoint string) (*fuse.Conn, error) {
	c, err := fuse.Mount(mountpoint)
	if err != nil {
		log.Fatal(err)
	}

	maybeErr := make(chan error, 1)
	go func(*fuse.Conn, chan error) {
		select {
		case <-c.Ready:
			break
		case <-time.After(time.Duration(60) * time.Millisecond):
			maybeErr <- errors.New("timeout waiting for fuse to become ready")
		}
	}(c, maybeErr)

	return c, err
}
