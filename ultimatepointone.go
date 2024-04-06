package piscine

func UltimatePointOne(n ***int) {
	var a = 1
	var b *int
	*n = &b
	*b = a
}
