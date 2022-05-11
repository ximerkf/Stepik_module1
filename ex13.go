package main

import "fmt"

func main() {
	var a, fbn int
	i := 1
	fb0 := 0
	fb1 := 1
	fmt.Scan(&a)
	for fbn <= a {
		fbn = fb0 + fb1
		fb0, fb1 = fb1, fbn
		i++
		if a == fbn {
			break
		}
	}
	if a == fbn {
		fmt.Println(i)
	} else {
		fmt.Println(-1)
	}
}
