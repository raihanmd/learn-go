package utils

func MergeSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	mid := len(a) / 2
	return merge(MergeSort(a[:mid]), MergeSort(a[mid:]))
}

func merge(a, b []int) []int {
	c := make([]int, len(a)+len(b))
	i, j, k := 0, 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			c[k] = a[i]
			i++
		} else {
			c[k] = b[j]
			j++
		}
		k++
	}
	for i < len(a) {
		c[k] = a[i]
		i++
		k++
	}
	for j < len(b) {
		c[k] = b[j]
		j++
		k++
	}
	return c
}
