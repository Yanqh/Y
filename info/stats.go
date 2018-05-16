package info

import (
	"time"
	"yanqh/info"
)


// 计算出函数运行时间
func ExeTime(t time.Time) {
	exeTime := time.Since(t).String()
	fmt := "Elapsed Time : " + exeTime
	info.YPN1(fmt)
}
