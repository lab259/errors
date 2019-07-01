package errors_test

import (
	lerrors "github.com/lab259/errors/v2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
)

var _ = Describe("ErrorWithCode", func() {
	It("should wrap a ReportableError", func() {
		nerr := errors.New("test")
		err := lerrors.WrapCode(nerr, "error code")
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal("error code: test"))

		reportableErr, ok := err.(lerrors.ErrorWithCode)
		Expect(ok).To(BeTrue())
		Expect(reportableErr.Code()).To(Equal("error code"))

		errResponse := NewMockErrorResponse()

		aggErr, ok := err.(lerrors.ErrorResponseAggregator)
		Expect(ok).To(BeTrue())
		aggErr.AppendData(errResponse)
		Expect(errResponse.Data).To(HaveKeyWithValue("code", "error code"))

		errWithReason, ok := err.(lerrors.Wrapper)
		Expect(ok).To(BeTrue())
		Expect(errWithReason.Unwrap()).To(Equal(nerr))
	})
})
