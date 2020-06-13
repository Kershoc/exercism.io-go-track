package grains

import (
	"errors"
)

//Square returns how many grains are on a given square
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("input out of range")
	}
	return 1 << (n - 1), nil
}

//Total return total number of grains on a chessboard
func Total() uint64 {
	return 1<<64 - 1
}
