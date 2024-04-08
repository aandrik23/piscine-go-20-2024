package piscine

func IsPrime(nb int) bool {
	i := 2

	if nb == 2 {
		return true
	} else {
		for i = 3; i <= nb; i = i + 2 {
			if nb%i != 0 {
				return false
			}
		}
	}
	return true
}
