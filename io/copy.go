package io

import (
	"os"
	"path/filepath"
	"strings"
)

//Copy 进行文件复制，source为文件时，进行文件复制，为文件夹时，进行文件夹复制
func Copy(source, dest string) error {
	fi, err := os.Lstat(source)
	if err != nil {
		return err
	}

	if fi.Mode().IsDir() {
		filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if path == source {
				return nil
			}
			relative := strings.TrimPrefix(path, source)
			return Copy(path, filepath.Join(dest, relative))
		})
	}
	return copyFile(source, dest)

}

func copyFile(source, dest string) error {
	return nil
}
