package kata_3_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestKata3(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Kata3 Suite")
}
