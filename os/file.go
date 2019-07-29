package os

import (
	"os"
	"path/filepath"
)

//读写文件（如果文件存在则打开文件，文件不存在则创建）
func RDRW(path string) (*os.File, error) {
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, os.ModeDir|os.ModePerm)
	if err != nil {
		return nil, err
	}
	return os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
}

//创建文件（判断文件夹存不存在，不存在则进行文件夹创建在进行文件创建）
func Create(path string) (*os.File, error) {
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, os.ModeDir|os.ModePerm)
	if err != nil {
		return nil, err
	}
	return os.Create(path)

}

//文件不存在则进行创建,文件存在则向文件中追加内容
func Append(path string) (*os.File, error) {
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, os.ModeDir)
	if err != nil {
		return nil, err
	}
	return os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
}

//判断文件或文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

//判断文件是否存在
func ExistsFile(path string) bool {
	info, err := os.Stat(path)
	if err == nil && !info.IsDir() {
		return true
	}
	return false
}

//判断文件夹是否存在
func ExistsDir(path string) bool {
	info, err := os.Stat(path)
	if err == nil && info.IsDir() {
		return true
	}
	return false
}
