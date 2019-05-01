package errors_test

import (
	lerrors "github.com/lab259/errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
)

var _ = Describe("HttpError", func() {
	It("should wrap a HttpError", func() {
		nerr := errors.New("test")
		err := lerrors.WrapHttp(nerr, fasthttp.StatusBadRequest)
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal("test"))
		httpErr, ok := err.(lerrors.HttpError)
		Expect(ok).To(BeTrue())
		Expect(httpErr.StatusCode()).To(Equal(fasthttp.StatusBadRequest))
		reasonErr, ok := err.(lerrors.Wrapper)
		Expect(ok).To(BeTrue())
		Expect(reasonErr.Unwrap()).To(Equal(nerr))
		aggregator, ok := err.(lerrors.ErrorResponseAggregator)
		Expect(ok).To(BeTrue())
		response := NewMockErrorResponse()
		aggregator.AppendData(response)
		Expect(response.Data).To(HaveKeyWithValue("statusCode", fasthttp.StatusBadRequest))
	})
})
