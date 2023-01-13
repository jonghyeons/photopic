package utils

import (
	"github.com/jonghyeons/photopic/models/constant"
	"strings"
)

func Filter(file string) bool {
	fn := strings.Split(file, ".")
	fileName := strings.Join(fn[:len(fn)-1], "")
	fileExtension := fn[len(fn)-1]
	if Contains(constant.ExtensionBlackList, fileExtension) {
		return true
	}

	if Contains(constant.NameBlackList, fileName) {
		return true
	}
	return false
}
