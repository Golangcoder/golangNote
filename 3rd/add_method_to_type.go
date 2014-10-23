package main

import (
	"fmt"
)

type Test int //声明一个底层类型为int的结构（类型别名）

func main() {
	var a Test
	/*这里我惯性的以为跟结构的初始化相同，类似：
	a:=Test{} 出现了错误error:invalid type for composite literal
	前面类型声明处，只是给int型取了个别名Test，这里需要显示声明变量a为类型Test
	*/
	a.Hello(200)
	fmt.Println(a)
}
func (a *Test) Hello(num int) {
	/*接收者为a,采用指针传递，为了更改类型Test的值，这样操作，也会节省值拷贝所需的内存空间*/
	*a += Test(num)
	/*虽然类型Test底层为int，但type关键字声明的类型包含了类型名称，所以Test和int为两个不同的类型，
	二者相加，需要进行强制类型转换*/
}

/*
前面学习了结构中对字段的操作，那么，结构是不是也有像类一样的方法呢？答案是肯定的。
为结构来绑定方法：
采用receiver来讲函数与结构进行绑定：
func (b *Test) Hello() {//func关键字后跟一对括号，即receiver，括号内是一个结构类型的参数b，(这里故意将参数设成了
与结构类型初始化赋值给的变量a不同，即此变量仅仅是个参数)
	b.Name = "golang"
	fmt.Println(b.Name)//func()内调用也要跟随传入参数b
}
接着看main函数中如何进行的方法调用：
func main() {
	a := Test{
		Name: "Hello",
	}
	a.Hello()
	fmt.Println(a.Name)
}
字面初始化了一个结构Test，并给字段Name赋值“hello”，
a.Hello()调用了与结构绑定的方法Hello
在结构声明中，使用了type关键字，所以，结构也是一种类型。
*/
