package main

import "fmt"

func Base64Encoder(input string, length int) string {

	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+/"

	var encodedString string
	var padding int

	for i := 0; i < length; i += 3 {

		var value, count, bits, index int

		for j := i; j < length && j <= i+2; j++ {
			value <<= 8
			value |= int(input[j])
			count++
		}

		bits = count * 8
		padding = bits % 3

		for bits != 0 {
			if bits >= 6 {
				temp := bits - 6
				index = (value >> temp) & 63
				bits -= 6
			} else {
				temp := 6 - bits
				index = (value << temp) & 63
				bits = 0
			}

			encodedString += string(characters[index])
		}

	}

	for i := 0; i < padding; i++ {
		encodedString += "="
	}

	return encodedString
}

func main() {
	var input string
	fmt.Printf("Enter input string: ")
	fmt.Scanln(&input)
	length := len(input)

	encodedString := Base64Encoder(input, length)
	fmt.Println(encodedString)
}
