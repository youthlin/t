package t

import (
	"io/fs"
	"os"
	"path/filepath"
)

// asFS path as FS. if path is dir return os.DirFS, or return singleFileFS
func asFS(path string) fs.FS {
	return pathFS(path)
}

type pathFS string

func (f pathFS) Open(name string) (fs.File, error) {
	path := string(f)
	// os.DirFS(path): Open(path + "/" + name)
	return os.Open(filepath.Join(path, name))
}
