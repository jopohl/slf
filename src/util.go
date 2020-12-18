package slf

import "fmt"

// convert a size in bytes to human readable string such as 4096 -> 4 KiB
func SizeToHumanReadable(size int64) string {
	suffixes := [...]string{"B", "KiB", "MiB", "GiB", "TiB", "PiB", "EiB", "ZiB", "YiB"}
	dsize := float64(size)
	var resultingSuffix string
	for _, suffix := range suffixes {
		resultingSuffix = suffix
		if dsize >= 1024 && suffix != "YiB" {
			dsize /= 1024
		} else {
			break
		}
	}

	return fmt.Sprintf("%-5.4g %-3s", dsize, resultingSuffix)
}

func PrintSizes(files []*Node, numEntries int) {
	i := 0
	for _, file := range files {
		if file == nil {
			continue
		}
		if i >= numEntries {
			return
		}
		i++
		fmt.Println(SizeToHumanReadable(file.Size), file.Path)
	}
}
