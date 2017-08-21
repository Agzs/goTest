// goSHA256 project main.go
package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	var s1 string = "A.com"
	var s2 string = "a.com"
	var s3 string = "0.com"

	printHash(26, s1)
	printHash(26, s2)
	printHash(10, s3)

	//	c = getHmacCode("hello world")
	//	fmt.Println(c)
}

func printHash(n int, s string) {
	var x int
	switch {
	case s == "A.com":
		x = 65 // 'A'= 65
	case s == "a.com":
		x = 97
	case s == "0.com":
		x = 48
	default:
		fmt.Printf("error\n")
	}
	for i := 0; i < n; i++ {
		sByte := []byte(s) //转换字符串的内容，先转换s的类型为[]byte
		t := x + i
		sByte[0] = byte(t) //赋值
		ss := string(sByte)
		//fmt.Println(ss)
		c := getSha256Code(ss)
		fmt.Println(c)
	}
}

func getSha256Code(s string) string {
	fmt.Printf("SHA256(%s) = ", s)
	h := sha256.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func getHmacCode(s string) string {
	h := hmac.New(sha256.New, []byte("ourkey"))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}
