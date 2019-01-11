package testing_test

import (
	"log"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jamillosantos/macchiato"
)

func TestTestingMachers(t *testing.T) {
	log.SetOutput(GinkgoWriter)
	RegisterFailHandler(Fail)
	macchiato.RunSpecs(t, "Testing Matchers Test Suite")
}
