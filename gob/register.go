package pifxregister

import (
	"encoding/gob"

	pifx "github.com/andrewmunsell/pifx-go-lib"
	easing "github.com/andrewmunsell/pifx-go-lib/easing"
	animations "github.com/andrewmunsell/pifx-go-lib/animations"
)

func RegisterGobTypes() {
	gob.Register(&pifx.Pixel{})
	gob.Register([]*pifx.Pixel{})

	gob.Register([]animations.Animation{})

	gob.Register(&animations.ColorWipe{})
	gob.Register(&animations.ColorWheel{})
	gob.Register(&animations.Chase{})

	gob.Register([]easing.Easing{})

	gob.Register(&easing.EaseInOutCubic{})
}