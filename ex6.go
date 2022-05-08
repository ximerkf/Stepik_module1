package main

import "fmt"

func main() {
	var a, b int
	var ar float32
	fmt.Scan(&a, &b)
	ar = (float32(a + b)) / 2
	fmt.Println(ar)
}
