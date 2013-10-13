package animations

import (
	"time"

	pifx "github.com/andrewmunsell/pifx-go-lib"
)

type Animation interface {
	Render(time time.Time, strand *pifx.Strand)
}
