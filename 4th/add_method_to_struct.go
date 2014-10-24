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
