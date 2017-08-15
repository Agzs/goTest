package fib

import (
	"testing"
)

func BenchmarkFibm(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibm(90)
	}
}
func TestFibm(t *testing.T) {
	ar := []int64{0, 1, 1, 2, 3, 5, 8, 13, 21}
	for i := 0; i < len(ar); i++ {
		if ar[i] != Fibm(int64(i)) {
			t.Fatalf("%d, %d != %d", i, ar[i], Fibm(int64(i)))
		}
	}
	ar1 := []int64{0, 1, -1, 2, -3, 5, -8, 13, -21}
	for i := 0; i < len(ar1); i++ {
		if ar1[i] != Fibm(int64(-i)) {
			t.Fatalf("%d, %d != %d", -i, ar1[i], Fibm(int64(-i)))
		}
	}
}
