/*
声明数组
Go 语言数组声明需要指定元素类型及元素个数，语法格式如下：
var variable_name [SIZE] variable_type
以上为一维数组的定义方式。数组长度必须是整数且大于 0。

例如以下定义了数组 balance 长度为 10 类型为 float32：
var balance [10] float32

初始化数组
以下演示了数组初始化：
var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
初始化数组中 {} 中的元素个数不能大于 [] 中的数字。

如果忽略 [] 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小：
 var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
or
  balance := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0} 

Go 语言支持多维数组，以下为常用的多维数组声明方式：
var variable_name [SIZE1][SIZE2]...[SIZEN] variable_type

*/
package main

import "fmt"

func main() {
   var n [10]int /* n 是一个长度为 10 的数组 */
   var a = [5][2]int{ {0,0}, {1,2}, {2,4}, {3,6},{4,8}} /* 数组 - 5 行 2 列*/
 
   var i,j int

   /* 为数组 n 初始化元素 */         
   for i = 0; i < 10; i++ {
      n[i] = i + 100 /* 设置元素为 i + 100 */
   }

   /* 输出每个数组元素的值 */
   for j = 0; j < 10; j++ {
      fmt.Printf("Element[%d] = %d\n", j, n[j] )
   }
   /* 输出数组元素 */
   for  i = 0; i < 5; i++ {
      for j = 0; j < 2; j++ {
         fmt.Printf("a[%d][%d] = %d\n", i,j, a[i][j] )
      }
   }
   //--------------------------------------------
    /* 数组长度为 5 */
   var  balance = []int {1000, 2, 3, 17, 50}
   var avg float32

   /* 数组作为参数传递给函数 */
   avg = getAverage( balance, 5 ) ;

   /* 输出返回的平均值 */
   fmt.Printf( "平均值为: %f ", avg );
}
func getAverage(arr []int, size int) float32 {
   var i,sum int
   var avg float32  

   for i = 0; i < size;i++ {
      sum += arr[i]
   }

   avg = float32(sum / size)

   return avg;
}


