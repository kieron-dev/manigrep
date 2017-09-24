package grepper_test

import (
	"github.com/kieron-pivotal/manigrep/grepper"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Grepper", func() {
	It("should assert something about grepper", func() {

		_, err := grepper.New("somefile")
		Expect(err).ToNot(HaveOccurred())
	})

})
