package piscine

func IsPrime(nb int) bool {
	i := 2

	if nb <= 0 {
		return false
	} else if nb == 2 {
		return true
	} else {
		for i = 2; i <= nb; i++ {
			if nb%i != 0 {
				return false
			}
		}
	}
	return true
}
