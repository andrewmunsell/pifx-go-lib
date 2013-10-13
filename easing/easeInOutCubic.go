package easing

type EaseInOutCubic struct{

}

func (e *EaseInOutCubic) Value(progress float64) float64 {
	if progress < .5 {
		return 4 * progress * progress * progress 
 	} else {
		return (progress - 1) * (2 * progress - 2) * (2 * progress - 2) + 1
	}
}