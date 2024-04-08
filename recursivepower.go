package piscine

func RecursivePower(nb int, power int) int {
	var repo int = 1
	a := 1
	if power < 0 {
		repo = 0
	} else if power == 0 {
		repo = 1
	} else if a <= power {
		a++
		repo = repo * nb
	}
	return repo
}
