package kata_2_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestKata2(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Kata2 Suite")
}
