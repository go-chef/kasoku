package kasoku

import (
	"log"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"github.com/golang/groupcache"
)

// File is a local FUSE type File, represented internally by a Node.
type File struct {
	Node
}

// ReadAll implements
func (f File) ReadAll(intr fs.Intr) ([]byte, fuse.Error) {
	var contents []byte
	err := f.Node.group.Get(nil, f.Path, groupcache.AllocatingByteSliceSink(&contents))
	if err != nil {
		log.Print(err)
		return nil, fuse.ENOENT
	}
	return contents, nil
}
