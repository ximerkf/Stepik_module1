package main

import "fmt"

func main() {
	var k, h, m int
	fmt.Scan(&k)
	h = k / 3600
	m = (k - h*3600) / 60
	fmt.Printf("It is %d hours %d minutes.", h, m)
}
