package kasoku

import (
	"log"
	"os"

	"github.com/golang/groupcache"

	"bazil.org/fuse"
)

// Node wraps around a Path
type Node struct {
	Path  string
	group *groupcache.Group
}

// Attr returns a fuse.Attr based on a node.
func (n Node) Attr() fuse.Attr {
	s, err := os.Stat(n.Path)
	if err != nil {
		log.Print(err)
		return fuse.Attr{}
	}

	return fuse.Attr{Size: uint64(s.Size()), Mtime: s.ModTime(), Mode: s.Mode()}
}
