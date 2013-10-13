package pifx

import "fmt"

type PixelCommand struct {
	Offset  int
	Length  int
	Action  int         // 0 write pixel data, 1 wipe pixels, 2 change animations
	Payload interface{} // Data for the above action
}

func (c *PixelCommand) String() string {
	return fmt.Sprintf("%d: %v", c.Action, c.Payload)
}
