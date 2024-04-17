package piscine

func Map(f func(int) bool, a []int) []bool {
	return_array := []bool{}
	for i := range a {
		return_array = append(return_array, f(a[i]))
	}
	return return_array
}
