package main

import (
	"fmt"
)

type Test struct { //type关键字+结构名称+struct关键字+{}
	Name string
	Age  int
}

func main() {
	a := Test{
		Name: "zhanghua",
		Age:  13,
	}
	fmt.Println(a) //a就是指向结构Test的指针
	Hello(a)       //调用函数Hello时，进行指针传递,并进行属性操作
	fmt.Println(a)
}
func Hello(per Test) {
	per.Name = "lisan"
	fmt.Println(per.Name)
}

/*
结构的声明，显示声明，匿名声明，声明时，字段匿名的话，则结构初始化时则需要注意字面初始化的顺序。结构是一种值类型，进行参数传递时
为值的拷贝，考虑计算效率和属性操作，可以将结构的指针在函数中进行传递。

	结构的匿名声明：
	a:=struct{
		Name string
		Age int
	}
	匿名结构的初始化：
	a:=struct{
		Name string,
		Age int,
	}{
		Name:"zhanghua",
		Age:13,
	}
	但如果结构中的字段也是匿名的呢？
	type Test struct {
		string
		int
	}

	func main() {
		a := Test{"hello", 12}//按照声明字段类型的顺序进行初始化，如果顺序错乱，编译器就会报错。
		fmt.Println(a)
	}
	结构的显示声明：
	var Test struct{}//这样就声明一个空的结构
	var Test2 struct{
		Name string
		Age int
	}//结构中包含了Name和Age两个字段，字段类型分别为string和int
	初始化：
	a:=Test{
		Name:"zhanghua",
		Age:12,
	}//每个字段赋值之后，都要跟一个逗号来分割。
结构嵌套匿名类型时，如何对匿名结构进行初始化和属性操作
var a struct{
	Name string
	Age int
	Content struct{//嵌套匿名结构
		Address string
		Phone string
	}
}
属性更改：
var b a=b{
	Name string,
	Age 12,
}
b.Content.Address = "beijing"//考虑下内嵌结构中的字段名称和外部结构的字段相同的情况，这样就可以正确的进行赋值操作
当然内嵌一个显示声明的结构也是一样的：
var a struct{
	Name string
	Age int
}
var b struct{
	Address string
	Content a
}
那么，如何对内嵌显示声明结构的结构进行初始化赋值呢？
type test1 struct{
	Name string
}
type test2 struct{
	Age int
	test1//结构test2中嵌套了结构test1
}//初始化和值操作：
a:=test2{
	Age:12,
	test1:test1{
		Name:"zhanghua",
	}
}
类似class的属性操作也可：
a.Name = "lisan",内嵌一个结构，那么内部结构的字段将给予外部结构中，那么属性操作就可以a.Name = "lisan",
但如果内嵌结构的字段名称与外层结构字段相同怎么办呢？
a.Test.Name = "lisan"即可

结构的字面初始化，属性更改
type Test struct {
	Name string
}

func main() {
	a := Test{
		Name: "zhanghua",
	}
	fmt.Println(a)
	Hello(a)
	fmt.Println(a)
}
func Hello(test Test) {
	test.Name = "test"
	fmt.Println(test)
}
结构为值类型，进行参数传递时将为值的复制，类型包含了结构名称和对应字段及其字段类型
type Test1 struct {
	Name string
}
type Test2 struct {
	Name string
}
t1 := Test1{
		Name: "Test1",
	}
	t2 := Test2{}
	t2 = t1
即使两个结构声明的字段和字段属性相同，但他们的结构名称不同，则两者之间也不可以进行比较，同样也就不能进行赋值
（提示cannot use t1 (type Test1) as type Test2 in assignment，即两个不同的类型之间不能赋值）
可以通过指针传递，来更改结构中的属性值
type Test struct {
	Name string
	Age  int
}

func main() {
	a := Test{
		Name: "zhanghua",
		Age:  13,
	}
	fmt.Println(a) //a就是指向结构Test的指针
	Hello(a)       //调用函数Hello时，进行指针传递,并进行属性操作
	fmt.Println(a)
}
func Hello(per Test) {
	per.Name = "lisan"
	fmt.Println(per.Name)
}最后输出的a的Name依旧为zhanghua,并没有被改变。
在结构初始化时就取出结构的地址，使变量为指向结构的指针，但指针的属性操作依旧为类似class的属性操作。a.Age = 15,在参数接受的地方，去参数的值，*arg
type Test struct {
	Name string
	Age  int
}

func main() {
	a := &Test{
		Name: "zhanghua",
		Age:  13,
	}
	fmt.Println(a) //a就是指向结构Test的指针
	Hello(a)       //调用函数Hello时，进行指针传递,并进行属性操作
	fmt.Println(a)
}
func Hello(per *Test) {
	per.Name = "lisan"
	fmt.Println(per.Name)
}
最后输出的结构属性Name已经为lisan，即指针传递改变了结构的字段值
*/
