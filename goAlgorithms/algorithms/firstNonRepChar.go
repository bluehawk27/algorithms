package algorithms

func FirstNotRepeatingCharacter(s string) string {
	hash := make(map[rune]uint, len(s)) //preallocate the map size
	for _, v := range s {
		hash[v]++
	}

	for _, v := range s {
		if hash[v] == 1 {
			return string(v)
		}
	}
	return "_"
}
