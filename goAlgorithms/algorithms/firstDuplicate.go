package algorithms

func FirstDuplicate(a []int) int {
	hash := make(map[int]int)
	for _, v := range a {
		_, ok := hash[v]
		if ok {
			hash[v]++
			if hash[v] > 1 {
				return v
			}
		} else {
			hash[v] = 1
		}
	}
	return -1
}
