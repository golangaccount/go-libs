package ftp

import (
	"errors"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	gos "github.com/golangaccount/go-libs/os"
	"github.com/jlaffaye/ftp"
)

const FtpScheme = "ftp"

//download file from remote
func DownLoad(remote, local string) error {
	resp, err := Response(remote)
	if err != nil {
		return err
	}
	defer resp.Close()
	f, err := gos.Create(local)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, resp)
	return err
}

type FtpInfo struct {
	User     string
	Password string
	Service  string
	Port     string
	Path     string
}

func (fi *FtpInfo) Check() bool {
	if strings.TrimSpace(fi.Service) == "" {
		return false
	}
	if strings.TrimSpace(fi.Port) == "" {
		fi.Port = "21"
		return true
	} else {
		i, err := strconv.Atoi(fi.Port)
		if err != nil {
			return false
		} else if i < 0 || i > 1<<16 {
			return false
		} else {
			return true
		}
	}
}

func (fi *FtpInfo) Address() string {
	return fi.Service + ":" + fi.Port
}

//get download file response stream
func Response(remote string) (io.ReadCloser, error) {
	ftpinfo, err := parse2FtpInfo(remote)
	if err != nil {
		return nil, err
	}

	ser, err := ftp.Connect(ftpinfo.Address())
	if err != nil {
		return nil, err
	}
	defer ser.Quit()
	err = ser.Login(ftpinfo.User, ftpinfo.Password)
	if err != nil {
		return nil, err
	}
	return ser.Retr(ftpinfo.Path)
}

//upload local file to remote
func Upload(remote, local string) error {
	f, err := os.Open(local)
	if err != nil {
		return err
	}
	defer f.Close()
	return Request(remote, f)

}

//upload reader stream data to remote
func Request(remote string, reader io.Reader) error {
	ftpinfo, err := parse2FtpInfo(remote)
	if err != nil {
		return err
	}
	ser, err := ftp.Connect(ftpinfo.Address())
	if err != nil {
		return err
	}
	defer ser.Quit()
	err = ser.Login(ftpinfo.User, ftpinfo.Password)
	if err != nil {
		return err
	}
	//判断目录是否存在
	err = ser.ChangeDir(filepath.Dir(ftpinfo.Path))
	if err == nil { //目录跳转成功，表示存在
	} else { //目录跳转失败，尝试创建目录
		err = ser.MakeDir(filepath.Dir(ftpinfo.Path))
		if err != nil {
			return err
		}
	}

	return ser.Stor(ftpinfo.Path, reader)
}

//
func parse2FtpInfo(remote string) (*FtpInfo, error) {
	info, err := url.Parse(remote)
	if err != nil {
		return nil, err
	}
	if info.Scheme != FtpScheme {
		return nil, errors.New("地址协议错误")
	}
	ftpinfo := &FtpInfo{}
	ftpinfo.User = info.User.Username()
	ftpinfo.Password, _ = info.User.Password()
	ftpinfo.Service = info.Hostname()
	ftpinfo.Port = info.Port()
	ftpinfo.Path = info.Path
	if !ftpinfo.Check() {
		return nil, errors.New("地址信息错误")
	}
	return ftpinfo, nil
}
