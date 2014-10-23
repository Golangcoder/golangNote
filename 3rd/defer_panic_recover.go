package main

import (
	"fmt"
)

func main() {
	c, err := Hello(1, 3)
	fmt.Println(c, err)
}
func Hello(a, b int) (c int, err error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("hello world")
		}
	}()
	panic("this is a panic info") //panic 接受任意的类型，比如数字，字符串，对象
	return a + b, err
}

/*
defer,panic,recover
func main() {
	c, err := Hello(1, 0)
	fmt.Println(c, err) //获取函数的多返回值，最后一个为error类型，如果函数执行没有错误，则返回的err为nil
}
func Hello(a, b int) (c int, err error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("something wrong happened")
		}
	}()//defer声明了一个匿名函数，匿名函数会调用参数err，而err并没有显式参数传入，所以匿名函数中的err变量会从他的外围函数
	   //获得，程序最后的除数为0会引发panic，程序终止，执行defer的匿名函数。最终，结果c将为整型变量的零值0，执行Println并
	   //没有出错，所以err为nil

	return a / b, err
}

但如果没有调用recover
func main() {
	c, err := Hello(1, 0)
	fmt.Println(c, err)
}
func Hello(a, b int) (c int, err error) {
	defer func() {
		fmt.Println("hello world")
	}()

	return a / b, err
}
程序依旧会执行refer,但最后还是会panic
panic可以手动触发
func main() {
	c, err := Hello(1, 3)
	fmt.Println(c, err)
}
func Hello(a, b int) (c int, err error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("hello world")
		}
	}()
	panic("this is a panic info")//panic 接受任意的类型，比如数字，字符串，对象
	return a + b, err
}
*/
