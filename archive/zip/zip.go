package zip

import (
	"archive/zip"
	"errors"
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
		fh := &zip.FileHeader{
			Name:   item.ZipPath,
			Flags:  1 << 11, // 使用utf8编码 解决中文乱码的问题
			Method: zip.Deflate,
		}
		if item.FileMode != 0 {
			//mode设置
			if fh.Mode().IsDir() != item.FileMode.IsDir() {
				return errors.New(Zip_FileModeErr)
			}
			fh.SetMode(item.FileMode)
		}
		if !item.ModTime.IsZero() {
			fh.SetModTime(item.ModTime)
		}
		dst, err := zf.CreateHeader(fh)
		if err != nil {
			return err
		}
		if fh.Mode().IsDir() {
			continue
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

func ZipDir(path string, skipempty bool, out ...string) error {
	zipstruct := make([]*ZipFileStruct, 0)
	var err error
	path, err = filepath.Abs(path)
	if err != nil {
		return nil
	}
	if info, err := os.Stat(path); err != nil {
		return err
	} else if info.IsDir() {
		length := len(path)
		filepath.Walk(path, func(walkpath string, info os.FileInfo, err error) error {
			if walkpath == path {
				return nil
			}
			if !info.IsDir() {
				zipstruct = append(zipstruct, &ZipFileStruct{
					FilePath: walkpath,
					ZipPath:  walkpath[length+1:],
					FileMode: info.Mode(),
					ModTime:  info.ModTime(),
				})
			} else if !skipempty {
				zipstruct = append(zipstruct, &ZipFileStruct{
					FilePath: walkpath,
					ZipPath:  walkpath[length+1:] + "/",
					FileMode: info.Mode(),
					ModTime:  info.ModTime(),
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
