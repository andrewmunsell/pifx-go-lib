package animations

import (
	"time"

	pifx "github.com/andrewmunsell/pifx-go-lib"
)

type ColorWheel struct {
	Speed float64

	lastFrame time.Time
	offset    float64
}

func (a *ColorWheel) Render(time time.Time, strand *pifx.Strand) {
	max := len(strand.Pixels)

	for n := 0; n < max; n++ {
		hue := (float64(n) / float64(max)) + a.offset

		for hue > 1 {
			hue -= 1
		}

		strand.Set(n, *pifx.NewPixelHSL(hue, 1, 0.5))
	}

	a.offset += float64(time.UnixNano()-a.lastFrame.UnixNano()) / 1e10 * a.Speed

	if a.offset > 1 {
		a.offset = 0
	}

	a.lastFrame = time
}
