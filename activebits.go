package piscine

func ActiveBits(n int) int {
	var counter int
	for n >= 1 {
		if n%2 != 0 {
			counter++
		}
		n /= 2
	}
	return counter
}
