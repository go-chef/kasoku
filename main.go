// Hellofs implements a simple "hello world" file system.
package main

import (
	"flag"
	"log"

	"github.com/go-chef/kasoku/cookbookfs"

	_ "bazil.org/fuse/fs/fstestutil"
)

var (
	mountpoint = flag.String("mountpoint", "", "The FUSE mount point target / connection")
	source     = flag.String("source", "", "The server-side cookbook storage location")
)

func main() {
	flag.Parse()

	if *mountpoint == "" && *source == "" {
		log.Fatal("kasoku.mountpoint cannot be empty on clients")
	} else {
		log.Println("server running with local FUSE mount")
	}

	err := cookbookfs.NewFilesystem(*mountpoint)
	if err != nil {
		log.Fatal(err)
	}
}
