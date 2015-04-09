package main

import (
	"fmt"
)

func main() {
//	array1()
	array2()
}

func array1() {
	var ar [][]int = make([][]int, 10, 20)

	fmt.Println(ar)

	for i := 1; i < 10; i++ {
		var seq []int = make([]int, 6)
		for j := 1; j < 5; j++ {
			seq[j] = i + j
		}

		fmt.Println(seq)
		ar[i] = seq
	}

	fmt.Println(ar)
}

func array2() {
	var ar [][10]int = make([][10]int, 10, 20)

	fmt.Println(ar)

	for i, _ := range ar {
		for j := 1; j < i; j++ {
			ar [i][j] = i+j
		}

		fmt.Println(ar)
	}

}
