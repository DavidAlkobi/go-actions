package main

import "fmt"

// sumInts sums up integers
func sumInts(nums ...int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	fmt.Println(sumInts(1, 2, 3, 4, 5))
}
