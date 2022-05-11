package main

import "fmt"

func main() {
	var b int
	var a string
	fmt.Scan(&a, &b)
	for _, elem := range a {
		if int(elem-'0') != b {
			fmt.Print(string(elem))
		}
	}
}
