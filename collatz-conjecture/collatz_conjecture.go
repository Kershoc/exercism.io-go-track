//Package collatzconjecture implements a formula to reduce any positive value to 1
//and count how long it takes to get there
package collatzconjecture

import "errors"

// CollatzConjecture computes number of iterations it takes to get to 1
func CollatzConjecture(i int) (int, error) {
	if i < 1 {
		return 0, errors.New("positive intergers above zero only please")
	}
	count := 0
	for i != 1 {
		if i%2 == 0 {
			i = i / 2
		} else {
			i = 3*i + 1
		}
		count++
	}
	return count, nil
}
