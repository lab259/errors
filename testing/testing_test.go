package testing_test

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

func TestTestingMachers(t *testing.T) {
	log.SetOutput(GinkgoWriter)
	RegisterFailHandler(Fail)
	description := "Testing Matchers Test Suite"

	if os.Getenv("CI") == "" {
		macchiato.RunSpecs(t, description)
	} else {
		reporterOutputDir := "../test-results/errors/testing"
		os.MkdirAll(reporterOutputDir, os.ModePerm)
		junitReporter := reporters.NewJUnitReporter(path.Join(reporterOutputDir, "results.xml"))
		macchiatoReporter := macchiato.NewReporter()
		RunSpecsWithCustomReporters(t, description, []Reporter{macchiatoReporter, junitReporter})
	}
}
