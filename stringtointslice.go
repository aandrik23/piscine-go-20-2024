package piscine

func StringToIntSlice(str string) []int {
	intSlice := make([]int, len(str))

	for i, char := range str {
		intSlice[i] = int(char)
	}

	return intSlice
}
