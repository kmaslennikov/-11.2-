package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	text := "a10 10 20b 20 30c30 30 dd"
  fmt.Println("Исходная строка:")
  fmt.Println(text)
  fmt.Println("В строке содержатся числа в десятичном формате:")
	for len(text) > 0 {
		spaceIndex := strings.Index(text, " ")
		if spaceIndex == -1 {
			spaceIndex = len(text) - 1
		}
		digit := text[:spaceIndex]
		num, err := strconv.ParseInt(digit, 10, 64)
		text = text[spaceIndex+1:]
		if err != nil {
			continue
		}
		fmt.Printf("%d ", num)
	}
fmt.Println()
}
