我们可以为类型绑定方法，那么，通过什么方式来实现方法的绑定呢？没错，就是receiver。下面的self就被称为
receiver
    `type T struct{name string}//定义了一个结构，并添加了一个string类型的name字段
    func (self *T) Hello(){self.name = "Hello"}//通过receiver的类型来判断这个方法和谁绑定，这里就是T类型
    func (self T) World(){self.name = "World"}//这里就是World方法和T类型的指针类型绑定`
这里说，类型T关联一个MethodSet，类型*T也关联一个MethodSet，但是，*T的MethodSet包含了T的MethodSet，要求
就是MethodSet的方法名唯一。那么这两者之间到底有什么联系呢？上代码：
    `a:=T{}//初始化了一个裸值类型
    a.Hello()
    ptr:=new(T)//初始化了一个指针类型
    ptr.Hello()
    `
虽然两者类型不同，一个为裸值，一个为指针，但他们都成功调用了Hello方法。也就是说，虽然receiver的类型为指针，
就如上面说的MethodSet，他也是包含了裸值类型T的MethodSet，对象在被传入的时候，系统会判断传入对象到底是裸值类型
还是指针类型，如果是指针类型，会发生类似(*self).name = "hello"这样的变化（这里并不准，只是为了理解）。

Golang.org的wiki中也有对MethodSet的说明：
[MethodSet](https://code.google.com/p/go-wiki/wiki/MethodSets)
