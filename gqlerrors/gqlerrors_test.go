package gqlerrors_test

import (
	"log"
	"testing"

	"github.com/jamillosantos/macchiato"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGqlErrors(t *testing.T) {
	log.SetOutput(GinkgoWriter)
	RegisterFailHandler(Fail)
	macchiato.RunSpecs(t, "Testing GraphQL Matchers Test Suite")
}
