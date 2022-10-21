package sorting

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func countSort(nums, aux []int, exp int) {
	n := len(nums)

	// radix is 10, could be something else
	count := make([]int, 10)

	// store count of occurrences of current digit
	for i := 0; i < n; i++ {
		count[(nums[i]/exp)%10]++
	}

	// modify count so it contains actual position
	// of this digit in sorted array (aux)
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	// build output array
	for i := n - 1; i >= 0; i-- {
		// get the actual position and put a number there
		aux[count[(nums[i]/exp)%10]-1] = nums[i]
		// decrease position
		count[(nums[i]/exp)%10]--
	}

	copy(nums, aux)
}

func radixSort(nums []int) {
	if len(nums) < 2 {
		return
	}

	// find the maximum number to know number of digits
	maxVal := nums[0]
	for i := range nums {
		maxVal = max(maxVal, nums[i])
	}

	aux := make([]int, len(nums))
	// exp is 10^i where i is current digit number
	// 1, 10, 100, 1000 ...
	for exp := 1; maxVal/exp > 0; exp *= 10 {
		countSort(nums, aux, exp)
	}
}
