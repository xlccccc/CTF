package utils

import (
	"fmt"
	"io/ioutil"
	"path"
)

func IsImgExist(imgUrl string) bool {
	exist := false
	files, err := ioutil.ReadDir("/tmp/img")
	if err != nil {
		return false
	}
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		if f.Name() == fmt.Sprintf("img-%s", path.Base(imgUrl)) {
			exist = true
			break
		}
	}
	return exist
}
