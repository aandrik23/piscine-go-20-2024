package piscine

func Unmatch(a []int) int {
	frequency := make(map[int]int)
	for _, value := range a {
		frequency[value]++
	}
	for _, number := range a {
		if count, exists := frequency[number]; exists && count%2 != 0 {
			return number
		}
	}
	return -1
}
