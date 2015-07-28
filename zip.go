package mgnzip

import (
	"io"
	"os"
	"path"

	"archive/zip"
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
	if err = os.MkdirAll(destdir, 0774); err != nil {
		return nil, err
	}

	r, err := zip.OpenReader(src)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	paths = make([]string, 0)

	for _, f := range r.File {
		if mgnstr.ContainsAny(f.Name, Excludes) {
			continue
		}

		rc, err := f.Open()
		if err != nil {
			return nil, err
		}
		defer rc.Close()

		// dir無しでくるときがあるのでつくってしまう
		tpath := filepath.Join(destdir, f.Name)
		os.MkdirAll(path.Dir(tpath), 0777)

		if f.FileInfo().IsDir() {
			os.MkdirAll(tpath, f.Mode())
		} else {
			paths = append(paths, tpath)

			fo, err := os.OpenFile(
				tpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())

			if err != nil {
				return nil, err
			}
			defer fo.Close()

			if _, err = io.Copy(fo, rc); err != nil {
				return nil, err
			}
		}
	}

	return paths, nil
}
