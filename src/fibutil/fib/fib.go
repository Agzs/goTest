// fib project fib.go
package fib

// 非递归方式实现Fibonacci算法
func Fibm(n int64) int64 {
	if n == 0 {
		return 0
	}
	var i, a, b int64
	b = 1
	if n < 0 {
		n = -n
		if n%2 == 0 {
			b = -1
		}
	}
	for i = 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// 递归方式实现Fibonacci算法
func Fibm2(n int64) int64 {
	if n < 0 {
		if n%2 == 0 {
			return _Fibm2(-n, 0, -1)
		} else {
			return _Fibm2(-n, 0, 1)
		}
	}
	return _Fibm2(n, 0, 1)
}

func _Fibm2(n, a1, a2 int64) int64 {
	if n == 0 {
		return a1
	}
	return _Fibm2(n-1, a2, a1+a2)
}
