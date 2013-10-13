package pifx

import (
	"encoding/binary"
	"reflect"
)

func IntToByte(number int) byte {
	i := make([]int, 1)
	i[0] = number

	intSize := int(reflect.TypeOf(i).Elem().Size())
	b := make([]byte, intSize*len(i))
	for n, s := range i {
		switch intSize {
		case 64 / 8:
			binary.BigEndian.PutUint64(b[intSize*n:], uint64(s))
		case 32 / 8:
			binary.BigEndian.PutUint32(b[intSize*n:], uint32(s))
		default:
			panic("unreachable")
		}
	}

	return b[intSize-1]
}

func NewPixelHSL(h, s, l float64) *Pixel {
	r, g, b := HSLToRGB(h, s, l)

	return &Pixel{r, g, b}
}

func HSLToRGB(h, s, l float64) (r, g, b uint8) {
	var fR, fG, fB float64
	if s == 0 {
		fR, fG, fB = l, l, l
	} else {
		var q float64
		if l < 0.5 {
			q = l * (1 + s)
		} else {
			q = l + s - s*l
		}
		p := 2*l - q
		fR = hueToRGB(p, q, h+1.0/3)
		fG = hueToRGB(p, q, h)
		fB = hueToRGB(p, q, h-1.0/3)
	}
	r = uint8((fR * 255) + 0.5)
	g = uint8((fG * 255) + 0.5)
	b = uint8((fB * 255) + 0.5)
	return
}

// hueToRGB is a helper function for HSLToRGB.
func hueToRGB(p, q, t float64) float64 {
	if t < 0 {
		t += 1
	}
	if t > 1 {
		t -= 1
	}
	if t < 1.0/6 {
		return p + (q-p)*6*t
	}
	if t < 0.5 {
		return q
	}
	if t < 2.0/3 {
		return p + (q-p)*(2.0/3-t)*6
	}
	return p
}
