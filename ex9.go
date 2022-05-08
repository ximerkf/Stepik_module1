package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a)
	for a >= 10 {
		b = a
		c = 0
		for b > 0 {
			c = c + b%10
			b = b / 10
		}
		a = c
	}
	fmt.Println(a)
}
