package main

import (
	"fmt"
	"log"
)

/*
 * @link http://www.spoj.com/problems/PRIME1/
 */
func main() {

	var t int
	fmt.Scanf("%d", &t)

	if t > 10 {
		log.Fatal("Must be less 10")
	}

	var primes [][]int = make([][]int, t)

	var seq []int

	var a, b int
	for i := 0; i < t; i++ {
		fmt.Scanf("%d %d", &a, &b)

		seq = findPrime(a, b)
		primes = append(primes, seq)
	}

	for _, items := range primes {
		if len(items) < 1 {
			continue
		}

		for _, value := range items {
			fmt.Println(value)
		}
		fmt.Println("")
	}

}

/*
 * Find primes between numbers
 */
func findPrime(start int, end int) []int {
	seq := []int{}
	for i := start; i <= end; i++ {
		if is_prime(i) {
			seq = append(seq, i)
		}
	}

	return seq
}

/*
 * Check number for prime
 */
func is_prime(number int) bool {
	if (number <= 3) {
		return number > 1
	}

	if (number%2 == 0 || number%3 == 0) {
		return false
	}

	for i := 5; i*i < number; i+=6 {
		if number%i == 0 || number%(i+2) == 0 {
			return false
		}
	}

	return true
}
