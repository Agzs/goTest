/*
Go 语言接口
Go 语言提供了另外一种数据类型即接口，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。
实例
// 定义接口 
type interface_name interface {
   method_name1 [return_type]
   method_name2 [return_type]
   method_name3 [return_type]
   ...
   method_namen [return_type]
}

// 定义结构体 
type struct_name struct {
   // variables 
}

//实现接口方法 
func (struct_name_variable struct_name) method_name1() [return_type] {
   // 方法实现 
}
...
func (struct_name_variable struct_name) method_namen() [return_type] {
   // 方法实现
}
*/

/*
例子中，我们定义了一个接口Phone，接口里面有一个方法call()。
然后我们在main函数里面定义了一个Phone类型变量，并分别为之赋值为NokiaPhone和IPhone。然后调用call()方法.
*/
package main

import (
    "fmt"
)

type Phone interface {
    call()
    print()
}

type NokiaPhone struct {
    name string
    number int
    price  int
}

func (nokiaPhone NokiaPhone) call() {
    fmt.Println("I am Nokia, I can call you!")
}
func (nokiaPhone NokiaPhone) print() {
    fmt.Println("name:",nokiaPhone.name,"number:",nokiaPhone.number,"price:",nokiaPhone.price)
}

type IPhone struct {
    name string
    number int
}

func (iPhone IPhone) call() {
    fmt.Println("I am iPhone, I can call you!")
}
func (iPhone IPhone) print() {
     fmt.Println("name:",iPhone.name,"number:",iPhone.number)
}

func main() {
    var phone Phone

    //phone = new(NokiaPhone)
   
    phone = &NokiaPhone{"Nokia",150300,2000}
    phone.call()
    phone.print()

    //phone = new(IPhone)
    phone = &IPhone{"iphone",153456}
    phone.call()
    phone.print()

}
