package md5

import (
	"testing"
)

func TestString(t *testing.T) {
	if "202cb962ac59075b964b07152d234b70" == String("123") {

	} else {
		t.Fail()
	}
}

func TestBytes(t *testing.T) {
	if "202cb962ac59075b964b07152d234b70" == Bytes([]byte{49, 50, 51}) {

	} else {
		t.Fail()
	}
}

func TestFile(t *testing.T) {
	if v, _ := File("./123.txt"); v == "202cb962ac59075b964b07152d234b70" {

	} else {
		t.Fail()
	}
}

func TestFileBuffer(t *testing.T) {
	if v, _ := FileBuffer("./123.txt", 1024*1024); v == "202cb962ac59075b964b07152d234b70" {

	} else {
		t.Fail()
	}
}
