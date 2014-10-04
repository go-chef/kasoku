package kasoku

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"syscall"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

// Dir wraps a Node
// TODO(fujin): See?
type Dir struct {
	Node
}

// Lookup satisfies the lookup interface for FUSE server
func (d Dir) Lookup(name string, intr fs.Intr) (fs fs.Node, error fuse.Error) {

	path := filepath.Join(d.Path, name)
	s, err := os.Stat(path)
	if err != nil {
		log.Print(err)
		return nil, fuse.ENOENT
	}
	node := Node{
		Path:  path,
		group: d.Node.group,
	}
	switch {
	case s.IsDir():
		fs = Dir{node}
	case s.Mode().IsRegular():
		fs = File{node}
	default:
		fs = node
	}

	return
}

// ReadDir satisfies ReadDir interface in FUSE server
func (d Dir) ReadDir(intr fs.Intr) ([]fuse.Dirent, fuse.Error) {
	var out []fuse.Dirent
	files, err := ioutil.ReadDir(d.Path)
	if err != nil {
		log.Print(err)
		return nil, fuse.Errno(err.(syscall.Errno))
	}
	for _, node := range files {
		de := fuse.Dirent{Name: node.Name()}
		if node.IsDir() {
			de.Type = fuse.DT_Dir
		}
		if node.Mode().IsRegular() {
			de.Type = fuse.DT_File
		}
		out = append(out, de)
	}

	return out, nil
}
