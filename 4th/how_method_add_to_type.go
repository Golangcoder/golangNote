package main

import (
	"fmt"
)

type Test int

func main() {
	var a Test
	a.Hello(100) //这种类似其他语言中类的调用方式，被称为method-value
	fmt.Println(a)
	(*Test).Hello(&a, 100) //这种通过类型进行调用方法，并向方法传递receiver，这种方法调用方式被称为method-expression
}

func (a *Test) Hello(num int) {
	*a = *a + Test(num)
	fmt.Println(*a)
}

/*关于Test这个底层类型为int的类型，他的初始化需要显式的声明变量为Test型
如果这样：
a:=0//这样编译器会将a设置为int型
a:=Test{}//Test是类型没错，但不是结构，不能这样初始化
再次解释下 方法如何与类型进行绑定：通过receiver
关键字func定义了一个方法。不过在方法名前强行插入了一个参数，没错，这就是receiver，a作为一个接收者的参数
而接收者类型为Test类型的指针，编译器就是靠着receiver的类型和对应的类型进行方法绑定。
底层类型为int这样的内置类型，不会将int类型的方法也传递到Test中。这和结构的嵌套是不同的。
方法的绑定需要在同一个包中实现，不能在包A中实现类型TZ，包B中实现对类型TZ的方法绑定
*/
