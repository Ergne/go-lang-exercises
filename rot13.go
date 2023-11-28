package main

import "fmt"

func Rot13(msg string) string {
	result := ""
	for _, char := range msg {
		if ('a' <= char && char <= 'm') || ('A' <= char && char <= 'M') {
			char = char + 13
		} else if ('n' <= char && char <= 'z') || ('N' <= char && char <= 'Z') {
			char -= 13
		}
		fmt.Println(char)
		result += string(char)
	}
	return result
}

func main() {
	fmt.Println(Rot13("Va gur ryringbef, gur rkgebireg ybbxf ng gur BGURE thl'f fubrf."))
}
