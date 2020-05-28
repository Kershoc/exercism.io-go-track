package accumulate

// Callback is any function that can take and return a string
type Callback func(x string) string

// Accumulate iterates over a list and runs the provided function on it returning a new list of the results
func Accumulate(s []string, f Callback) []string {
	var out []string
	for _, v := range s {
		out = append(out, f(v))
	}
	return out
}
