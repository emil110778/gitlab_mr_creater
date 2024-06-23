package helper

import (
	"golang.org/x/exp/constraints"
)

func SliceToInt[T constraints.Integer](val []T) []int {
	result := make([]int, len(val))
	for i, v := range val {
		result[i] = int(v)
	}
	return result
}
