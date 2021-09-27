package plain

import (
	"github.com/fogleman/gg"
	"math"
)

/*Dibuja colores hasta W y H usando suma para red(izq->der) y green(arr->abj) y resta para blue(
izq+arr->der+abj)*/
func DrawBasic(w, h float64) *gg.Context {
	dc := gg.NewContext(int(w), int(h))

	for x := 0.0; x < w; x++ {
		for y := 0.0; y < h; y++ {

			r := addColor(x, w)
			g := addColor(y, h)
			b := subtractColor(x, w)

			dc.SetRGB(r, g, b)
			dc.SetPixel(int(x), int(y))

		}
	}
	return dc
}

/*Dibuja colores hasta W y H usando Sin para red (izq->der) y green(arr->abj) y Cos para blue(
izq+arr->der+abj)*/
func DrawSin(w, h float64) *gg.Context {
	dc := gg.NewContext(int(w), int(h))

	for x := 0.0; x < w; x++ {
		for y := 0.0; y < h; y++ {

			r := getSinColor(x, w)
			g := getSinColor(y, h)
			b := getCosColor(x+y, w+h)

			dc.SetRGB(r, g, b)
			dc.SetPixel(int(x), int(y))

		}

	}
	return dc
}

func addColor(current, max float64) float64 {
	return current / max
}
func subtractColor(current, max float64) float64 {
	return 1.0 - (current / max)
}

const straightAngle = 90

func getSinColor(current, max float64) float64 {
	radians := gg.Radians((current / max) * straightAngle)
	return math.Sin(radians)
}
func getCosColor(current, max float64) float64 {
	radians := gg.Radians((current / max) * straightAngle)
	return math.Cos(radians)
}
