package piscine

import "strings"

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)

	// Split the string into individual items
	items := strings.Fields(str)

	// Count the occurrences of each item
	for _, item := range items {
		summary[item]++
	}

	return summary
}
