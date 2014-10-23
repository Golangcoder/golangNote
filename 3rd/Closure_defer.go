package main

import (
	"fmt"
)

func main() {
	fs := [4]func(){}              //定义了一个长度为4的数组，数组元素类型为func()
	for i := 0; i < len(fs); i++ { //对数组进行遍历
		defer func() { //defer一个匿名函数，这个函数没有参数传递
			fmt.Println("hello")
		}()
		defer fmt.Println("defer as :", i) //defer一个函数，没有参数，但调用的参数i会从函数外部获取
		defer func() {
			fmt.Println("defer_closure", i)
		}()
		fs[i] = func() {
			fmt.Println("closure i = ", i) //将函数存入数组中，每次存储获得参数i同样从外层函数中获取。
		}
	}
	for _, v := range fs {
		v()
	}
}

/*
#匿名函数，闭包，defer
defer 定义的函数将在函数体return之根据定义的顺序前进行逆向调用
函数定义时没有定义传递参数，函数体内的参数将从外围函数中获取。获取变量地址。
在执行完第一个for循环后，变量i的值变为了4，所以在整个函数体执行到return，refer部分执行时，参数i的值已经为4，所以打印出的i值均为4
而存入数组中的func()中的i是在for循环进行中的i地址，所以他们打印出来的i的值会有变化
*/
