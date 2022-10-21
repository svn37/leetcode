package sorting

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) / 2

	leftCopy := append(arr[:0:0], arr[:mid]...)
	rightCopy := append(arr[:0:0], arr[mid:]...)

	left := mergeSort(leftCopy)
	right := mergeSort(rightCopy)

	i, j := 0, 0
	for i < len(left) || j < len(right) {
		if j == len(right) || (i < len(left) && left[i] <= right[j]) {
			arr[i+j] = left[i]
			i++
		} else {
			arr[i+j] = right[j]
			j++
		}
	}
	return arr
}
