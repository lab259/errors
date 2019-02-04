package testing_test

import (
	"github.com/lab259/errors"
	"github.com/lab259/errors/testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ErrorWithModuleMatcher Test Suite", func() {
	It("should initialize the matcher", func() {
		m := testing.ErrorWithModule("Module1")
		Expect(m.Expected).To(Equal("Module1"))
	})

	It("should match an ErrorWithModule", func() {
		m := testing.ErrorWithModule("Module1")

		errModule := errors.Wrap(errors.New("error1"), errors.Module("Module1"))

		result, err := m.Match(errModule)
		Expect(result).To(BeTrue())
		Expect(err).ToNot(HaveOccurred())
	})

	It("should match a deep ErrorWithModule", func() {
		m := testing.ErrorWithModule("Module1")

		errModule := errors.Wrap(errors.New("error1"), errors.Module("Module1"), errors.Message("message1"), errors.Http(404))

		result, err := m.Match(errModule)
		Expect(result).To(BeTrue())
		Expect(err).ToNot(HaveOccurred())
	})

	It("should fail matching an ErrorWithModule", func() {
		m := testing.ErrorWithModule("Module1")

		errModule := errors.Wrap(errors.New("error1"), errors.Module("Module2"))

		result, err := m.Match(errModule)
		Expect(result).To(BeFalse())
		Expect(err).ToNot(HaveOccurred())
	})

	It("should fail matching a deep ErrorWithModule", func() {
		m := testing.ErrorWithModule("Module1")

		errModule := errors.Wrap(errors.New("error1"), errors.Module("Module2"), errors.Message("message1"), errors.Http(404))

		result, err := m.Match(errModule)
		Expect(result).To(BeFalse())
		Expect(err).ToNot(HaveOccurred())
	})

	It("should fail matching an error chain with no ErrorWithModule", func() {
		m := testing.ErrorWithModule("Module1")

		errModule := errors.Wrap(errors.New("error1"), errors.Message("message1"), errors.Http(404))

		result, err := m.Match(errModule)
		Expect(result).To(BeFalse())
		Expect(err).ToNot(HaveOccurred())
	})

	It("should fail matching when the actual is not an error", func() {
		m := testing.ErrorWithModule("Module1")

		result, err := m.Match("not an error")
		Expect(result).To(BeFalse())
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("`actual` is not an `error`"))
	})

	It("should run through the Gomega", func() {
		errModule := errors.Wrap(errors.New("error1"), errors.Module("Module1"), errors.Message("message1"), errors.Http(404))

		Expect(errModule).To(testing.ErrorWithModule("Module1"))
	})

	It("should run through the Gomega (negating)", func() {
		errModule := errors.Wrap(errors.New("error1"), errors.Message("message1"), errors.Http(404))

		Expect(errModule).ToNot(testing.ErrorWithModule("Module1"))
	})

	It("should generate a failure message with different error Modules", func() {
		errModule := errors.Wrap(errors.New("error1"), errors.Module("Module1"), errors.Message("message1"), errors.Http(404))

		m := testing.ErrorWithModule("Module2")
		message := m.FailureMessage(errModule)
		Expect(message).To(ContainSubstring("Module2"))
		Expect(message).To(ContainSubstring("to have module"))
		Expect(message).To(ContainSubstring("Module1"))
	})

	It("should generate a failure message for non existing ErrorWithModule in the error chain", func() {
		errModule := errors.Wrap(errors.New("error1"), errors.Message("message1"), errors.Http(404))

		m := testing.ErrorWithModule("Module1")
		message := m.FailureMessage(errModule)

		Expect(message).To(ContainSubstring("does not have ErrorWithModule"))
		Expect(message).To(ContainSubstring("Module1"))
	})

	It("should generate a failure message negating the Module found", func() {
		errModule := errors.Wrap(errors.New("error1"), errors.Module("Module1"), errors.Message("message1"), errors.Http(404))

		m := testing.ErrorWithModule("Module1")
		message := m.NegatedFailureMessage(errModule)
		Expect(message).To(ContainSubstring("not to have module"))
		Expect(message).To(ContainSubstring("Module1"))
	})
})
