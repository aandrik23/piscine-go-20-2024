package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	if len(a) > 1 {
		for i := 1; i < len(a); {
			if a[i-1] > a[i] {
				a[i-1], a[i] = a[i], a[i-1]
				i = 1
			} else {
				i++
			}
		}
	}
}
