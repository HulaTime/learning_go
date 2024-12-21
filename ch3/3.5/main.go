// // Mandelbrot emits a PNG image of the Mandelbrot fractal in full color
// package main
//
// import (
//
//	"fmt"
//	"image"
//	"image/color"
//	"image/png"
//	"math/cmplx"
//	"os"
//
// )
//
//	func main() {
//		const (
//			xmin, ymin, xmax, ymax = -2, -2, +2, +2
//			width, height          = 1024, 1024
//		)
//
//		img := image.NewRGBA(image.Rect(0, 0, width, height))
//		for py := 0; py < height; py++ {
//			y := float64(py)/height*(ymax-ymin) + ymin
//			for px := 0; px < width; px++ {
//				x := float64(px)/width*(xmax-xmin) + xmin
//				z := complex(x, y)
//				// Image point (px, py) represents complex value z.
//				img.Set(px, py, mandelbrot(z))
//			}
//		}
//		png.Encode(os.Stdout, img) // NOTE: ignoring errors
//	}
//
//	func mandelbrot(z complex128) color.Color {
//		const iterations = 200
//		const contrast = 15
//		const rgb_boundary = iterations / 9
//		const rgb_steps = 255 / rgb_boundary
//
//		var v complex128
//		for n := uint8(0); n < iterations; n++ {
//			v = v*v + z
//			if cmplx.Abs(v) > 2 {
//				// 00-FF 00-FF 00-FF 00-FF
//				var blue, green, red uint8 = 0x00, 0x00, 0x00
//				if n <= uint8(rgb_boundary) {
//					blue += rgb_steps
//				}
//				if uint8(rgb_boundary) > n && uint16(n) <= rgb_boundary*2 {
//					green += rgb_steps
//					fmt.Fprintf(os.Stderr, "R: %d G: %d B: %d\n", red, green, blue)
//				}
//				if uint16(rgb_boundary*2) > uint16(n) && uint16(n) <= uint16(rgb_boundary*3) {
//					red += rgb_steps
//				}
//				return color.RGBA{red, green, blue, 0xFF}
//			}
//		}
//		return color.Black
//	}
//
// ex3.5 emits a full-color PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			switch {
			case n > 50: // dark red
				return color.RGBA{100, 0, 0, 255}
			default:
				// logarithmic blue gradient to show small differences on the
				// periphery of the fractal.
				logScale := math.Log(float64(n)) / math.Log(float64(iterations))
				return color.RGBA{0, 0, 255 - uint8(logScale*255), 255}
			}
		}
	}
	return color.Black
}

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//
//	= z - (z^4 - 1) / (4 * z^3)
//	= z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
