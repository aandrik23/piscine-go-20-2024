package piscine

func IsNotPaired(n int, a []int) bool {
	var counter int
	for i := 0; i < len(a); i++ {
		if n == a[i] {
			counter++
		}
	}
	if counter == 0 {
		return true
	}
	if counter%2 == 0 {
		return true
	}
	return false
}

func Unmatch(a []int) int {
	var check []int
	var result int
	for i := 0; i < len(a)-1; i++ {
		if IsNotPaired(a[i], a) {
			check = append(check, a[i])
		}
	}
	result = check[0]
	return result
}
