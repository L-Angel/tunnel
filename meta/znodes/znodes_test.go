package znodes

import (
	"fmt"
	"strings"
	"testing"
)

func _CreatePathWithDataIfNecessary(path string) {
	var pSeg []string
	if strings.HasPrefix(path, "/") {
		pSeg = strings.Split(path[1:], "/")
	} else {
		pSeg = strings.Split(path, "/")
	}
	p := ""
	for idx, s := range pSeg {
		p += "/" + s
		existed := false
		if !existed && idx == len(pSeg)-1 {
			fmt.Println(p, "demo data")
		} else if !existed {
			fmt.Println(p)
		}
	}
}
func TestCreatePathWithDataIfNecessary(t *testing.T) {
	path1 := "/path1/path2/path3"
	_CreatePathWithDataIfNecessary(path1)
	path2 := "path11/path12/path13"
	_CreatePathWithDataIfNecessary(path2)

	path3 := "/path1/path2/path3/"
	_CreatePathWithDataIfNecessary(path3)

}
