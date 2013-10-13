package animations

import (
	"time"

	pifx "github.com/andrewmunsell/pifx-go-lib"
)

type ColorWipe struct {
	Offset int
	Color  *pifx.Pixel
}

func (a *ColorWipe) Render(time time.Time, strand *pifx.Strand) {
	strand.Wipe(*a.Color)
}
