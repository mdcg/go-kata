package kata_1

func SimpleArraySum(ar []int32) int32 {
	total := int32(0)

	for _, a := range ar {
		total += a
	}
	return total
}
