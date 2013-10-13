package pifx

import "fmt"

/**
 * Struct representing a single LED pixel
 */
type Pixel struct {
	R, G, B byte
}

/**
 * Return a string representation of the pixel object
 */
func (p *Pixel) String() (string){
	return fmt.Sprintf("rgb(%d, %d, %d)", p.R, p.G, p.B)
}

func NewPixel(r, g, b byte) *Pixel{
	return &Pixel{
		R: r,
		G: g,
		B: b,
	}
}