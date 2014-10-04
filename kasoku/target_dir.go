package kasoku

import (
	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"github.com/golang/groupcache"
)

// TODO(fujin): These are all stupid
type TargetDir struct {
	Path  string
	group *groupcache.Group
}

func NewTargetDir(path string, cache *groupcache.Group) TargetDir {
	return TargetDir{
		Path:  path,
		group: cache,
	}
}

// Root satisfies the interface specified by FUSE fs.Serve
func (td TargetDir) Root() (fs.Node, fuse.Error) {
	return Dir{
		Node{
			Path:  td.Path,
			group: td.group,
		},
	}, nil
}
