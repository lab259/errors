package gqlerrors_test

import (
	"log"
	"os"
	"path"
	"testing"

	"github.com/jamillosantos/macchiato"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"
)

func TestGqlErrors(t *testing.T) {
	log.SetOutput(GinkgoWriter)
	RegisterFailHandler(Fail)
	description := "Testing GraphQL Matchers Test Suite"

	if os.Getenv("CI") == "" {
		macchiato.RunSpecs(t, description)
	} else {
		reporterOutputDir := "../test-results/errors/gqlerrors"
		os.MkdirAll(reporterOutputDir, os.ModePerm)
		junitReporter := reporters.NewJUnitReporter(path.Join(reporterOutputDir, "results.xml"))
		macchiatoReporter := macchiato.NewReporter()
		RunSpecsWithCustomReporters(t, description, []Reporter{macchiatoReporter, junitReporter})
	}
}
