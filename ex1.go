package main

import "fmt"

func main() {
	var a, sum int
	fmt.Scan(&a)
	sum = a/100 + a/10%10 + a%10
	fmt.Println(sum)
}
