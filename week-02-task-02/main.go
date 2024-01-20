package main

import (
	"fmt"

	"github.com/shayan-2k1/go-examples/week-02-task-02/findian"
)

func main() {
	var input string
	fmt.Println("Enter any string:")
	fmt.Scan(&input)
	check := findian.Find(input)
	switch check {
	case true:
		fmt.Println("Found!")
	case false:
		fmt.Println("Not Found!")
	}

}
