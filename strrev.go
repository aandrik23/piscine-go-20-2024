package piscine

func StrRev(s string) string {
	b := []byte(s)
	i := 0
	for a := len(s) - 1; a >= 0; a-- {
		b[i] = s[a]
		i++
	}
	return (string(b))
}
