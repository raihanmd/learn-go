package sorter

// Time complexity O(n^2)
func BubbleSort(arr *[]int) *[]int {
	var swapper bool = true

	for i := 0; i < len(*arr); i++ {
		swapper = false
		for j := 0; j < len(*arr)-i-1; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
				swapper = true
			}
		}
		if !swapper {
			break
		}
	}

	return arr
}
