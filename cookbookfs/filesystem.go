package cookbookfs

import (
	"errors"
	"log"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

// NewFilesystem creates a new connection/mountpoint for FUSE
func NewFilesystem(mountpoint string) (err error) {
	conn, err := fuse.Mount(
		mountpoint,
		fuse.FSName("kasoku"),
		fuse.Subtype("cookbookfs"),
		fuse.LocalVolume(),
		fuse.VolumeName("Accelerated local cookbook distribution and storage"),
	)

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	err = NewServer(conn)
	if err != nil {
		return err
	}

	// check if the mount process has an error to report
	<-conn.Ready
	if err := conn.MountError; err != nil {
		return err
	}

	return nil
}

// NewServer either creates a new FUSE server backend "on" a connection or returns err
func NewServer(conn *fuse.Conn) (err error) {
	if conn == nil {
		return errors.New("connection was nil")
	}
	err = fs.Serve(conn, FS{})
	return err
}
