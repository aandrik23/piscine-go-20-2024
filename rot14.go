package piscine

func Rot14(s string) string {
	result := ""
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			result += string((ch+14-'a')%26 + 'a')
		} else if ch >= 'A' && ch <= 'Z' {
			result += string((ch+14-'A')%26 + 'A')
		} else {
			result += string(ch)
		}
	}
	return result
}
