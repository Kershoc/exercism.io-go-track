package allyourbase

import (
	"errors"
	"math"
)

//ConvertToBase converts digits in base a to digits in base b
func ConvertToBase(a int, digits []int, b int) ([]int, error) {
	o := []int{}
	if a < 2 {
		return o, errors.New("input base must be >= 2")
	}
	if b < 2 {
		return o, errors.New("output base must be >= 2")
	}
	n := 0
	p := float64(len(digits) - 1)
	for _, d := range digits {
		if d < 0 || d >= a {
			return o, errors.New("all digits must satisfy 0 <= d < input base")
		}
		n += d * int(math.Pow(float64(a), p))
		p--
	}
	for n > 0 {
		o = append([]int{n % b}, o...)
		n /= b
	}
	if len(o) == 0 {
		o = []int{0}
	}

	return o, nil
}
