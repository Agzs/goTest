/*
Go 语言切片是对数组的抽象。
Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

定义切片
声明一个未指定大小的数组来定义切片：
var identifier []type  切片不需要说明长度。

或使用make()函数来创建切片:
var slice1 []type = make([]type, len)
or
slice1 := make([]type, len)

也可以指定容量，其中capacity为可选参数。
make([]T, length, capacity)这里 len 是数组的长度并且也是切片的初始长度。

切片初始化
s :=[] int {1,2,3 } 直接初始化切片，[]表示是切片类型，{1,2,3}初始化值依次是1,2,3.其cap=len=3
s := arr[:] 初始化切片s,是数组arr的引用
s := arr[startIndex:endIndex] 将arr中从下标startIndex到endIndex-1 下的元素创建为一个新的切片
s := arr[startIndex:] 缺省endIndex时将表示一直到arr的最后一个元素
s := arr[:endIndex] 缺省startIndex时将表示从arr的第一个元素开始
s1 := s[startIndex:endIndex] 通过切片s初始化切片s1
s :=make([]int,len,cap) 通过内置函数make()初始化切片s,[]int 标识为其元素类型为int的切片

切片是可索引的，并且可以由 len() 方法获取长度。
切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少。
拷贝切片的 copy 方法和向切片追加新元素的 append 方法。
*/
package main

import "fmt"

func main() {
   /* 创建切片 */
   numbers := []int{0,1,2,3,4,5,6,7,8}   
   printSlice(numbers)

   /* 打印原始切片 */
   fmt.Println("numbers ==", numbers)

   /* 打印子切片从索引1(包含) 到索引4(不包含)*/
   fmt.Println("numbers[1:4] ==", numbers[1:4])

   /* 默认下限为 0*/
   fmt.Println("numbers[:3] ==", numbers[:3])

   /* 默认上限为 len(s)*/
   fmt.Println("numbers[4:] ==", numbers[4:])

   numbers1 := make([]int,0,5)
   printSlice(numbers1)

   /* 打印子切片从索引  0(包含) 到索引 2(不包含) */
   number2 := numbers[:2]
   printSlice(number2)

   /* 打印子切片从索引 2(包含) 到索引 5(不包含) */
   number3 := numbers[2:5]
   printSlice(number3)

   fmt.Printf("------------\n")
   var nnumbers []int
   printSlice(nnumbers)

    /* 允许追加空切片 */
   nnumbers = append(nnumbers, 0)
   printSlice(nnumbers)

   /* 向切片添加一个元素 */
   nnumbers = append(nnumbers, 1)
   printSlice(nnumbers)

   /* 同时添加多个元素 */
   nnumbers = append(nnumbers, 2,3,4,5,6)
   printSlice(nnumbers)

   /* 创建切片 numbers1 是之前切片的两倍容量*/
   nnumbers1 := make([]int, len(nnumbers), (cap(nnumbers))*2)

   /* 拷贝 numbers 的内容到 numbers1 */
   copy(nnumbers1,nnumbers)
   printSlice(nnumbers1)   

}

func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

