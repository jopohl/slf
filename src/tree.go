package slf

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

var numTraverseThreads = runtime.NumCPU() + 1

// add a node to list and keep it sorted (biggest first)
// if node's size is smaller than smallest size in list, do nothing
// the list will keep it's size and do not grow
func insertSorted(nodes *[]*Node, node *Node) {
	for i, n := range *nodes {
		if n == nil {
			(*nodes)[i] = node
			break
		}

		if node.Size > n.Size {
			for j := len(*nodes) - 1; j >= i+1; j-- {
				(*nodes)[j] = (*nodes)[j-1]
			}
			(*nodes)[i] = node
			break
		}
	}
}

func UpdateSizes(node *Node, entries *[]*Node, dirMode bool) int64 {
	result := node.Size
	for _, child := range node.Children {
		result += UpdateSizes(child, entries, dirMode) / int64(child.Nlink)
	}
	node.Size = result

	// file mode
	if !dirMode && !node.IsDir {
		insertSorted(entries, node)
	}

	// scan for top level directories in dir mode
	if dirMode && node.IsDir && node.Parent != nil && node.Parent.Parent == nil {
		insertSorted(entries, node)
	}

	return result
}

func scanDirectoryTree(directory string, dirTree *Node, wg *sync.WaitGroup, routineStarted bool) {
	startNewRoutine := runtime.NumGoroutine() <= numTraverseThreads
	if routineStarted {
		defer wg.Done()
	}

	f, err := os.Open(directory)
	if err != nil {
		return
	}
	defer f.Close()

	entries, err := f.Readdir(-1)
	if err != nil {
		return
	}
	for _, entry := range entries {
		path := filepath.Join(directory, entry.Name())
		if strings.HasPrefix(path, "/proc") {
			continue
		}

		child := &Node{dirTree, map[string]*Node{}, path, entry.Size(), entry.IsDir(), 1}
		dirTree.Children[entry.Name()] = child

		if entry.IsDir() {
			if startNewRoutine {
				wg.Add(1)
				go scanDirectoryTree(path, child, wg, true)
			} else {
				scanDirectoryTree(path, child, wg, false)
			}

		} else {
			setNlink(entry, child)
		}
	}
}

func ScanDirectoryTree(directory string, numThreads int) *Node {
	numTraverseThreads = numThreads
	root := &Node{nil, map[string]*Node{}, "", 0, true, 1}
	var wg sync.WaitGroup
	scanDirectoryTree(directory, root, &wg, false)
	wg.Wait()

	return root
}
