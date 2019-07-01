package testing_test

import (
	"github.com/lab259/errors/v2"
	"github.com/lab259/errors/v2/testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ErrorWithCodeMatcher Test Suite", func() {
	It("should initialize the matcher", func() {
		m := testing.ErrorWithCode("code1")
		Expect(m.Expected).To(Equal("code1"))
	})

	It("should match an ErrorWithCode", func() {
		m := testing.ErrorWithCode("code1")

		errCode := errors.Wrap(errors.New("error1"), errors.Code("code1"))

		result, err := m.Match(errCode)
		Expect(result).To(BeTrue())
		Expect(err).ToNot(HaveOccurred())
	})

	It("should match a deep ErrorWithCode", func() {
		m := testing.ErrorWithCode("code1")

		errCode := errors.Wrap(errors.New("error1"), errors.Code("code1"), errors.Message("message1"), errors.Http(404))

		result, err := m.Match(errCode)
		Expect(result).To(BeTrue())
		Expect(err).ToNot(HaveOccurred())
	})

	It("should fail matching an ErrorWithCode", func() {
		m := testing.ErrorWithCode("code1")

		errCode := errors.Wrap(errors.New("error1"), errors.Code("code2"))

		result, err := m.Match(errCode)
		Expect(result).To(BeFalse())
		Expect(err).ToNot(HaveOccurred())
	})

	It("should fail matching a deep ErrorWithCode", func() {
		m := testing.ErrorWithCode("code1")

		errCode := errors.Wrap(errors.New("error1"), errors.Code("code2"), errors.Message("message1"), errors.Http(404))

		result, err := m.Match(errCode)
		Expect(result).To(BeFalse())
		Expect(err).ToNot(HaveOccurred())
	})

	It("should fail matching an error chain with no ErrorWithCode", func() {
		m := testing.ErrorWithCode("code1")

		errCode := errors.Wrap(errors.New("error1"), errors.Message("message1"), errors.Http(404))

		result, err := m.Match(errCode)
		Expect(result).To(BeFalse())
		Expect(err).ToNot(HaveOccurred())
	})

	It("should fail matching when the actual is not an error", func() {
		m := testing.ErrorWithCode("code1")

		result, err := m.Match("not an error")
		Expect(result).To(BeFalse())
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("`actual` is not an `error`"))
	})

	It("should run through the Gomega", func() {
		errCode := errors.Wrap(errors.New("error1"), errors.Code("code1"), errors.Message("message1"), errors.Http(404))

		Expect(errCode).To(testing.ErrorWithCode("code1"))
	})

	It("should run through the Gomega (negating)", func() {
		errCode := errors.Wrap(errors.New("error1"), errors.Message("message1"), errors.Http(404))

		Expect(errCode).ToNot(testing.ErrorWithCode("code1"))
	})

	It("should generate a failure message with different error codes", func() {
		errCode := errors.Wrap(errors.New("error1"), errors.Code("code1"), errors.Message("message1"), errors.Http(404))

		m := testing.ErrorWithCode("code2")
		message := m.FailureMessage(errCode)
		Expect(message).To(ContainSubstring("code2"))
		Expect(message).To(ContainSubstring("to have code"))
		Expect(message).To(ContainSubstring("code1"))
	})

	It("should generate a failure message for non existing ErrorWithCode in the error chain", func() {
		errCode := errors.Wrap(errors.New("error1"), errors.Message("message1"), errors.Http(404))

		m := testing.ErrorWithCode("code1")
		message := m.FailureMessage(errCode)

		Expect(message).To(ContainSubstring("does not have ErrorWithCode"))
		Expect(message).To(ContainSubstring("code1"))
	})

	It("should generate a failure message negating the code found", func() {
		errCode := errors.Wrap(errors.New("error1"), errors.Code("code1"), errors.Message("message1"), errors.Http(404))

		m := testing.ErrorWithCode("code1")
		message := m.NegatedFailureMessage(errCode)
		Expect(message).To(ContainSubstring("not to have code"))
		Expect(message).To(ContainSubstring("code1"))
	})
})
