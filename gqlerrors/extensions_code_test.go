package gqlerrors_test

import (
	"github.com/lab259/errors"
	"github.com/lab259/errors/gqlerrors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
	"gopkg.in/gavv/httpexpect.v1"
)

var _ = Describe("GraphQL Extensions Code", func() {
	It("should check the return code", func() {
		jsonData := httpexpect.NewObject(&httpGomegaFail{}, map[string]interface{}{
			"data": map[string]interface{}{"mutate": nil},
			"errors": []map[string]interface{}{
				{
					"extensions": map[string]interface{}{
						"code": "validation",
					},
				},
			},
		})

		a := gqlerrors.ErrWithGraphQLCode("mutate", "validation")
		ok, err := a.Match(jsonData)
		Expect(ok).To(BeTrue())
		Expect(err).ToNot(HaveOccurred())
	})

	It("should fail check the return code when input not object", func() {
		a := gqlerrors.ErrWithGraphQLCode("mutate", "validate")
		ok, err := a.Match("it is not an object of the type `*httpexpect.Object`")
		Expect(ok).To(BeFalse())
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("`actual` is not an json object"))
	})

	It("should fail when not matcher extension error", func() {
		m := gqlerrors.ErrWithGraphQLCode("mutate", "validate")

		ok, err := m.Match(&httpexpect.Object{})
		Expect(ok).To(BeFalse())
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("expected an error is not `&{{<nil> %!s(bool=false)} map[]}`"))
	})

	It("should fail checking the return code when mutate not matcher", func() {
		jsonData := httpexpect.NewObject(&httpGomegaFail{}, map[string]interface{}{
			"data": map[string]interface{}{"mutate": nil},
			"errors": []map[string]interface{}{
				{
					"extensions": map[string]interface{}{
						"code": "validation",
					},
				},
			},
		})

		code := errors.Code("invalid")
		a := gqlerrors.ErrWithGraphQLCode("mutateInvalid", code)
		ok, err := a.Match(jsonData)
		Expect(ok).To(BeFalse())
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("expected mutate or query name [mutate] not is equal [mutateInvalid]"))
	})

	It("should fail checking the return when error not contains code", func() {
		jsonData := httpexpect.NewObject(&httpGomegaFail{}, map[string]interface{}{
			"data": map[string]interface{}{"mutate": nil},
			"errors": []map[string]interface{}{
				{
					"extensions": map[string]interface{}{
						"module": "accounts",
					},
				},
			},
		})

		a := gqlerrors.ErrWithGraphQLCode("mutate", "users")
		ok, err := a.Match(jsonData)
		Expect(ok).To(BeFalse())
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("couldn't have key `code` map[\"module\":\"accounts\"]"))
	})

	It("should fail to decode when error not matcher", func() {
		jsonData := httpexpect.NewObject(&httpGomegaFail{}, map[string]interface{}{
			"data": map[string]interface{}{"mutate": nil},
			"errors": map[string]interface{}{
				"extensions": map[string]interface{}{
					"code": "invalid-account-id",
				},
			},
		})

		code := errors.Code("invalid-account-id")
		a := gqlerrors.ErrWithGraphQLCode("mutate", code)
		ok, err := a.Match(jsonData)
		Expect(ok).To(BeFalse())
		Expect(err).To(HaveOccurred())
	})

	It("should fail checking the return when error not matcher code option", func() {
		jsonData := httpexpect.NewObject(&httpGomegaFail{}, map[string]interface{}{
			"data": map[string]interface{}{"mutate": nil},
			"errors": []map[string]interface{}{
				{
					"extensions": map[string]interface{}{
						"code": "validation",
					},
				},
			},
		})

		op := errors.Code("new-code")
		a := gqlerrors.ErrWithGraphQLCode("mutate", op)
		ok, err := a.Match(jsonData)
		Expect(ok).To(BeFalse())
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("code [new-code] not equal [validation]"))
	})

	It("should fail checking the return when error not matcher code option nil", func() {
		jsonData := httpexpect.NewObject(&httpGomegaFail{}, map[string]interface{}{
			"data": map[string]interface{}{"mutate": nil},
			"errors": []map[string]interface{}{
				{
					"extensions": map[string]interface{}{
						"code": "validation",
					},
				},
			},
		})

		a := gqlerrors.ErrWithGraphQLCode("mutate", nil)
		ok, err := a.Match(jsonData)
		Expect(ok).To(BeFalse())
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("the code cannot be null"))
	})

	It("should fail checking the return when error not matcher code text", func() {
		jsonData := httpexpect.NewObject(&httpGomegaFail{}, map[string]interface{}{
			"data": map[string]interface{}{"mutate": nil},
			"errors": []map[string]interface{}{
				{
					"extensions": map[string]interface{}{
						"code": "validation",
					},
				},
			},
		})

		a := gqlerrors.ErrWithGraphQLCode("mutate", "graphql")
		ok, err := a.Match(jsonData)
		Expect(ok).To(BeFalse())
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("code [graphql] not equal [validation]"))
	})

	It("should checking failure message", func() {
		mutate := "mutate"
		code := "graphql"
		a := gqlerrors.ErrWithGraphQLCode(mutate, code)
		failureMessage := a.FailureMessage("mutate")
		fMessage := format.Message(mutate, "to have any code equal field", code)
		Expect(failureMessage).To(Equal(fMessage))
	})

	It("should checking negative failure message", func() {
		mutate := "mutate"
		code := "graphql"
		a := gqlerrors.ErrWithGraphQLCode(mutate, code)
		failureMessage := a.NegatedFailureMessage("mutate")
		fMessage := format.Message(mutate, "to have any code equal field", code)
		Expect(failureMessage).To(Equal(fMessage))
	})

})
