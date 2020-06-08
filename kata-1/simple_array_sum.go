package main

import "fmt"

func SimpleArraySum(ar []int32) int32 {
	total := int32(0)

	for _, a := range ar {
		total += a
	}
	return total
}

func main() {
	fmt.Println(SimpleArraySum([]int32{1, 2, 3, 4, 5, 6}))
}
