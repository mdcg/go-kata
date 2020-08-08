package kata_1_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestKata1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Kata1 Suite")
}
