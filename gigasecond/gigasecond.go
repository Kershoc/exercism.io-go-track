// Package gigasecond should have a package comment that summarizes what it's about.
// It really should, but it is late and I am tired, so it doesn't.
package gigasecond

import (
	"math"
	"time"
)

// AddGigasecond adds a gigasecond to the given time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Duration(math.Pow10(9)) * time.Second)
}
