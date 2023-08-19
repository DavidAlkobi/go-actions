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
<<<<<<< HEAD
	fmt.Println(sumInts(1, 2, 3, 4, 5))
}
=======
	printText("Hello, World!")

>>>>>>> a25d67f312111f861fb7bda89323d660f8d9f16e
