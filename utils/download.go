package utils

import (
	"archive/zip"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

func Download(rawurl string, destDir string) (string, error) {

	u, err := url.Parse(rawurl)
	if err != nil {
		return "", err
	}

	filename := filepath.Base(u.Path)

	resp, err := http.Get(rawurl)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	destPath := filepath.Join(destDir, filename)
	f, err := os.OpenFile(destPath, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return "", err
	}

	_, err = io.Copy(f, resp.Body)

	return destPath, nil
}

func Unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	os.MkdirAll(dest, 0755)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), f.Mode())
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}
