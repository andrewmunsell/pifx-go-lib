package animations

import (
	"time"

	pifx "github.com/andrewmunsell/pifx-go-lib"
	easing "github.com/andrewmunsell/pifx-go-lib/easing"
)

type Chase struct {
	Color  *pifx.Pixel
	Speed  float64
	Easing easing.Easing

	lastFrame time.Time
	offset    float64
}

func (a *Chase) Render(time time.Time, strand *pifx.Strand) {
	max := len(strand.Pixels)

	for n := 0; n < max; n++ {
		opacity := (1 - (float64(n) / float64(max))) + a.offset

		for opacity > 1 {
			opacity -= 1
		}

		opacity = 1 - opacity

		opacity = a.Easing.Value(opacity)

		r := pifx.IntToByte(int(float64(a.Color.R) * opacity))
		g := pifx.IntToByte(int(float64(a.Color.G) * opacity))
		b := pifx.IntToByte(int(float64(a.Color.B) * opacity))

		strand.Set(n, *pifx.NewPixel(r, g, b))
	}

	a.offset += float64(time.UnixNano()-a.lastFrame.UnixNano()) / 1e10 * a.Speed

	if a.offset > 1 {
		a.offset = 0
	}

	if a.offset < 0 {
		a.offset = 1
	}

	a.lastFrame = time
}
