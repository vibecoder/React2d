package main

import (
	"fmt"

	quoteV3 "rsc.io/quote/v3"
)

func main() {
	fmt.Println(Hello())
}

func Hello() string {
	return quoteV3.HelloV3()
}

func Proverb() string {
	return quoteV3.Concurrency()
}
