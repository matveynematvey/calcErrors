package main

import (
	"fmt"

	calc "./calc"
)

func check(msg string, err error) {
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(msg)
	}
}

func main() {
	check(calc.Add(1, 2))
	check(calc.Div(1, 0))
	check(calc.Div(3, 2))
	check(calc.Div(4, 2))
}
