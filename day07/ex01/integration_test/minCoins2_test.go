package integration_test

//
//import (
//	"day07/ex01/app"
//	"testing"
//)
//
//func BenchmarkMyOne(b *testing.B) {
//	tmp := []int{1, 2, 2, 4, 12, 13, 144}
//	var res []int
//	for i := 0; i < b.N; i++ {
//		res = app.MinCoins2(14, tmp)
//	}
//	exp := []int{13, 1}
//	for i := 0; i != len(exp); i++ {
//		if res[i] != exp[i] {
//			b.Errorf("Expect %d, but res = %d", exp, res)
//			break
//		}
//	}
//}
