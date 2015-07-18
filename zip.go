package mgnzip

import (
	"archive/zip"
	"github.com/tetuyoko/mgnstr"
	"io"
	"os"
	"path/filepath"
)

var Excludes = []string{"__MACOSX", ".DS_Store"}

// src : source zip
// dest : outputs path
// paths: outputed all paths
// err: error
func Unzip(src, dest string) (paths []string, err error) {
	r, err := zip.OpenReader(src)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	//var paths []string
	paths = make([]string, len(r.File))

	for i, f := range r.File {
		if mgnstr.ContainsAny(f.Name, Excludes) {
			continue
		}
		rc, err := f.Open()
		if err != nil {
			return nil, err
		}
		defer rc.Close()

		path := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			paths[i] = path

			f, err := os.OpenFile(
				path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return nil, err
			}
			defer f.Close()

			if _, err = io.Copy(f, rc); err != nil {
				return nil, err
			}
		}
	}

	return paths, nil
}
