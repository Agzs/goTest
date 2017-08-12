/*
定义 Map
可以使用内建函数 make 也可以使用 map 关键字来定义 Map:
var map_variable map[key_data_type]value_data_type  //声明变量，默认 map 是 nil
map_variable = make(map[key_data_type]value_data_type)  // 使用 make 函数 

如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对

delete() 函数用于删除集合的元素, 参数为 map 和其对应的 key
*/

package main

import "fmt"

func main() {
   var countryCapitalMap map[string]string
   /* 创建集合 */
   countryCapitalMap = make(map[string]string)
   
   /* map 插入 key-value 对，各个国家对应的首都 */
   countryCapitalMap["France"] = "Paris"
   countryCapitalMap["Italy"] = "Rome"
   countryCapitalMap["Japan"] = "Tokyo"
   countryCapitalMap["India"] = "New Delhi"
   
   /* 使用 key 输出 map 值 */
   for country := range countryCapitalMap {
      fmt.Println("Capital of",country,"is",countryCapitalMap[country])
   }
   
   /* 查看元素在集合中是否存在 */
   captial, ok := countryCapitalMap["United States"]
   /* 如果 ok 是 true, 则存在，否则不存在 */
   if(ok){
      fmt.Println("Capital of United States is", captial)  
   }else {
      fmt.Println("Capital of United States is not present") 
   }
   captial, ok = countryCapitalMap["France"]
   if(ok){
      delete(countryCapitalMap,"France")
      fmt.Println("France is deleted")
   } else {
      fmt.Println("France is not present") 
   }
   /* 使用 key, value 输出 map 值 */
   for k,v := range countryCapitalMap {
      fmt.Println(k," -> ",v)
   }
}

