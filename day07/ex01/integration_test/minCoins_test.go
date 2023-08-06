package integration_test

import (
	app "day07/ex01/app"
	"sort"
	"testing"
)

func BenchmarkOne(t *testing.B) {
	tmp := []int{1, 2, 2, 4, 12, 13, 144}
	sort.Ints(tmp)
	var res []int
	for i := 0; i < t.N; i++ {
		res = app.MinCoins(14, tmp)
	}
	exp := []int{13, 1}
	for i := 0; i != len(exp); i++ {
		if res[i] != exp[i] {
			t.Errorf("Expect %d, but res = %d", exp, res)
			break
		}
	}
}

func BenchmarkTwo(t *testing.B) {
	res := app.MinCoins(13, []int{1, 5, 10})
	exp := []int{10, 1, 1, 1}
	for i := 0; i != len(exp); i++ {
		if res[i] != exp[i] {
			t.Errorf("Expect %d, but res = %d", exp, res)
			break
		}
	}
}

func BenchmarkThree(t *testing.B) {
	res := app.MinCoins(13, []int{1, 5, 10, 13})
	exp := []int{13}
	for i := 0; i != len(exp); i++ {
		if res[i] != exp[i] {
			t.Errorf("Expect %d, but res = %d", exp, res)
			break
		}
	}
}

func BenchmarkFour(t *testing.B) {
	res := app.MinCoins(8, []int{1, 2, 2, 2, 10, 13})
	exp := []int{2, 2, 2, 2}
	if len(res) != len(exp) {
		t.Errorf("Expect %d, but res = %d", exp, res)
	}
	for i := 0; i != len(exp); i++ {
		if res[i] != exp[i] {
			t.Errorf("Expect %d, but res = %d", exp, res)
			break
		}
	}
}

func BenchmarkFive(t *testing.B) {
	res := app.MinCoins(9, []int{11, 51, 101, 13})
	exp := []int{}
	if len(res) != len(exp) {
		t.Errorf("Expect %d, but res = %d", exp, res)
	}
	for i := 0; i != len(exp); i++ {
		if res[i] != exp[i] {
			t.Errorf("Expect %d, but res = %d", exp, res)
			break
		}
	}
}
