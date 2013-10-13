package animations

import (
	"time"

	pifx "github.com/andrewmunsell/pifx-go-lib"
)

type ColorWipe struct {
	Offset int
	Length int

	Color *pifx.Pixel
}

func (a *ColorWipe) Render(time time.Time, strand *pifx.Strand) {
	if a.Offset == 0 && a.Length == -1 {
		strand.Wipe(*a.Color)
	} else {
		if a.Length < 0 {
			a.Length = len(strand.Pixels) - a.Offset
		}

		for i := a.Offset; i < a.Offset+a.Length; i++ {
			strand.Set(i, *a.Color)
		}
	}
}
