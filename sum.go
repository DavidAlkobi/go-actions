// This program sums up integers passed as arguments
// Usage: go run sum.go 1 2 3 4 5
// Output: 15
// test 23
package main

import (
	"os"
	"strconv"
)

// sumInts sums up integers
func sumInts(nums ...int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	return sum
}

// Get integers from arguments and sum them up
func main() {
	// get all the arguments
	args := os.Args[1:]

	// convert the arguments to integers
	var nums []int
	for _, arg := range args {
		num, err := strconv.Atoi(arg)
		if err != nil {
			continue
		}
		nums = append(nums, num)
	}

	// sum up the integers
	sum := sumInts(nums...)

	// convert the sum to string
	sum_string := strconv.Itoa(sum)

	// print the sum
	os.Stdout.WriteString(sum_string)

}
