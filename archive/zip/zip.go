package zip

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"

	gos "github.com/golangaccount/go-libs/os"
)

func Zip(files []*ZipFileStruct, out string) error {
	f, err := gos.Create(out)
	if err != nil {
		return err
	}
	defer f.Close()
	zf := zip.NewWriter(f)
	defer zf.Close()
	for _, item := range files {
		dst, err := zf.CreateHeader(&zip.FileHeader{
			Name:   item.ZipPath,
			Flags:  1 << 11, // 使用utf8编码 解决中文乱码的问题
			Method: zip.Deflate,
		})
		if err != nil {
			return err
		}
		src, err := item.Reader()
		if err != nil {
			return err
		}
		_, err = io.Copy(dst, src)
		if err != nil {
			src.Close()
			return err
		}
		src.Close()
	}
	return zf.Close()
}

func ZipDir(path string, out ...string) error {
	zipstruct := make([]*ZipFileStruct, 0)
	if info, err := os.Stat(path); err != nil {
		return err
	} else if info.IsDir() {
		length := len(path)
		filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				zipstruct = append(zipstruct, &ZipFileStruct{
					FilePath: path,
					ZipPath:  path[length+1:],
				})
			}
			return nil
		})
	} else {
		zipstruct = append(zipstruct, &ZipFileStruct{FilePath: path, ZipPath: info.Name()})
	}
	var outpath string
	if len(out) == 0 {
		path, err := filepath.Abs(path)
		if err != nil {
			return err
		}
		dic := filepath.Dir(path)
		name := filepath.Base(path)
		index := strings.LastIndex(name, ".")
		if index == -1 {
			index = len(name)
		}
		name = name[:index]
		outpath = filepath.Join(dic, name+".zip")
	} else {
		outpath = out[0]
	}
	return Zip(zipstruct, outpath)
}
