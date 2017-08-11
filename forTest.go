package main

import "fmt"

func main() {

   var b int = 15
   var a int

   numbers := [6]int{1, 2, 3, 5}  // define array, size is 6

   /* for 循环 */
   for a := 0; a < 10; a++ {
      fmt.Printf("a 的值为: %d\n", a)
   }

   for a < b {    // like while
      a++
      fmt.Printf("a 的值为: %d\n", a)
      }

   for i,x:= range numbers {  //for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环
      fmt.Printf("第 %d 位 x 的值 = %d\n", i,x)
   }   
}
