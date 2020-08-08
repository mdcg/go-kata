package kata_3

import "sort"

// https://www.codewars.com/kata/5ae62fcf252e66d44d00008e
func ExpressionMatter(a int, b int, c int) int {
	arr := []int{a * (b + c), a * b * c, a + b + c, (a + b) * c}
	sort.Ints(arr)
	return arr[len(arr)-1]
}
