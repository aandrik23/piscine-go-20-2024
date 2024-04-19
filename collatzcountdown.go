package piscine

func CollatzCountdown(start int) int {
	var count int
	if start < 1 {
		return -1
	} else if start == 1 {
		return 0
	} else {
		for start != 1 {
			if start%2 == 0 {
				start /= 2
				count++
			} else {
				start = 3*start + 1
				count++
			}
		}
	}
	return count
}
