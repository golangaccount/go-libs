package ftp

import (
	"testing"
)

func TestDownload(t *testing.T) {
	t.Log(DownLoad("ftp://bkUpload:bkUpload@127.0.0.1/D7D1936EAC7447C5B9EAE49B3BE2D438.STD", "./ftp.std"))
}

func TestUpload(t *testing.T) {
	t.Log(Upload("ftp://bkUpload:bkUpload@127.0.0.1/123/456/ftp.std", "./ftp.std"))
}
