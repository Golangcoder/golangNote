package main

import (
	"fmt"
)

//问各位老师一个问题：
func main() {
	var a = [3]int{1, 2, 3}
	var b = &a
	fmt.Println(b)
	b[1]++ //b中不是存放a的地址吗？为什么不是*b[1]++
	fmt.Println(a, *b)
}

// b[1]++ ,Go语言这里对“指针的操作”,实际上隐含着对指针的dereference,也可以对指针进行显式的取值，需要添加一个
//括号，来保证运行的优先级
