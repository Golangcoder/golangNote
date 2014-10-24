package main

import (
	"fmt"
)

func main() {
	Hello()
}
func Hello() {
	defer func() {
		fmt.Println("hello world", recover())
	}()
	defer func() {
		fmt.Println("recover panic")
	}()
	panic("raise panic")
}

//recover在defer的匿名函数中被调用，而非其他函数调用时，recover函数会将panic值进行传递，并停止panic队列，如果程序正常执行退出。
//recover 同样defer调用，则recover返回nil
