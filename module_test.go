package errors_test

import (
	lerrors "github.com/lab259/errors/v2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
)

var _ = Describe("ModuleError", func() {
	It("should wrap a ModuleError", func() {
		nerr := errors.New("test")
		err := lerrors.WrapModule(nerr, "test module")
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal("test module: test"))
		moduleErr, ok := err.(lerrors.ModuleError)
		Expect(ok).To(BeTrue())
		Expect(moduleErr.Module()).To(Equal("test module"))
		reasonErr, ok := err.(lerrors.Wrapper)
		Expect(ok).To(BeTrue())
		Expect(reasonErr.Unwrap()).To(Equal(nerr))
		aggregator, ok := err.(lerrors.ErrorResponseAggregator)
		Expect(ok).To(BeTrue())
		errResponse := NewMockErrorResponse()
		aggregator.AppendData(errResponse)
		Expect(errResponse.Data).To(HaveKeyWithValue("module", "test module"))
	})

	It("should wrap a ModuleError with no reason", func() {
		err := lerrors.WrapModule(nil, "test module")
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal("test module: unknown error"))
		moduleErr, ok := err.(lerrors.ModuleError)
		Expect(ok).To(BeTrue())
		Expect(moduleErr.Module()).To(Equal("test module"))
		reasonErr, ok := err.(lerrors.Wrapper)
		Expect(ok).To(BeTrue())
		Expect(reasonErr.Unwrap()).To(BeNil())
		aggregator, ok := err.(lerrors.ErrorResponseAggregator)
		Expect(ok).To(BeTrue())
		errResponse := NewMockErrorResponse()
		aggregator.AppendData(errResponse)
		Expect(errResponse.Data).To(HaveKeyWithValue("module", "test module"))
	})
})
