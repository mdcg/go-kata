package kata_3_test

import (
	"github.com/mdcg/go-kata/kata_3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Book", func() {
	Describe("Checking expression matter function", func() {
		Context("Results must be correct", func() {
			It("2, 1, 2 => 6", func() {
				Expect(kata_3.ExpressionMatter(2, 1, 2)).To(Equal(6))
			})
			It("2, 1, 1 => 4", func() {
				Expect(kata_3.ExpressionMatter(2, 1, 1)).To(Equal(4))
			})
			It("1, 10, 1 => 12", func() {
				Expect(kata_3.ExpressionMatter(1, 10, 1)).To(Equal(12))
			})
			It("1, 1, 1 => 3", func() {
				Expect(kata_3.ExpressionMatter(1, 1, 1)).To(Equal(3))
			})
			It("1, 2, 3 => 9", func() {
				Expect(kata_3.ExpressionMatter(1, 2, 3)).To(Equal(9))
			})
			It("1, 3, 1 => 5", func() {
				Expect(kata_3.ExpressionMatter(1, 3, 1)).To(Equal(5))
			})
			It("2, 2, 2 => 8", func() {
				Expect(kata_3.ExpressionMatter(2, 2, 2)).To(Equal(8))
			})
			It("5, 1, 3 => 20", func() {
				Expect(kata_3.ExpressionMatter(5, 1, 3)).To(Equal(20))
			})
			It("3, 5, 7 => 105", func() {
				Expect(kata_3.ExpressionMatter(3, 5, 7)).To(Equal(105))
			})
			It("5, 6, 1 => 35", func() {
				Expect(kata_3.ExpressionMatter(5, 6, 1)).To(Equal(35))
			})
			It("1, 6, 1 => 8", func() {
				Expect(kata_3.ExpressionMatter(1, 6, 1)).To(Equal(8))
			})
			It("2, 6, 1 => 14", func() {
				Expect(kata_3.ExpressionMatter(2, 6, 1)).To(Equal(14))
			})
			It("6, 7, 1 => 48", func() {
				Expect(kata_3.ExpressionMatter(6, 7, 1)).To(Equal(48))
			})
			It("2, 10, 3 => 60", func() {
				Expect(kata_3.ExpressionMatter(2, 10, 3)).To(Equal(60))
			})
			It("1, 8, 3 => 27", func() {
				Expect(kata_3.ExpressionMatter(1, 8, 3)).To(Equal(27))
			})
			It("9, 7, 2 => 126", func() {
				Expect(kata_3.ExpressionMatter(9, 7, 2)).To(Equal(126))
			})
			It("1, 1, 10 => 20", func() {
				Expect(kata_3.ExpressionMatter(1, 1, 10)).To(Equal(20))
			})
			It("9, 1, 1 => 18", func() {
				Expect(kata_3.ExpressionMatter(9, 1, 1)).To(Equal(18))
			})
			It("10, 5, 6 => 300", func() {
				Expect(kata_3.ExpressionMatter(10, 5, 6)).To(Equal(300))
			})
		})
	})
})
