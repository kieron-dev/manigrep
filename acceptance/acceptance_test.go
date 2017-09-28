package acceptance_test

import (
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Manigrep", func() {
	Context("invoked with a matching key value contained in the given file", func() {
		It("returns the path from root to the matched key", func() {
			Expect(pathToCmd).To(HaveSuffix("manigrep"))

			command := exec.Command(pathToCmd, "--val-search=bar", "--file=assets/small.yml")
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())

			Eventually(session.Out).Should(gbytes.Say("line 3: foo.bar"))
		})

	})
})
