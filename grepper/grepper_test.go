package grepper_test

import (
	"path/filepath"

	"github.com/kieron-pivotal/manigrep/grepper"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Grepper", func() {
	Context("instantiation", func() {
		Context("when a valid yaml file passed to New", func() {
			It("succeeds without error", func() {
				_, err := grepper.New(assetPath("small.yml"))
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("when an invalid yaml file passed to New", func() {
			It("fails with appropriate error message", func() {
				_, err := grepper.New(assetPath("small-error.yml"))
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Error parsing YAML"))
			})
		})

		Context("when a non-existant file is passed to New", func() {
			It("fails with appropriate error message", func() {
				_, err := grepper.New(assetPath("does-not-exist.yml"))
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("no such file or directory"))
			})
		})
	})

})

func assetPath(filename string) string {
	path, err := filepath.Abs(filepath.Join("assets", filename))
	Expect(err).ToNot(HaveOccurred())
	return path
}
