package piscine

func IterativePower(nb int, power int) int {
	var itpo int = 1

	if power < 0 {
		itpo = 0
	} else if power == 0 {
		itpo = 1
	} else {
		for a := 1; a <= power; a++ {
			itpo = itpo * nb
		}
	}
	return itpo
}
