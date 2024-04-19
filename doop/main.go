package main

import "os"

func Atoi(s string) (int64, bool) {
	sign := int64(1)
	var result int64
	if len(s) == 0 {
		return 0, false
	}
	start := 0
	if s[0] == '-' {
		sign = -1
		start = 1
	} else if s[0] == '+' {
		start = 1
	}
	for _, char := range s[start:] {
		if char < '0' || char > '9' {
			return 0, false
		}
		digit := int64(char - '0')
		if result > (9223372036854775807-digit)/10 {
			return 0, false // overflow check
		}
		result = result*10 + digit
	}
	return result * sign, true
}

func Itoa(num int64) []byte {
	if num == 0 {
		return []byte{'0'}
	}
	negative := false
	if num < 0 {
		negative = true
		num = -num
	}
	var bytes []byte
	for num > 0 {
		remainder := num % 10
		bytes = append([]byte{byte('0' + remainder)}, bytes...)
		num /= 10
	}
	if negative {
		bytes = append([]byte{'-'}, bytes...)
	}
	return bytes
}

func main() {
	if len(os.Args[1:]) != 3 {
		return
	}
	num1, ok1 := Atoi(os.Args[1])
	operator := os.Args[2]
	num2, ok2 := Atoi(os.Args[3])
	if !ok1 || !ok2 {
		return // If either argument is not a valid number, return immediately.
	}
	var result int64
	var output []byte
	switch operator {
	case "+":
		// Check for overflow in addition
		if (num1 > 0 && num2 > 9223372036854775807-num1) || (num1 < 0 && num2 < -9223372036854775808-num1) {
			return
		}
		result = num1 + num2
	case "-":
		// Check for overflow in subtraction
		if (num1 < 0 && num2 > 9223372036854775807+num1) || (num1 > 0 && num2 < -9223372036854775808+num1) {
			return
		}
		result = num1 - num2
	case "*":
		// Check for overflow in multiplication
		if num1 != 0 && (num2 > 9223372036854775807/num1 || num2 < -9223372036854775808/num1) {
			return
		}
		result = num1 * num2
	case "/":
		if num2 == 0 {
			output = []byte("No division by 0")
		} else if num1 == -9223372036854775808 && num2 == -1 {
			return // Specific case where result is out of int64 bounds
		} else {
			result = num1 / num2
		}
	case "%":
		if num2 == 0 {
			output = []byte("No modulo by 0")
		} else {
			result = num1 % num2
		}
	default:
		return // Exit if the operator is not recognized
	}
	if output == nil {
		output = Itoa(result)
	}
	os.Stdout.Write(output)
	os.Stdout.Write([]byte("\n"))
}
