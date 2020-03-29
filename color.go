package colorpicker

import (
	"image/color"
	"math"
)

func fromHSV(h, s, v float64) *color.RGBA {
	if s == 0 {
		return &color.RGBA{
			R: uint8(v * 255),
			G: uint8(v * 255),
			B: uint8(v * 255),
			A: 0xff,
		}
	}

	h = h * 6
	if h == 6 {
		h = 0
	}
	i := math.Floor(h)
	v1 := v * (1 - s)
	v2 := v * (1 - s*(h-i))
	v3 := v * (1 - s*(1-(h-i)))

	switch int(i) {
	case 0:
		return &color.RGBA{
			R: uint8(v * 255),
			G: uint8(v3 * 255),
			B: uint8(v1 * 255),
			A: 0xff,
		}
	case 1:
		return &color.RGBA{
			R: uint8(v2 * 255),
			G: uint8(v * 255),
			B: uint8(v1 * 255),
			A: 0xff,
		}
	case 2:
		return &color.RGBA{
			R: uint8(v1 * 255),
			G: uint8(v * 255),
			B: uint8(v3 * 255),
			A: 0xff,
		}
	case 3:
		return &color.RGBA{
			R: uint8(v1 * 255),
			G: uint8(v2 * 255),
			B: uint8(v * 255),
			A: 0xff,
		}
	case 4:
		return &color.RGBA{
			R: uint8(v3 * 255),
			G: uint8(v1 * 255),
			B: uint8(v * 255),
			A: 0xff,
		}
	default:
		return &color.RGBA{
			R: uint8(v * 255),
			G: uint8(v1 * 255),
			B: uint8(v2 * 255),
			A: 0xff,
		}
	}
}
