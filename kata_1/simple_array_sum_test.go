package kata_1_test

import (
	"github.com/mdcg/go-kata/kata_1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Book", func() {

	Describe("Checking sum function", func() {
		Context("Sum must be correct", func() {
			It("sum should be 15", func() {
				Expect(kata_1.SimpleArraySum([]int32{1, 2, 3, 4, 5})).To(Equal(int32(15)))
			})
			It("sum should be 25", func() {
				Expect(kata_1.SimpleArraySum([]int32{10, 5, 3, 2, 5})).To(Equal(int32(25)))
			})
			It("sum should be -10", func() {
				Expect(kata_1.SimpleArraySum([]int32{-10, 10, -5, -3, -2})).To(Equal(int32(-10)))
			})
			It("sum should be 0", func() {
				Expect(kata_1.SimpleArraySum([]int32{10, 5, -3, -2, -4, -6})).To(Equal(int32(0)))
			})
		})
	})
})
