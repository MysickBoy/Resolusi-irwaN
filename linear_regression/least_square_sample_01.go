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

func m