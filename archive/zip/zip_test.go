package zip

import (
	"testing"
)

func TestZipDir(t *testing.T) {
	t.Log(ZipDir("./../../archive", false))
}

func TestZipInfo(t *testing.T) {
	t.Log(ZipInfo("./../../archive.zip"))
	t.Log(Unzip("./../../archive.zip", "./../123"))
}
