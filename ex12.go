package main

import "fmt"

func main() {
	var a int
	c := 1
	fmt.Scan(&a)
	for c <= a {
		fmt.Print(c, " ")
		c = c * 2

	}
}
