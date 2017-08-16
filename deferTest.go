/**
 * golang的defer关键字，它可以在函数返回前执行一些操作，最常用的就是打开一个资源（例如一个文件、数据库连接等）时就用defer延迟关闭改资源，以免引起内存泄漏。
 * 在函数块中使用defer，相当于函数对应的栈空间，先进后出。函数结束后调用栈，进行defer操作。
*/
package main

import "fmt"
import "time"

type User struct {
        username string
}

func (this *User) Close() {
        fmt.Println("\n", this.username, "Closed !!!")
}


func print(z int) {
       fmt.Println("print z = ", z)
}

func deferRet(x,y int) (int){
       z := 0
       defer print(z) // defer后面只能跟函数？？？
       z = x + y
       fmt.Println("deferRet z = ", z)
       return z + 50 
}

func main() {
        fmt.Println("print name:")
        u1 := &User{"jack"}
        defer u1.Close()
        u2 := &User{"lily"}
        defer u2.Close()

        time.Sleep(10 * time.Second)  // 实际上，线程Sleep的10秒，u1，和u2早就可以Close()了，但却需要依赖main()函数的结束，才能defer执行。

        fmt.Println("main Done !")

        // 函数中的defer
        res := deferRet(1,1)
        fmt.Println("main res = ", res)

        // 我们可以在官方的文档中看到defer的执行顺序是逆序的，也就是先进后出的顺序：
        // 打印结果是：4,3,2,1,0
        fmt.Printf("print i:")
        for i := 0; i < 5; i++ {
            defer fmt.Printf("%d ", i)
        }
        fmt.Println("")
}
