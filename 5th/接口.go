package main

import (
	"fmt"
)

type USB interface {
	Name() string
	Connector
}
type Connector interface {
	Connect()
}
type Phoneconnector struct {
	name string
}

func (self Phoneconnector) Connect() {
	fmt.Println(self.name, "Conenected")
}
func (self Phoneconnector) Name() string {
	return self.name
}
func (self Phoneconnector) Test() {
	fmt.Println("hello test")
}
func main() {
	var a interface{}
	fmt.Println(a == nil)
	var p *int = nil
	a = p
	fmt.Println(a == nil)
}
func Disconnect(usb interface{}) {
	// if pc, ok := usb.(Phoneconnector); ok {
	// 	fmt.Println(pc.name, "disconnected")
	// 	return
	// }
	switch v := usb.(type) {
	case Phoneconnector:
		fmt.Println("Disconnected:", v.name)
	default:
		fmt.Println("Unknown Device")
	}
	return
}

/*
1.接口是什么？自己理解就是一个方法集
2.接口如何定义？
type USB interface{
	Name()
	Connector() string
}
用type关键字+接口名+关键字interface+{},大括号内是这个接口的方法，这样，接口就定义好了。
3.如何实现接口呢？
Golang中无需显式声明我实现了某个接口，只需要我有该接口所有的方法，那么，就说，我实现了这个接口。
当然，我也可以有该接口以外的方法，总之，如果一个结构实现了某个接口，那么该结构的方法集是这个接口的超集。
定义一个接口，为该结构绑定方法，这些方法需要满足接口中定义的方法，那么这个结构也就实现了该接口，代码：
type PhoneConnecter struct{name string}
func (self PhoneConnecter)Connector()string{
	fmt.Println("Connecter")
}
func (self PhoneConnecter)Name(){fmt.Println("name")}
var a USB//声明一个类型为USB接口的变量a
a = PhoneConnecter{name:"Nokia"}//初始化PhoneConnecter这个结构，并字面赋值name为Nokia，再将它赋值给USB接口类型的变量a
那么，可以通过对象a来调用接口的方法，a.Name() a.Connecter()
结构PhoneConnecter中还有一个字段name，那么，USB类型的a中并没有这个字段，该如何访问呢？
通过类型判定来调用这个属性。教程中给出了两种方法：
第一种：
	func Disconnect(usb USB){//定义了一个方法，传入了类型为USB的参数
	if k,ok:=usb.(PhoneConnecter);ok{
	fmt.Println(k.name,"Disconnected")
	return
	}
	fmt.Println("Uknown Device")
	}
第二种：
	func Disconnecter(usb interface{}){
	switch v:=usb.(type){
	case Phoneconnecter:
		fmt.Println(v.name,"Disconnected")
	default:
		fmt.Println("Uknown Device")
	}
	}
4.接口的转换：拥有方法超集的接口可以向拥有方法子集的接口进行转换，但逆向转换是不可以的。也就是说，接口转换遵循着从
大到小的转换原则。
5.接口对象和类型都为nil才为nil
6.对象赋值给接口时，是值的复制，接口内部存储的是指向这个复制品的指针。无法获取复制品的指针，也无法修改复制品的状态。
7.指针和值的方法集问题，前面遇到了这样一种情况：
var p *int 
var a int = 10
p = &a
p++
fmt.Println(*p)//11
也就是说，传入一个指针，因为指针的方法集中包含了取指针对应值这样的方法，那么内部就有如同：(*p)++这样的存在，但是反过来却不成立。
传入一个值，那么，他只有值本身的方法。不会进行隐式的调用转换。
*/
