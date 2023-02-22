
/**
 * 線形回帰サンプル
 * 参考 : http://gihyo.jp/dev/serial/01/machine-learning/0011?page=1&ard=1400930362 
 */

package main

import (
	"code.google.com/p/plotinum/plot"
	"code.google.com/p/plotinum/plotter"
	"code.google.com/p/plotinum/plotutil"
	"code.google.com/p/plotinum/vg"
	"github.com/skelterjohn/go.matrix"
	"image/color"
	"math"
)

func defaultBaseFunction(a_x float64) []float64 {
	return []float64{1, a_x, math.Pow(a_x, 2), math.Pow(a_x, 3)}
}

func makePhiMatrix(a_vec []float64, a_baseFunction func(float64) []float64) (matrix [][]float64) {
	matrix = make([][]float64, 0)
	for _, x := range a_vec {
		matrix = append(matrix, a_baseFunction(x))
	}
	return
}

func f(a_w []float64, a_x float64, a_baseFunction func(float64) []float64) float64 {
	vecW := matrix.MakeDenseMatrix(a_w, 1, len(a_w))
	phiX := a_baseFunction(a_x)
	vecPhiX := matrix.MakeDenseMatrix(phiX, len(phiX), 1)
	return matrix.Product(vecW, vecPhiX).Get(0, 0)
}

func linspace(a_start, a_end float64, a_n int) (ret []float64) {
	ret = make([]float64, a_n)
	if a_n == 1 {
		ret[0] = a_end
		return ret
	}
	delta := (a_end - a_start) / (float64(a_n) - 1)
	for i := 0; i < a_n; i++ {
		ret[i] = float64(a_start) + (delta * float64(i))
	}
	return
}

func addLine(a_p *plot.Plot, a_xVec, a_yVec []float64) {
	length := len(a_xVec)