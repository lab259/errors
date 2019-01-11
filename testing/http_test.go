package testing_test

import (
	"github.com/lab259/errors"
	"github.com/lab259/errors/testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/valyala/fasthttp"
)

var _ = Describe("HttpErrorMatcher Test Suite", func() {
	It("should initialize the matcher", func() {
		m := testing.HttpStatus(fasthttp.StatusNotFound)
		Expect(m.Expected).To(Equal(fasthttp.StatusNotFound))
	})

	It("should match an HttpError", func() {
		m := testing.HttpStatus(fasthttp.StatusNotFound)

		errCode := errors.Wrap(errors.New("error1"), errors.Http(fasthttp.StatusNotFound))

		result, err := m.Match(errCode)
		Expect(result).To(BeTrue())
		Expect(err).ToNot(HaveOccurred())
	})

	It("should match a deep HttpError", func() {
		m := testing.HttpStatus(fasthttp.StatusNotFound)

		errCode := errors.Wrap(errors.New("error1"), errors.Http(fasthttp.StatusNotFound), errors.Message("message1"), errors.Http(404))

		result, err := m.Match(errCode)
		Expect(result).To(BeTrue())
		Expect(err).ToNot(HaveOccurred())
	})

	It("should fail matching an HttpError", func() {
		m := testing.HttpStatus(fasthttp.StatusNotFound)

		errCode := errors.Wrap(errors.New("error1"), errors.Http(fasthttp.StatusInternalServerError))

		result, err := m.Match(errCode)
		Expect(result).To(BeFalse())
		Expect(err).ToNot(HaveOccurred())
	})

	It("should fail matching a deep HttpError", func() {
		m := testing.HttpStatus(fasthttp.StatusNotFound)

		errCode := errors.Wrap(errors.New("error1"), errors.Http(fasthttp.StatusInternalServerError), errors.Message("message1"), errors.Code("error2"))

		result, err := m.Match(errCode)
		Expect(result).To(BeFalse())
		Expect(err).ToNot(HaveOccurred())
	})

	It("should fail matching an error chain with no HttpError", func() {
		m := testing.HttpStatus(fasthttp.StatusNotFound)

		errCode := errors.Wrap(errors.New("error1"), errors.Message("message1"), errors.Code("error2"))

		result, err := m.Match(errCode)
		Expect(result).To(BeFalse())
		Expect(err).ToNot(HaveOccurred())
	})

	It("should fail matching when the actual is not an error", func() {
		m := testing.HttpStatus(fasthttp.StatusNotFound)

		result, err := m.Match("not an error")
		Expect(result).To(BeFalse())
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("`actual` is not an `error`"))
	})

	It("should run through the Gomega", func() {
		errCode := errors.Wrap(errors.New("error1"), errors.Http(fasthttp.StatusNotFound), errors.Message("message1"), errors.Http(404))

		Expect(errCode).To(testing.HttpStatus(fasthttp.StatusNotFound))
	})

	It("should run through the Gomega (negating)", func() {
		errCode := errors.Wrap(errors.New("error1"), errors.Message("message1"), errors.Code("error2"))

		Expect(errCode).ToNot(testing.HttpStatus(fasthttp.StatusNotFound))
	})

	It("should generate a failure message with different error codes", func() {
		errCode := errors.Wrap(errors.New("error1"), errors.Http(fasthttp.StatusNotFound), errors.Message("message1"), errors.Http(404))

		m := testing.HttpStatus(fasthttp.StatusInternalServerError)
		message := m.FailureMessage(errCode)
		Expect(message).To(ContainSubstring("500"))
		Expect(message).To(ContainSubstring("to have status code"))
		Expect(message).To(ContainSubstring("404"))
	})

	It("should generate a failure message for non existing HttpError in the error chain", func() {
		errCode := errors.Wrap(errors.New("error1"), errors.Message("message1"), errors.Code("error1"))

		m := testing.HttpStatus(fasthttp.StatusNotFound)
		message := m.FailureMessage(errCode)

		Expect(message).To(ContainSubstring("does not have HttpError"))
		Expect(message).To(ContainSubstring("404"))
	})

	It("should generate a failure message negating the code found", func() {
		errCode := errors.Wrap(errors.New("error1"), errors.Http(fasthttp.StatusNotFound), errors.Message("message1"), errors.Http(404))

		m := testing.HttpStatus(fasthttp.StatusNotFound)
		message := m.NegatedFailureMessage(errCode)
		Expect(message).To(ContainSubstring("not to have status code"))
		Expect(message).To(ContainSubstring("404"))
	})
})
