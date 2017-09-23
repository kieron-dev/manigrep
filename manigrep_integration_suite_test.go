package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"

	"testing"
)

func TestManigrepIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Manigrep IntegrationSuite")
}

var pathToCmd string

var _ = BeforeSuite(func() {
	var err error
	pathToCmd, err = gexec.Build("github.com/kieron-pivotal/manigrep")
	Expect(err).ToNot(HaveOccurred())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})
