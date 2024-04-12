package piscine

func TrimAtoi(s string) int {
	var result int
	sign := 1
	encounternumber := false
	for _, v := range s {
		if v == '-' && !encounternumber {
			sign = -1
		}
		if v >= '0' && v <= '9' {
			result = result*10 + int(v-'0')
			encounternumber = true
		}
	}
	return result * sign
}
