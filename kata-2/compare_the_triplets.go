package main

func CompareTriplets(a, b []int32) []int32 {
	var sumA, sumB int32
	for i, v := range a {
		if v > b[i] {
			sumA += 1
		}
		if b[i] > v {
			sumB += 1
		}
	}
	return []int32{sumA, sumB}
}

func main() {}
