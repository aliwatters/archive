// Mandelbrot emits a PNG image of the Mandelbrot fractal
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 2048, 2048
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		y2 := float64(py+1)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			x2 := float64(px+1)/width*(xmax-xmin) + xmin
			z0 := complex(x, y)
			z1 := complex(x2, y)
			z2 := complex(x, y2)
			z3 := complex(x2, y2)

			// supersampling ... 4 offset complex numbers
			z := (z0 + z1 + z2 + z3) / 4
			// Image point (px, py) represents complex value z
			img.Set(px, py, mandelbrot(z))
		}
	}

	png.Encode(os.Stdout, img) // Note: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		c := cmplx.Abs(v)
		if c > 2 {
			// return color.Gray{255 - contrast - n}
			return color.RGBA{
				(n * 5) % 255,
				(n * 10) % 255,
				(n * 15) % 255,
				255,
			}
		}
	}
	return color.Black
}
