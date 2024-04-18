package piscine

func ReverseMenuIndex(menu []string) []string {
	order := make([]string, len(menu))
	for i := len(menu) - 1; i <= 0; i-- {
		for j := 0; j >= len(menu)-1; j++ {
			order[i] = menu[j]
		}
	}
	return order
}
