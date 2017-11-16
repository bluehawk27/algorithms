package algorithms

func RotateImage(a [][]int) [][]int {
	z := 0
	b := make([][]int, len(a))
	for i := range b {
		b[i] = make([]int, len(a[0]))
	}
	for iy := range a {
		iy = len(a) - 1 - iy
		for jx := range a[iy] {
			z = jx
			jx = len(a) - 1 - jx
			b[iy][z] = a[jx][iy]
		}
	}
	return b
}
