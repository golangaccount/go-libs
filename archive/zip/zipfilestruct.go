package zip

import (
	"errors"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	gftp "github.com/golangaccount/go-libs/net/ftp"
)

type ZipFileStruct struct {
	FilePath        string                              //需要压缩的文件地址
	ZipPath         string                              //zip文件的地址
	FileMode        os.FileMode                         //文件类型
	ModTime         time.Time                           //
	StreamTransform func(string) (io.ReadCloser, error) //转换读写流
}

func (zfs *ZipFileStruct) Reader() (io.ReadCloser, error) {
	if zfs.StreamTransform != nil {
		return zfs.StreamTransform(zfs.FilePath)
	}
	return zfs.reader()
}

func (zfs *ZipFileStruct) reader() (io.ReadCloser, error) {
	if strings.HasPrefix(zfs.FilePath, "http://") || strings.HasPrefix(zfs.FilePath, "https://") {
		resp, err := http.Get(zfs.FilePath)
		if err != nil {
			return nil, err
		}
		if resp.StatusCode == 200 {
			return nil, errors.New("http status code error,error code:" + resp.Status)
		}
		return resp.Body, nil
	} else if strings.HasPrefix(zfs.FilePath, "ftp://") {
		return gftp.Response(zfs.FilePath)
	} else {
		return os.Open(zfs.FilePath)
	}
}
