package io

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	gos "github.com/golangaccount/go-libs/os"
)

//Copy 进行文件复制，source为文件时，进行文件复制，为文件夹时，进行文件夹复制
func Copy(source, dest string) error {
	fi, err := os.Lstat(source)
	if err != nil {
		return err
	}

	if fi.Mode().IsDir() {
		return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if path == source {
				return nil
			}
			if info.IsDir() {
				return nil
			}
			relative := strings.TrimPrefix(path, source)
			return Copy(path, filepath.Join(dest, relative))
		})
	}
	return copyFile(source, dest)
}

func copyFile(source, dest string) error {
	info, err := os.Stat(source)
	if err != nil {
		return err
	}
	size := info.Size()

	s, err := os.Open(source)
	if err != nil {
		return err
	}
	defer s.Close()

	d, err := gos.Create(dest)
	if err != nil {
		return err
	}
	defer d.Close()
	var bytes []byte
	//缓存策略 需要注意的是，会增大内存的开销 但是会加快文件的复制速度
	if size < 1024*1024 {
		bytes = make([]byte, 32*1024) //system copy default
	} else if size < 100*1024*1024 {
		bytes = make([]byte, 1024*1024) //1m~100m 使用1m的缓存
	} else {
		bytes = make([]byte, 10*1024*1024) //100m~ 使用10m的缓存
	}
	write, err := io.CopyBuffer(d, s, bytes)
	if write != size {
		return errors.New("不完整复制，请检查磁盘空间")
	}
	if err != nil {
		return err
	}
	return nil
}
