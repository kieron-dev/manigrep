package grepper_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGrepper(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Grepper Suite")
}
