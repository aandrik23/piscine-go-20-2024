package piscine

func Index(s string, toFind string) int {
	l1 := len(s)
	l2 := len(toFind)
	if l1 >= l2 {
		for i := 0; i < l1; i++ {
			j := 0
			t := i
			for j < l2 && s[t] == toFind[j] {
				j++
				t++
			}
			if j == l2 {
				return i
			}
		}
	}
	return -1
}
