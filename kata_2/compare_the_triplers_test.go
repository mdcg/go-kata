package kata_2_test

import (
	"github.com/mdcg/go-kata/kata_2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Book", func() {

	Describe("Checking compare triplets function", func() {
		Context("Comparison must be correct", func() {
			It("SumA = 1, SumB = 1", func() {
				Expect(kata_2.CompareTriplets([]int32{5, 6, 7}, []int32{3, 6, 10})).To(Equal([]int32{1, 1}))
			})
			It("SumA = 0, SumB = 0", func() {
				Expect(kata_2.CompareTriplets([]int32{1, 2, 3}, []int32{1, 2, 3})).To(Equal([]int32{0, 0}))
			})
			It("SumA = 0, SumB = 3", func() {
				Expect(kata_2.CompareTriplets([]int32{6, 8, 12}, []int32{7, 9, 15})).To(Equal([]int32{0, 3}))
			})
			It("SumA = 3, SumB = 0", func() {
				Expect(kata_2.CompareTriplets([]int32{10, 15, 20}, []int32{5, 6, 7})).To(Equal([]int32{3, 0}))
			})
		})
	})
})
