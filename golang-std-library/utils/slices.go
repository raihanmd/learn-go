package utils

import "slices"

func SortWithSlices(arr []int) []int {
	slices.Sort(arr)

	return arr
}
