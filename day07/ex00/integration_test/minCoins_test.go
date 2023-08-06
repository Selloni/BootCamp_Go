package integration_test

import (
	app2 "day07/ex00/app"
	"sort"
	"testing"
)

func TestOne(t *testing.T) {
	tmp := []int{1, 2, 144, 12, 13, 2, 4}
	sort.Ints(tmp)
	res := app2.MinCoins(14, tmp)
	exp := []int{13, 1}
	for i := 0; i != len(exp); i++ {
		if res[i] != exp[i] {
			t.Errorf("Expect %d, but res = %d", exp, res)
			break
		}
	}
}

func TestTwo(t *testing.T) {
	res := app2.MinCoins(13, []int{1, 5, 10})
	exp := []int{10, 1, 1, 1}
	for i := 0; i != len(exp); i++ {
		if res[i] != exp[i] {
			t.Errorf("Expect %d, but res = %d", exp, res)
			break
		}
	}
}

func TestThree(t *testing.T) {
	res := app2.MinCoins(13, []int{1, 5, 10, 13})
	exp := []int{13}
	for i := 0; i != len(exp); i++ {
		if res[i] != exp[i] {
			t.Errorf("Expect %d, but res = %d", exp, res)
			break
		}
	}
}

func TestFour(t *testing.T) {
	res := app2.MinCoins(8, []int{1, 2, 2, 2, 10, 13})
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

func TestFive(t *testing.T) {
	res := app2.MinCoins(9, []int{11, 51, 101, 13})
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

func TestFalling(t *testing.T) {
	res := app2.MinCoins(12, []int{2, 5, 10, 2})
	exp := []int{10, 2}
	for i := 0; i != len(exp); i++ {
		if res[i] != exp[i] {
			t.Errorf("Expect %d, but res = %d", exp, res)
			break
		}
	}
}

func TestMyOne(t *testing.T) {
	tmp := []int{1, 2, 144, 12, 13, 2, 4}
	sort.Ints(tmp)
	res := app2.MinCoins2(14, tmp)
	exp := []int{13, 1}
	for i := 0; i != len(exp); i++ {
		if res[i] != exp[i] {
			t.Errorf("Expect %d, but res = %d", exp, res)
			break
		}
	}
}

func TestMyTwo(t *testing.T) {
	res := app2.MinCoins2(13, []int{1, 5, 10})
	exp := []int{10, 1, 1, 1}
	for i := 0; i != len(exp); i++ {
		if res[i] != exp[i] {
			t.Errorf("Expect %d, but res = %d", exp, res)
			break
		}
	}
}

func TestMyThree(t *testing.T) {
	res := app2.MinCoins2(13, []int{1, 5, 10, 13})
	exp := []int{13}
	for i := 0; i != len(exp); i++ {
		if res[i] != exp[i] {
			t.Errorf("Expect %d, but res = %d", exp, res)
			break
		}
	}
}

func TestMyFour(t *testing.T) {
	res := app2.MinCoins2(8, []int{1, 2, 2, 2, 10, 13})
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

func TestMyFive(t *testing.T) {
	res := app2.MinCoins2(9, []int{11, 51, 101, 13})
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

func TestMyFalling(t *testing.T) {
	res := app2.MinCoins2(12, []int{2, 5, 10, 2})
	exp := []int{10, 2}
	for i := 0; i != len(exp); i++ {
		if res[i] != exp[i] {
			t.Errorf("Expect %d, but res = %d", exp, res)
			break
		}
	}
}
