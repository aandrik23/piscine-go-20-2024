package piscine

func ShoppingListSort(slice []string) []string {
	if len(slice) > 1 {
		for i := 0; i < len(slice)-1; {
			if len(slice[i]) > len(slice[i+1]) {
				slice[i], slice[i+1] = slice[i+1], slice[i]
				i = 0
			} else {
				i++
			}
		}
	}
	return slice
}
