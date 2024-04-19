package piscine

func Max(a []int) int {
	if len(a) < 1 {
		return 0
	}
	for i := 1; i < len(a)-1; {
		if a[i-1] > a[i] {
			a[i-1], a[i] = a[i], a[i-1]
			i = 1
		} else {
			i++
		}
	}
	return a[len(a)-1]
}
