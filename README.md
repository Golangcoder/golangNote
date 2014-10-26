#Golang学习笔记
零基础学习Golang编程
####//目录
#####[第三天](https://github.com/Golangcoder/golangNote/tree/master/3rd)**更新时间** 2014年10月23日
#####[第四天](https://github.com/Golangcoder/golangNote/tree/master/4th)**更新时间** 2014年10月24日
#####[第五天](https://github.com/Golangcoder/golangNote/tree/master/5th)**更新时间** 2014年10月25日
#####[第六天](https://github.com/Golangcoder/golangNote/tree/master/6th)**更新时间** 2014年10月26日
####//内容
#####第三天：
- 指针和变量
- 闭包
- 结构，包含了结构的声明，初始化，结构属性的操作，以及结构在函数中的传递与操作
- 一个文件复制的小例子
- defer,panic,recover
- 文件读取写入，并应用冒泡排序的小例子
- 为类型添加方法，利用receiver将方法与类型进行绑定

#####第四天:
- 再叙指针，Golang中没有对指针的操作，如p是个指向切片的指针，类似p[1] = 10是可以执行成功的。(*p)[1] = 10
- 结构中匿名字段如何在嵌入结构中操作呢？子结构会将父结构中的字段进行隐藏，当然也可以显式的调用更改。
- panic与recover,recover在defer中调用则会终止panic序列（panic sequence)并将panic信息传递给调用者。
- 类型方法绑定以及两种方法调用的方式：method-value, method-expression

#####第五天：
- 接口，嵌套接口，为结构实现接口，如何通过接口条用结构中的字段值，将对象赋值给接口实际上接口内部存储的是对象的复制品的指针，所以，对接口赋值之后，再对原对象进行属性修改，将不会影响到接口内部。

#####第六天：
- MethodSet，自己理解。类型*T的方法设置包含了T类型的方法。程序会根据传入对象的类型来对应实现各自的方法。
