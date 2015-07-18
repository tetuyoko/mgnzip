package mgnzip

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"

	"github.com/tetuyoko/mgnstr"
)

var Excludes = []string{"__MACOSX", ".DS_Store"}

// check if directory or not from name
func IsDirectory(name string) (isDir bool, err error) {
	fInfo, err := os.Stat(name)
	if err != nil {
		return false, err
	}

	return fInfo.IsDir(), nil
}

// unzip
// src : source zip
// destdir : outputs path(directory)
// paths: outputed all paths
// err: error
func Unzip(src, destdir string) (paths []string, err error) {
	if error := os.MkdirAll(destdir, 0774); error != nil {
		return nil, error
	}

	r, err := zip.OpenReader(src)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	paths  = make([]string, 0)

	for _, f := range r.File {
		if mgnstr.ContainsAny(f.Name, Excludes) {
			continue
		}

		rc, err := f.Open()
		if err != nil {
			return nil, err
		}
		defer rc.Close()

		path := filepath.Join(destdir, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			paths = append(paths, path)

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
