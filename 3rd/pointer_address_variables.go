package main

import (
	"fmt"
)

func main() {
	var a int = 10                           //定义一个整型变量，并赋值10
	var p *int = &a                          //定义了一个整型的指针变量p,并赋值变量a的内存地址
	fmt.Println("the value of a is :", a)    //打印变量a的值
	fmt.Println("the address of a is :", &a) //打印变量a的地址
	fmt.Println("the value of p is :", p)    //打印指针变量p的值，即变量a的地址
	fmt.Println("the address of p is :", &p) //打印指针变量p的地址
	fmt.Println("the value of ptr p", *p)    //打印指针变量p中的值(变量a的地址)的值
}

/*
一些关于指针的解释：
一些概念：指针，指针变量，内存地址，值
自己认为：
	指针：变量的内存地址
	指针变量：一个变量，其值为一个内存地址，当然，这个变量也是有自己的内存地址的。
示例中定义了一个变量a,a的值为int型的10
&符号为取地址符号，将变量a的内存地址取出， &a为变量a的内存地址
将&a赋值给p,则变量p的值为变量a的地址，即&a
&p即将指针变量p的内存地址取出
*为取值符号，将指针p对应的内存地址中的值取出
*/
