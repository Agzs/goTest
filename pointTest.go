/*
指针使用流程：
定义指针变量。
为指针变量赋值。
访问指针变量中指向地址的值。
在指针类型前面加上 * 号（前缀）来获取指针所指向的内容。

当一个指针被定义后没有分配到任何变量时，它的值为 nil。
nil 指针也称为空指针。
nil在概念上和其它语言的null、None、nil、NULL一样，都指代零值或空值。
一个指针变量通常缩写为 ptr。
*/

package main

import "fmt"

func main() {
   var a int = 4
   var b int32
   var c float32
   
   /* 运算符实例 */
   fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a );
   fmt.Printf("第 2 行 - b 变量类型为 = %T\n", b );
   fmt.Printf("第 3 行 - c 变量类型为 = %T\n", c );

   /*  & 和 * 运算符实例 */
   
   fmt.Printf("a 的值为  %d\n", a);

   //-----------------------------------
   var ptr *int
   ptr = &a	/* 'ptr' 包含了 'a' 变量的地址 */
   fmt.Printf("*ptr 为 %d\n", *ptr);
   if(ptr != nil){     /* ptr 不是空指针 */
      fmt.Printf("ptr != nil ,值为 : %x\n", ptr  )
   }
   if(ptr == nil){    /* ptr 是空指针 */
      fmt.Printf("ptr == nil ,值为 : %x\n", ptr  )
   }
   var ptr1 *int   
   if(ptr1 == nil){    /* ptr 是空指针 */
      fmt.Printf("ptr1 == nil ,值为 : %x\n", ptr1  )
   }
   
   //------------------------------------------------
   var aa int= 20   /* 声明实际变量 */
   var ip *int        /* 声明指针变量 */

   ip = &aa  /* 指针变量的存储地址 */

   fmt.Printf("aa 变量的地址是: %x\n", &aa  )

   /* 指针变量的存储地址 */
   fmt.Printf("ip 变量储存的指针地址: %x\n", ip )

   /* 使用指针访问值 */
   fmt.Printf("*ip 变量的值: %d\n", *ip )

}
