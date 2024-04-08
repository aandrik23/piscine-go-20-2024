package piscine

func IterativePower(nb int, power int) int {
	var itpo int = 1

	if power < 0 {
		itpo = 0
	} else if power == 0 {
		itpo = 1
	} else if power == 1 {
		itpo = nb
	} else {
		for a := 2; a <= power; a++ {
			itpo = itpo * nb
		}
	}
	return itpo
}
