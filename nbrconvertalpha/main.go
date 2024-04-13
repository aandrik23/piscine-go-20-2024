package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// create a map to the alphabet
	alphabet := map[string]rune{
		"1":  'a',
		"2":  'b',
		"3":  'c',
		"4":  'd',
		"5":  'e',
		"6":  'f',
		"7":  'g',
		"8":  'h',
		"9":  'i',
		"10": 'j',
		"11": 'k',
		"12": 'l',
		"13": 'm',
		"14": 'n',
		"15": 'o',
		"16": 'p',
		"17": 'q',
		"18": 'r',
		"19": 's',
		"20": 't',
		"21": 'u',
		"22": 'v',
		"23": 'w',
		"24": 'x',
		"25": 'y',
		"26": 'z',
	}
	// use a boolean flag if --upper is the first argument
	flag := false
	var arguments []string
	// give arguments list of strings value depending on flag
	if len(os.Args) > 1 {
		if os.Args[1] == "--upper" {
			flag = true
			arguments = os.Args[2:]
		} else {
			arguments = os.Args[1:]
		}
	}
	for _, argument := range arguments {
		// if value exists print value as lower case if flag false or upper case if flag true
		if value, exists := alphabet[argument]; exists {
			if !flag {
				z01.PrintRune(value)
			} else {
				z01.PrintRune(value - 32)
			}
		} else {
			z01.PrintRune(' ')
		}
	}
	if len(arguments) > 0 {
		z01.PrintRune('\n')
	}
}
