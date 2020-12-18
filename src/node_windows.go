package slf

import (
	"os"
)

type Node struct {
	Parent   *Node
	Children map[string]*Node
	Path     string
	Size     int64
	IsDir    bool
	Nlink    uint64
}

func setNlink(entry os.FileInfo, child *Node) {
	// No hardlinks on Windows
}
