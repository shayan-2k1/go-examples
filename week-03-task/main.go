package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	sli := make([]int, 0, 3)
	var usrInput string
	fmt.Println("Enter the numbers")
	for usrInput != "X" {
		fmt.Scan(&usrInput)
		if usrInput != "X" {

			intVar, err := strconv.Atoi(usrInput)
			if err != nil {
				fmt.Println("Unable to convert string to integer!")
			}
			sli = append(sli, intVar)
			sort.Ints(sli)
		} else if usrInput == "X" {
			break
		}
	}
	fmt.Println(sli)
}
