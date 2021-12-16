package main

import (
	"fmt"

	str "github.com/d0uble3L/stringmod/strings"

	qt "github.com/d0uble3L/stringmod/quotes"
)

func main() {
	o, e := str.CountOddEven("12345")
	fmt.Println(o, e)

	fmt.Println(qt.GetEmoji("turtle"))
}
