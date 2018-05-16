package info

import (
	"runtime"
	"fmt"
	. "Y/file"
	"strconv"
)


// 跳出0层，输出打印信息
func YPN(str string) {
	fmt.Println(getFmt(2) + str)
}


// 跳出1层，输出打印信息
func YPN1(str string) {
	fmt.Println(getFmt(3) + str)
}


// 跳出2层，输出打印信息
func YPN2(str string) {
	fmt.Println(getFmt(4) + str)
}


// 自定义跳出层数，输出打印信息
func YPNS(skip int, str string) {
	fmt.Println(getFmt(skip + 2) + str)
}


// 跳过skip层，获取格式化字符串
func getFmt(skip int) (fmt string) {
	pc, fileName, line, _ := runtime.Caller(skip)
	funcName := runtime.FuncForPC(pc).Name()
	fileName = YFileName(fileName)
	fmt = "[" + fileName + "]" + "[" + funcName + "]" + "[" + strconv.Itoa(line)  + "]> "
	return
}