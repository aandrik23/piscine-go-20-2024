package piscine

func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for a := range tab {
		if f(tab[a]) == true {
			count++
		}
	}
	return count

}
