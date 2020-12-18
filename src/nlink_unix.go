// +build aix darwin dragonfly freebsd js,wasm linux nacl netbsd openbsd solaris

package slf

import (
	"os"
	"syscall"
)

func setNlink(entry os.FileInfo, child *Node) {
	if sys := entry.Sys(); sys != nil {
		if stat, ok := sys.(*syscall.Stat_t); ok {
			child.Nlink = stat.Nlink
		}
	}
}
