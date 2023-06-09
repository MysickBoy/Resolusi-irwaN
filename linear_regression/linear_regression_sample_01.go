
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
     	// train data
	vec_x := []float64{0.02, 0.12, 0.19, 0.27, 0.42, 0.51, 0.64, 0.84, 0.88, 0.99}
	vec_t := []float64{0.05, 0.87, 0.94, 0.92, 0.54, -0.11, -0.78, -0.89, -0.79, -0.04}

	// base function
	φ:= func(a_x float64) []float64 {
		return []float64{1, a_x, math.Pow(a_x, 2), math.Pow(a_x, 3), math.Pow(a_x, 4)}
	}
	// φ = defaultBaseFunction

	Φ := matrix.MakeDenseMatrixStacked(makePhiMatrix(vec_x, φ))
	w := Dot(Inv(Dot(T(Φ), Φ)), Dot(T(Φ), matrix.MakeDenseMatrix(vec_t, 10, 1)))

        // 求めた重みでグラフを描いてみる
	xlist := linspace(0, 1, 100)
	ylist := make([]float64, 0)
	for _, x := range xlist {
		ylist = append(ylist, f(w.Array(), x, φ))
	}

	// 描画
	p, _ := plot.New()
	p.Title.Text = "Linear regression"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	addLine(p, xlist, ylist)
	addPoints(p, vec_x, vec_t)
	p.Save(4, 4, "linear_regression_sample_01.png")
}