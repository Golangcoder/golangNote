package main

import (
	"fmt"
)

type Ract struct {
	long  int
	width int
}

func main() {
	a := Ract{
		long:  10,
		width: 2,
	}
	c := a.Hello()
	fmt.Println(c)
}
func (x *Ract) Hello() int {
	return x.long * x.width
}

//为结构体添加方法并调用
/*
哦，当然这里的绑定参数是不是改为self会让你觉得更为习惯呢？
func (self *Ract) Hello int {
	return self.long * self.width
}
*/
