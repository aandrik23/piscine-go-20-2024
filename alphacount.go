package piscine

func AlphaCount(s string) int {
	srunes := []rune(s)
	var counter int
	if len(s) > 0 {
		for _, v := range srunes {
			if (v >= 'A' && v <= 'Z') || (v >= 'a' && v <= 'z') {
				counter++
			}
		}
	}
	return counter
}
