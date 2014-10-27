    package main
    import (
        "fmt"
    )
    func main() {
        a := [10]string{"a", "b", "c", "d", "e"}
        b := make(chan string, 3)
        //将chan的缓存设置成3，当chan里的内容没到3，chan就不会阻塞
        go func() {
            for i := range a {
                b <- a[i]
                fmt.Println(i, a[i])
            }
        }()
        for i := 0; i < len(a); i++ {
            fmt.Println("the value of the chan is ", <-b)
        }
    }
    