package pifx

import "fmt"

type PixelCommand struct {
	Offset  int
	Action  int         // 0 write pixel, 1 change animation
	Payload interface{} // Data for the above action
}

func (c *PixelCommand) String() string {
	return fmt.Sprintf("%d: %v", c.Action, c.Payload)
}
