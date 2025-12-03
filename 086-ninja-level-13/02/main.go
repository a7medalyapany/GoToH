package main

import (
	"GoToH/086-ninja-level-13/02/quote"
	"GoToH/086-ninja-level-13/02/word"
	"fmt"
)

func main() {
	fmt.Println(word.Count(quote.Quote))
	fmt.Println(word.Count(quote.ArQuote))

	fmt.Println(word.UseCount(quote.Quote))
	fmt.Println(word.UseCount(quote.ArQuote))
}