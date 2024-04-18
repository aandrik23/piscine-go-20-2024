package piscine

func DescendAppendRange(max, min int) []int {
	var s []int
	for i := max; i < min; i-- {
		s = append(s, i)
	}
	return s
}
