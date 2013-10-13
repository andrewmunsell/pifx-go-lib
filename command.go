package pifx

import "fmt"

type PixelCommand struct {
	Action  int         // 0 write pixel data, 1 clear pixels, 2 append animation(s)
	Payload interface{} // Data for the above action
}

func (c *PixelCommand) String() string {
	return fmt.Sprintf("%d: %v", c.Action, c.Payload)
}
