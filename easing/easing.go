package easing

type Easing interface {
	Value(progress float64) float64
}
