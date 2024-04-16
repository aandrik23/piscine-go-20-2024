package piscine

func StrLen(s string) int {
	count := 0
	for v := 1; v <= len(s); v++ {
		count = count + 1
	}
	return count
}
