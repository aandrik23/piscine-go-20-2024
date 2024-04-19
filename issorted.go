package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	counter := 0
	kounter := 0
	answer := true
	for i := 1; i < len(a); i++ {
		if f(a[i-1], a[i]) <= 0 {
			counter++
		}
		if f(a[i-1], a[i]) >= 0 {
			kounter++
		}
	}
	if len(a)-1 > counter && len(a)-1 > kounter {
		answer = false
	}

	return answer
}
