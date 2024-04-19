package piscine

func Abort(a, b, c, d, e int) int {
	var arr [5]int
	arr[0] = a
	arr[1] = b
	arr[2] = c
	arr[3] = d
	arr[4] = e
	for i := 0; i < 4; {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
			i = 0
		} else {
			i++
		}
	}
	return arr[2]
}
