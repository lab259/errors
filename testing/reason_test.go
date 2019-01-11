package testing_test

import (
	"github.com/lab259/errors"
	"github.com/lab259/errors/testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ErrorWithReasonMatcher Test Suite", func() {
	anyError := errors.New("any")
	anotherError := errors.New("another error")

	It("should initialize the matcher", func() {
		m := testing.ErrorWithReason(anyError)
		Expect(m.Expected).To(Equal(anyError))
	})

	It("should match an ErrorWithReason", func() {
		m := testing.ErrorWithReason(anyError)

		errWithReason := errors.Wrap(anyError)

		result, err := m.Match(errWithReason)
		Expect(result).To(BeTrue())
		Expect(err).ToNot(HaveOccurred())
	})

	It("should match a deep ErrorWithReason", func() {
		m := testing.ErrorWithReason(anyError)

		errWithReason := errors.Wrap(anyError, errors.Code("code1"), errors.Message("message1"), errors.Http(404))

		result, err := m.Match(errWithReason)
		Expect(result).To(BeTrue())
		Expect(err).ToNot(HaveOccurred())
	})

	It("should fail matching an ErrorWithReason", func() {
		m := testing.ErrorWithReason(anyError)

		errWithReason := errors.Wrap(anotherError, errors.Code("code2"))

		result, err := m.Match(errWithReason)
		Expect(result).To(BeFalse())
		Expect(err).ToNot(HaveOccurred())
	})

	It("should fail matching a deep ErrorWithReason", func() {
		m := testing.ErrorWithReason(anyError)

		errWithReason := errors.Wrap(anotherError, errors.Code("code2"), errors.Message("message1"), errors.Http(404))

		result, err := m.Match(errWithReason)
		Expect(result).To(BeFalse())
		Expect(err).ToNot(HaveOccurred())
	})

	It("should fail matching an error chain with no ErrorWithReason", func() {
		m := testing.ErrorWithReason(anyError)

		errWithReason := anotherError

		result, err := m.Match(errWithReason)
		Expect(result).To(BeFalse())
		Expect(err).ToNot(HaveOccurred())
	})

	It("should fail matching when the actual is not an error", func() {
		m := testing.ErrorWithReason(anyError)

		result, err := m.Match("not an error")
		Expect(result).To(BeFalse())
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("`actual` is not an `error`"))
	})

	It("should fail matching with actual = nil", func() {
		m := testing.ErrorWithReason(anyError)

		result, err := m.Match(nil)
		Expect(result).To(BeFalse())
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("`actual` is not an `error`"))
	})

	It("should run through the Gomega", func() {
		errWithReason := errors.Wrap(anyError, errors.Code("code1"), errors.Message("message1"), errors.Http(404))

		Expect(errWithReason).To(testing.ErrorWithReason(anyError))
	})

	It("should run through the Gomega (negating)", func() {
		errWithReason := errors.Wrap(anyError, errors.Message("message1"), errors.Http(404))

		Expect(errWithReason).ToNot(testing.ErrorWithReason(anotherError))
	})

	It("should generate a failure message with different error codes", func() {
		errWithReason := errors.Wrap(anyError, errors.Code("code1"), errors.Message("message1"), errors.Http(404))

		m := testing.ErrorWithReason(anotherError)
		message := m.FailureMessage(errWithReason)
		Expect(message).To(ContainSubstring("to have any reason equal"))
	})

	It("should generate a failure message negating the code found", func() {
		errWithReason := errors.Wrap(anyError, errors.Code("code1"), errors.Message("message1"), errors.Http(404))

		m := testing.ErrorWithReason(anyError)
		message := m.NegatedFailureMessage(errWithReason)
		Expect(message).To(ContainSubstring("not to have any reason equal"))
	})
})
