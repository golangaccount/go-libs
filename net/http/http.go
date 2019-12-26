package http

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/golangaccount/go-libs/os"
)

//Save 保存到本地文件
func Save(url string, dest string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	f, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer f.Close()
	size, err := io.Copy(f, resp.Body)
	if resp.ContentLength != size {
		return fmt.Errorf("保存数据不完整,需要保存数据为%d,实际保存%d", resp.ContentLength, size)
	}
	if err != nil {
		return err
	}
	return nil
}

//GetBytes 获取所有bytes
func GetBytes(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

//GetReader 返回reader流
//需要注意的是，此时，http请求会被hode住，并且用完后关闭reader流
func GetReader(url string) (io.Reader, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}
