package piscine

func ReverseMenuIndex(menu []string) []string {
	order := make([]string, len(menu))
	for i := len(menu) - 1; i <= 0; i-- {
		order[i] = menu[len(menu)-1-i]
	}
	return order
}
