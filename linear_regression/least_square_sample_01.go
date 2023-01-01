/**
 * 最小二乗法サンプル
 * 参考 : イラストで学ぶ機械学習 Chapter3.1
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
	"math/rand"
//	"fmt"
)

func makeBaseFunction(a_m int) func(float64) []float64 {
	return func(a_x float64) []float64 {
		ret := make([]float64, 0, 2*a_m+1)
		for i := 0; i <= a_m; i++ {
			switch {
			case i == 0:
				ret = append(ret, 1)
			default:
				ret = append(ret, math.Sin(float64(i)*a_x/2), math.Cos(float64(i)*a_x/2))
			}
		}
		return ret
	}
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
	xys := make(plotter.XYs, length)
	for i := 0; i < length; i++ {
		xys[i].X = a_xVec[i]
		xys[i].Y = a_yVec[i]
	}
	plotutil.AddLinePoints(a_p, "f", xys)
}

func addPoints(a_p *plot.Plot, a_xVec, a_yVec []float64) {
	length := len(a_xVec)
	xyzs := make(plotter.XYZs, length)
	for i := 0; i < length; i++ {
		xyzs[i].X = a_xVec[i]
		xyzs[i].Y = a_yVec[i]
		xyzs[i].Z = 1
	}
	bs, _ := plotter.NewBubbles(xyzs, vg.Points(2), vg.Points(2))
	bs.Color = color.RGBA{R: 196, B: 128, A: 255}
	a_p.Add(bs)
}

func main() {
	// alias
	Dot := matrix.Product
	Inv := matrix.Inverse
	T := matrix.Transpose
	Sin := math.Sin
	Pi := math.Pi

	// rand
	r := rand.New(rand.NewSource(0))

	// train data
	train_data_size := 50
	vec_x := linspace(-3, 3, train_data_size)
	vec_t := make([]float64, 0, train_data_size)
	for _, x := range vec_x {
		vec_t = append(vec_t, Sin(Pi*x)/Pi*x+0.1*x+0.05*r.NormFloat64())
	}

	// base function
	φ := makeBaseFunction(5) // ここ，例は 15 だけど，15 にすると上手くフィットしない

	// estimate
	Φ := matrix.MakeDenseMatrixStacked(makePhiMatrix(vec_x, φ))
	w := Dot(Inv(Dot(T(Φ), Φ)), Dot(T(Φ), matrix.MakeDenseMatrix(vec_t, len(vec_t), 1)))

	// 求めた重みでグラフを描いてみる
	xlist := linspace(-3, 3, 1000)
	ylist := make([]float64, 0, 1000)
	for _, x := range xlist {
		ylist = append(ylist, f(w.Array(), x, φ))
	}

	// 描画
	p, _ := plot.New()
	p.Title.Text = "Linear regression"
	p.X.Label.Text = "X"
	p.Y.Label