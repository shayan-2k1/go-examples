package main

import "fmt"

func main() {
	var x float64
	fmt.Println("Enter a number in floating point representation:")
	fmt.Scan(&x)
	fmt.Printf("The truncated number is %d", int(x))
}
