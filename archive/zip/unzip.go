package zip

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"

	"github.com/golangaccount/go-libs/errors"
	gos "github.com/golangaccount/go-libs/os"
)

//解压文件
func Unzip(input, output string) ([]string, error) {
	result := make([]string, 0)
	if info, err := os.Lstat(input); err != nil || info.IsDir() {
		return result, errors.Error(err, Unzip_InputTypeErr)
	}

	if err := os.MkdirAll(output, os.ModeDir); err != nil {
		return result, errors.New(Unzip_MakeDirErr)
	}

	r, err := zip.OpenReader(input)
	if err != nil {
		return result, errors.New(Unzip_InputZipErr)
	}
	defer r.Close()

	var rc io.ReadCloser
	var dst *os.File
	defer func() {
		if rc != nil {
			rc.Close()
		}
		if dst != nil {
			dst.Close()
		}
	}()
	for _, f := range r.File {
		result = append(result, f.Name) //包含路径和文件
		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(filepath.Join(output, f.Name), os.ModeDir); err != nil {
				return result, errors.New(Unzip_MakeDirErr)
			}
			continue
		}
		rc, err = f.Open()
		if err != nil {
			return result, err
		}
		dst, err = gos.Create(filepath.Join(output, f.Name))
		if err != nil {
			rc.Close()
			return result, err
		}
		_, err = io.Copy(dst, rc) //文件内容copy
		if err != nil {
			rc.Close()
			dst.Close()
			return result, err
		}
		rc.Close()
		dst.Close()
	}
	return result, nil
}

//压缩文件信息
func ZipInfo(filepath string) ([]fileInfo, error) {
	r, err := zip.OpenReader(filepath)
	if err != nil {
		return nil, errors.New(Unzip_InputZipErr)
	}
	result := make([]fileInfo, 0)
	for _, f := range r.File {
		result = append(result, fileInfo{
			f.FileInfo().IsDir(),
			f.Name,
			f.CompressedSize64,
			f.UncompressedSize64,
			f.Comment,
		})
	}
	return result, nil
}

type fileInfo struct {
	IsDir              bool   //是否是文件夹
	Name               string //文件名称(相对路径)
	CompressedSize64   uint64 //压缩后size
	UncompressedSize64 uint64 //压缩前size
	Comment            string //备注
}
