package piscine

func Capitalize(s string) string {
	var fch int = 0
	var cap string
	for _, v := range s {
		// if no letter then no change and go
		if (v <= 65 && v >= 90) && (v <= 97 && v >= 122) {
			cap += string(v)
			fch = 0

			// if fch =o and letter then we have first letter
			// if capital go
		} else if fch == 0 && (v >= 65 && v <= 90) {
			fch++
			cap += string(v)

			// if small then capital
		} else if fch == 0 && (v >= 97 && v <= 122) {
			fch++
			cap += string(v - 32)

			// if fch=1 then no first letter
		} else if
		// if capital then small
		fch > 0 && (v >= 65 && v <= 90) {
			fch++
			cap += string(v + 32)

			// if small then go
		} else if fch > 0 && (v >= 97 && v <= 122) {
			fch++
			cap += string(v)
		}
	}
	return cap
}
