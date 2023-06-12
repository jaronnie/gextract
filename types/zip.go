package types

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

type ZIP struct {
}

func (g *ZIP) Extract(source string, p string) error {
	r, err := zip.OpenReader(source)
	if err != nil {
		return err
	}
	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		filename := filepath.Join(p, f.Name)
		if f.FileInfo().IsDir() {
			if err = os.MkdirAll(filename, f.Mode()); err != nil {
				return err
			}
		} else {
			all, err := io.ReadAll(rc)
			if err != nil {
				return err
			}
			if err = os.WriteFile(filename, all, 0o644); err != nil {
				return err
			}
		}
		if err = rc.Close(); err != nil {
			return err
		}
	}
	return r.Close()
}
