package zip

import (
	"os"
	"path/filepath"
	"testing"
)

func TestT(t *testing.T) {
	filepath.Walk("C:\\workspace\\ftp", func(path string, info os.FileInfo, err error) error {
		t.Log(info.Name())
		return nil
	})
}

func TestZipDir(t *testing.T) {
	t.Log(ZipDir("."))
}
