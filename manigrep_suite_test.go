package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestManigrep(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Manigrep Suite")
}
