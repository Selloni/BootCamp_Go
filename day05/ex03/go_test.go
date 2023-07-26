package main

import (
	present "day05/ex02/struct"
	"testing"
)

func TestStd(t *testing.T) {
	h := present.PresentHeap{
		present.Present{Value: 5, Size: 1},
		present.Present{Value: 4, Size: 5},
		present.Present{Value: 3, Size: 1},
		present.Present{Value: 5, Size: 2},
	}
	res := grabPresents(h, 2)
	if res != 8 {
		t.Error("expect a result of 8, but get: ", res)
	}
}

func TestMin(t *testing.T) {
	h := present.PresentHeap{
		present.Present{Value: 0, Size: 1},
		present.Present{Value: 4, Size: 5},
		present.Present{Value: 3, Size: 2},
		present.Present{Value: 5, Size: 2},
	}
	res := grabPresents(h, 1)
	if res != 0 {
		t.Error("expect a result of 0, but get: ", res)
	}
}

func TestMax(t *testing.T) {
	h := present.PresentHeap{
		present.Present{Value: 0, Size: 1},
		present.Present{Value: 4, Size: 5},
		present.Present{Value: 3, Size: 2},
		present.Present{Value: 5, Size: 2},
	}
	res := grabPresents(h, 10)
	if res != 12 {
		t.Error("expect a result of 0, but get: ", res)
	}
}
