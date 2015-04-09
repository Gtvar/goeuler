package main

import (
	"fmt"
)

// http://www.spoj.com/problems/TEST/
func main() {

	var number int

	for {
		fmt.Scanf("%d", &number)
		if (number == 42) {
			break
		}
		fmt.Println(number)
	}

}
