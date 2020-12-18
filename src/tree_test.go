package slf

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestScanDirectoryTree(t *testing.T) {
	dir, err := ioutil.TempDir(os.TempDir(), "slf_test")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)

	f, err := os.Create(filepath.Join(dir, "data.txt"))

	if err != nil {
		log.Fatal(err)
	}

	_, err2 := f.WriteString("1234\n")
	f.Close()

	if err2 != nil {
		log.Fatal(err)
	}

	tree := ScanDirectoryTree(dir, runtime.NumCPU()+1)
	if tree.Children["data.txt"].Size != 5 {
		t.Error("Expected", f.Name(), "with 5 bytes but got", tree.Children["data.txt"].Size, "bytes")
	}
}

func BenchmarkRootScan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ScanDirectoryTree("/", runtime.NumCPU())
	}
}
