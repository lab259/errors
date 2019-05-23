package gqlerrors_test

import (
	"github.com/lab259/errors"
	"github.com/lab259/errors/gqlerrors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
	"gopkg.in/gavv/httpexpect.v1"
)

var _ = Describe("GraphQL Extensions Module", func() {

	It("should check the return extension module", func() {
		jsonData := httpexpect.NewObject(&httpGomegaFail{}, map[string]interface{}{
			"data": map[string]interface{}{"mutate": nil},
			"errors": []map[string]interface{}{
				{
					"extensions": map[string]interface{}{
						"module": "users",
					},
				},
			},
		})

		a := gqlerrors.ErrWithGraphQLModule("mutate", "users")
		ok, err := a.Match(jsonData)
		Expect(ok).To(BeTrue())
		Expect(err).ToNot(HaveOccurred())
	})

	It("should check the return module", func() {
		module := errors.WrapModule(nil, "users")
		jsonData := httpexpect.NewObject(&httpGomegaFail{}, map[string]interface{}{
			"data": map[string]interface{}{"mutate": nil},
			"errors": []map[string]interface{}{
				{
					"extensions": map[string]interface{}{
						"code":   "validation",
						"module": "users",
					},
				},
			},
		})

		a := gqlerrors.ErrWithGraphQLModule("mutate", module)
		ok, err := a.Match(jsonData)
		Expect(ok).To(BeTrue())
		Expect(err).ToNot(HaveOccurred())
	})

	It("should check the return module type string", func() {
		jsonData := httpexpect.NewObject(&httpGomegaFail{}, map[string]interface{}{
			"data": map[string]interface{}{"mutate": nil},
			"errors": []map[string]interface{}{
				{
					"extensions": map[string]interface{}{
						"code":   "validation",
						"module": "users",
					},
				},
			},
		})

		a := gqlerrors.ErrWithGraphQLModule("mutate", "users")
		ok, err := a.Match(jsonData)
		Expect(ok).To(BeTrue())
		Expect(err).ToNot(HaveOccurred())
	})

	It("should fail check the return module empty", func() {
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

		a := gqlerrors.ErrWithGraphQLModule("mutate", "users")
		ok, err := a.Match(jsonData)
		Expect(ok).To(BeFalse())
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("couldn't have key `module`"))
	})

	It("should fail checking the return module when mutate not matcher", func() {
		jsonData := httpexpect.NewObject(&httpGomegaFail{}, map[string]interface{}{
			"data": map[string]interface{}{"mutate": nil},
			"errors": []map[string]interface{}{
				{
					"extensions": map[string]interface{}{
						"module": "validation",
					},
				},
			},
		})

		module := errors.Module("invalid")
		a := gqlerrors.ErrWithGraphQLModule("mutateInvalid", module)
		ok, err := a.Match(jsonData)
		Expect(ok).To(BeFalse())
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("expected mutate or query name [mutate] not is equal [mutateInvalid]"))
	})

	It("should fail checking the return when error not contains module", func() {
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

		a := gqlerrors.ErrWithGraphQLModule("mutate", "users")
		ok, err := a.Match(jsonData)
		Expect(ok).To(BeFalse())
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("expected module [accounts] not equal [users]"))
	})

	It("should fail checking the return when error not matcher module option", func() {
		jsonData := httpexpect.NewObject(&httpGomegaFail{}, map[string]interface{}{
			"data": map[string]interface{}{"mutate": nil},
			"errors": []map[string]interface{}{
				{
					"extensions": map[string]interface{}{
						"module": "validation",
					},
				},
			},
		})

		option := errors.Module("new-module")
		a := gqlerrors.ErrWithGraphQLModule("mutate", option)
		ok, err := a.Match(jsonData)
		Expect(ok).To(BeFalse())
		Expect(err).To(HaveOccurred())
	})

	It("should fail checking the return when error not matcher module text", func() {
		jsonData := httpexpect.NewObject(&httpGomegaFail{}, map[string]interface{}{
			"data": map[string]interface{}{"mutate": nil},
			"errors": []map[string]interface{}{
				{
					"extensions": map[string]interface{}{
						"module": "validation",
					},
				},
			},
		})

		a := gqlerrors.ErrWithGraphQLModule("mutate", "graphql")
		ok, err := a.Match(jsonData)
		Expect(ok).To(BeFalse())
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("expected module [validation] not equal [graphql]"))
	})

	It("should checking failure module", func() {
		mutate := "mutate"
		module := "graphql"
		a := gqlerrors.ErrWithGraphQLModule(mutate, module)
		failureModule := a.FailureMessage("mutate")
		fModule := format.Message(mutate, "to have any module equal field", module)
		Expect(failureModule).To(Equal(fModule))
	})

	It("should checking negative failure module", func() {
		mutate := "mutate"
		module := "graphql"
		a := gqlerrors.ErrWithGraphQLModule(mutate, module)
		failureModule := a.NegatedFailureMessage("mutate")
		fModule := format.Message(mutate, "to have any module equal field", module)
		Expect(failureModule).To(Equal(fModule))
	})

})
