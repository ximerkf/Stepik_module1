package main

import "fmt"

func main() {
	var a, b, c int
	count := 1
	fmt.Scan(&a)
	fmt.Scan(&b)
	for i := 1; i < a; i++ {
		fmt.Scan(&c)
		if c < b {
			b = c
			count = 1
		} else if c == b {
			count++
		}
	}
	fmt.Println(count)
}
