package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)

	// Initialize variables to keep track of item boundaries
	start := 0

	// Iterate over the characters in the string
	for i := 0; i <= len(str); i++ {
		// If the current character is a space or it's the last character in the string
		if i == len(str) || str[i] == ' ' {
			// Extract the item from the substring
			item := str[start:i]

			// Increment the count of the item in the summary map
			summary[item]++

			// Update start to the next character after the space
			start = i + 1
		}
	}

	return summary
}
