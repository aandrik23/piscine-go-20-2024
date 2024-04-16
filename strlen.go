package piscine

func StrLen(s string) int {
	count := 0
	for v := 0; v < len(s)-1; v++ {
		count = count + 1
	}
	return count
}
