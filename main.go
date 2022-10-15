package main

import "fmt"

func main() {
	a := 0
	fmt.Println(a)
	a = 505

	for {
		fmt.Println(a)
		if a == 5000000000000000005 {
			break
		}
		a = (a-5)*10 + 5

	}

}
