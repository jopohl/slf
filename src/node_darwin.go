package slf

type Node struct {
	Parent   *Node
	Children map[string]*Node
	Path     string
	Size     int64
	IsDir    bool
	Nlink    uint16
}
