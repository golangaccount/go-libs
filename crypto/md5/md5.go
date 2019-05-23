package md5

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"
	"os"
	"strings"
)

var (
	MinBufferSize = 32 * 1024 //最小的buffer尺寸
)

/*
* 计算字符串的md5，并以小写的形式输出
 */
func String(value string) string {
	return Bytes([]byte(value))
}

/*
* 计算byte数组的md5
 */
func Bytes(bts []byte) string {
	result := md5.Sum(bts)
	return strings.ToLower(hex.EncodeToString(result[:]))
}

/*
* 计算文件的md5值(使用默认的buffer尺寸)
 */
func File(path string) (string, error) {
	return FileBuffer(path, 0)
}

/*
* 计算数据流的md5值(使用默认的buffer尺寸)
 */
func Reader(reader io.Reader) (string, error) {
	return ReaderBuffer(reader, 0)
}

/*
* 计算文件的md5值(使用指定的buffer尺寸)
 */
func FileBuffer(path string, size int) (string, error) {
	if fs, err := os.Open(path); err != nil {
		return "", err
	} else {
		defer fs.Close()
		return ReaderBuffer(fs, size)
	}
}

/*
* 计算数据流的md5值(使用=指定的buffer尺寸)
 */
func ReaderBuffer(reader io.Reader, size int) (string, error) {
	if reader == nil {
		return "", errors.New("reader is nil")
	}
	hs := md5.New()
	var err error
	if size > MinBufferSize {
		_, err = io.CopyBuffer(hs, reader, make([]byte, size))
	} else {
		_, err = io.CopyBuffer(hs, reader, make([]byte, MinBufferSize))
	}
	if err != nil {
		return "", err
	}
	return strings.ToLower(hex.EncodeToString(hs.Sum(nil)[:])), nil
}
