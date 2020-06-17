package clock

import (
	"fmt"
)

const dayInMinutes = 1440

//Clock type for keeping time
type Clock struct {
	timeInMinutes int
}

//New constructs new clocks
func New(h, m int) Clock {
	clock := Clock{h*60 + m}
	for clock.timeInMinutes < 0 {
		clock.timeInMinutes += dayInMinutes
	}
	clock.timeInMinutes %= dayInMinutes
	return clock
}

func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.timeInMinutes/60, clock.timeInMinutes%60)
}

//Add returns new clock with m added to current clock
func (clock Clock) Add(m int) Clock {
	return New(0, clock.timeInMinutes+m)
}

//Subtract returns new clock with m removed from current clock
func (clock Clock) Subtract(m int) Clock {
	return New(0, clock.timeInMinutes-m)
}
