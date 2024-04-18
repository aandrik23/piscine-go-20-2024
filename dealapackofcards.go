package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func DealAPackOfCards(deck []int) {
	players := 4
	cardsPerPlayer := len(deck) / players

	for i := 0; i < players; i++ {
		fmt.Printf("Player %d: ", i+1)
		for j := 0; j < cardsPerPlayer; j++ {
			index := i*cardsPerPlayer + j
			fmt.Printf("%d ", deck[index])
		}
		z01.PrintRune('\n')
	}
}
