package main

import "fmt"

func main() {
	var a, b int
	c := 1
	fmt.Scan(&a, &b)
	for i := a; i <= b; i++ {
		if i%7 == 0 {
			c = i
		}
	}
	if c == 1 {
		fmt.Println("NO")
	} else {
		fmt.Println(c)
	}
}
