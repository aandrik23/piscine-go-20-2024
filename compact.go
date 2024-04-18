package piscine

func Compact(slice *[]string) int {
	count := 0

	// Iterate over the slice backwards to delete elements with zero values
	for i := len(*slice) - 1; i >= 0; i-- {
		if (*slice)[i] == "" {
			// Delete the element with zero value
			*slice = append((*slice)[:i], (*slice)[i+1:]...)
		} else {
			// Increment the count for non-zero values
			count++
		}
	}

	return count
}
