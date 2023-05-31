package main

import (
	"testing"
)

func TestMedian1(t *testing.T) {
	var buff = []float64{-10003, 2032, 3, -75, 64}
	res, _ := median(buff)
	if res != 3 {
		t.Errorf("exepted %f != %f output", res, 3.0)
	}
}

func TestMedian2(t *testing.T) {
	var buff = []float64{-10003, 2032, 1, -173, -75, 64, 27, 27, 12, 1222}
	res1, res2 := median(buff)
	if res1 != 12 || res2 != 27 {
		t.Errorf("exepted %f||%f != %f,%f output ", res1, res2, 12.0, 27.0)
	}
}

func TestMean(t *testing.T) {
	var buff = []float64{-10003, 2032, 1, -173, -75, 64, 27, 27, 12, 1222}
	var sum float64
	for _, i := range buff {
		sum += i
	}
	itog := mean(sum, float64(len(buff)))
	if itog != -686.6 {
		t.Errorf("exrpted %f != %f output", itog, -686.6)
	}
}

func TestMode(t *testing.T) {
	var buff1 = []float64{-10003, 2032, 1, -173, -75, 64, 27, 27, 12, 1222}
	var buff2 = []float64{-10003, 2032, 12, -27, 1, 27, -173, -75, 64, 27, 12, 1222, 12, 27}
	value1 := mode(buff1)
	value2 := mode(buff2)
	if value1 != 27 {
		t.Errorf("exrpted %f != %f output", value1, 27.0)
	}
	if value2 != 12 {
		t.Errorf("exrpted %f != %f output", value1, 12.0)
	}
}
