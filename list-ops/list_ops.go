//Package listops implements a series of methods to operate on a list of integers
package listops

type binFunc func(x, y int) int
type predFunc func(n int) bool
type unaryFunc func(x int) int

// IntList is a slice of integers with a bunch of helper methods attached
type IntList []int

// Foldl folds each item into the accumulator a from the left via fn
func (list IntList) Foldl(fn binFunc, a int) int {
	for _, item := range list {
		a = fn(a, item)
	}
	return a
}

// Foldr folds each item into the accumulator a from the right via fn
func (list IntList) Foldr(fn binFunc, a int) int {
	for i := list.Length() - 1; i >= 0; i-- {
		a = fn(list[i], a)
	}
	return a
}

// Concat combines all items in all lists into one flattened list
func (list IntList) Concat(listOfIntLists []IntList) IntList {
	out := list
	for _, item := range listOfIntLists {
		out = out.Append(item)
	}
	return out
}

//Append add all items in the passedInList to the end of list
func (list IntList) Append(passedInList IntList) IntList {
	out := list
	for _, item := range passedInList {
		out = append(out, item)
	}
	return out
}

//Reverse returns the original list backwards
func (list IntList) Reverse() IntList {
	reversed := IntList{}
	for i := list.Length() - 1; i >= 0; i-- {
		reversed = append(reversed, list[i])
	}

	return reversed
}

//Map returns a list of the results of applying fn on all items
func (list IntList) Map(fn unaryFunc) IntList {
	out := IntList{}
	for _, item := range list {
		out = append(out, fn(item))
	}
	return out
}

//Length returns total number of items in list
func (list IntList) Length() int {
	return len(list)
}

//Filter returns a list of all items for which appling fn on item is True
func (list IntList) Filter(fn predFunc) IntList {
	out := IntList{}
	for _, item := range list {
		if fn(item) {
			out = append(out, item)
		}
	}
	return out
}
