
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
