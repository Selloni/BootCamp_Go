package integration_test

import (
	app2 "day07/ex01/app"
	"testing"
)

func BenchmarkMyOne(b *testing.B) {
	tmp := []int{1, 2, 2, 4, 12, 13, 144}
	res := app2.MinCoins2(14, tmp)
	exp := []int{13, 1}
	for i := 0; i != len(exp); i++ {
		if res[i] != exp[i] {
			b.Errorf("Expect %d, but res = %d", exp, res)
			break
		}
	}
}

func BenchmarkMyTwo(b *testing.B) {
	res := app2.MinCoins2(13, []int{1, 5, 10})
	exp := []int{10, 1, 1, 1}
	for i := 0; i != len(exp); i++ {
		if res[i] != exp[i] {
			b.Errorf("Expect %d, but res = %d", exp, res)
			break
		}
	}
}

func BenchmarkMyThree(b *testing.B) {
	res := app2.MinCoins2(13, []int{1, 5, 10, 13})
	exp := []int{13}
	for i := 0; i != len(exp); i++ {
		if res[i] != exp[i] {
			b.Errorf("Expect %d, but res = %d", exp, res)
			break
		}
	}
}

func BenchmarkMyFour(b *testing.B) {
	res := app2.MinCoins2(8, []int{1, 2, 2, 2, 10, 13})
	exp := []int{2, 2, 2, 2}
	if len(res) != len(exp) {
		b.Errorf("Expect %d, but res = %d", exp, res)
	}
	for i := 0; i != len(exp); i++ {
		if res[i] != exp[i] {
			b.Errorf("Expect %d, but res = %d", exp, res)
			break
		}
	}
}

func BenchmarkMyFive(b *testing.B) {
	res := app2.MinCoins2(9, []int{11, 51, 101, 13})
	exp := []int{}
	if len(res) != len(exp) {
		b.Errorf("Expect %d, but res = %d", exp, res)
	}
	for i := 0; i != len(exp); i++ {
		if res[i] != exp[i] {
			b.Errorf("Expect %d, but res = %d", exp, res)
			break
		}
	}
}

func BenchmarkMyFalling(b *testing.B) {
	res := app2.MinCoins2(12, []int{2, 5, 10, 2})
	exp := []int{10, 2}
	for i := 0; i != len(exp); i++ {
		if res[i] != exp[i] {
			b.Errorf("Expect %d, but res = %d", exp, res)
			break
		}
	}
}
