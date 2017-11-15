package algorithms

func BinarySearch(ints []int, targetValue int) int {
	min := 0
	max := len(ints) - 1
	for min <= max {
		mid := (max + min) / 2

		if ints[mid] == targetValue {
			return mid
		} else if ints[mid] < targetValue {
			min = mid + 1
		} else {
			max = mid - 1
		}

	}
	return -1
}
