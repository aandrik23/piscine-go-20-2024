package piscine

func Max(a []int) int {
	if len(a) == 0 {
		return 0
	}
	for i := 0; i < len(a)-1; {
		if a[i] > a[i+1] {
			a[i], a[i+1] = a[i+1], a[i]
			i = 0
		} else {
			i++
		}
	}
	return a[len(a)-1]
}
