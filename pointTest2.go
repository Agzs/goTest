package main

import "fmt"

const MAX int = 3

func main() {
   arr := []int{10,100,200}
   var i int
   var ptrr [MAX]*int;

   for  i = 0; i < MAX; i++ {
      ptrr[i] = &arr[i] /* 整数地址赋值给指针数组 */
   }

   for  i = 0; i < MAX; i++ {
      fmt.Printf("arr[%d] = %d\n", i,*ptrr[i] )
   }
   fmt.Printf("----------------\n")
   
   var aa int
   var ptr *int
   var pptr **int

   aa = 3000

   /* 指针 ptr 地址 */
   ptr = &aa

   /* 指向指针 ptr 地址 */
   pptr = &ptr

   /* 获取 pptr 的值 */
   fmt.Printf("变量 aa = %d\n", aa )
   fmt.Printf("指针变量 *ptr = %d\n", *ptr )
   fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)
   fmt.Printf("----------------\n")

}
