package pifx

import (
	"os"
	"strings"
)

/**
 * Pixel strand struct
 */
type Strand struct {
	Pixels []Pixel
	Device string

	spi *os.File

	modified bool

	bytes []byte
}

/**
 * Initialize the pixel strand by getting the SPI device file
 * descriptor and initializing the pixel slice with a capacity
 * specified.
 */
func (s *Strand) Init(pixels int, device string) {
	s.Pixels = make([]Pixel, pixels)
	s.Device = device

	f, err := os.OpenFile(s.Device, os.O_RDWR, 0600)

	if err != nil {
		panic(err)
	}

	s.bytes = make([]byte, pixels*3)

	s.modified = true

	s.spi = f
}

/**
 * Set a specific pixel with an offset to a color
 */
func (s *Strand) Set(offset int, pixel Pixel) {
	if offset >= len(s.Pixels) {
		panic("Index out of bounds of pixel array.")
	}

	s.Pixels[offset].R = pixel.R
	s.Pixels[offset].G = pixel.G
	s.Pixels[offset].B = pixel.B

	s.modified = true
}

func (s *Strand) SetBytes(buf []byte) {
	s.bytes = buf
}

func (s *Strand) Wipe(color Pixel) {
	for i := range s.Pixels {
		s.Set(i, color)
	}
}

/**
 * Write the pixels to the SPI device
 */
func (s *Strand) Write() {
	if !s.modified {
		return
	}

	s.modified = false

	if _, err := s.spi.Write(*s.Bytes()); err != nil {
		panic(err)
	}
}

func (s *Strand) Bytes() *[]byte {
	for i, pixel := range s.Pixels {
		offset := i * 3

		s.bytes[offset] = pixel.R
		s.bytes[offset+1] = pixel.G
		s.bytes[offset+2] = pixel.B
	}

	return &s.bytes
}

/**
 * Return string formatted strand
 */
func (s *Strand) String() string {
	output := "{"

	for _, pixel := range s.Pixels {
		output += pixel.String() + ", "
	}

	output = strings.TrimRight(output, ", ") + "}"

	return output
}

/**
 * Create a new strand
 */
func NewStrand(pixels int, device string) Strand {
	strand := Strand{}

	strand.Init(pixels, device)

	return strand
}
