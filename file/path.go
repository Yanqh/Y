package file

import (
	. "path"
	"strings"
)


// 获取文件的后缀
func YFileSuffix(path string) (ext string) {
	ext = Ext(path)
	return
}


// 获取文件的全名
func YFileFullName(path string) (fullName string) {
	fullName = Base(path)
	return
}


// 获取文件名
func YFileName(path string) (name string) {
	name = strings.TrimSuffix(YFileFullName(path), YFileSuffix(path))
	return
}

