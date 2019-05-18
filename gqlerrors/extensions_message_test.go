package gqlerrors_test

import (
	"github.com/lab259/errors"
	"github.com/lab259/errors/gqlerrors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
	"gopkg.in/gavv/httpexpect.v1"
)

var _ = Describe("GraphQL Extensions Message", func() {
	It("should check the return extension message", func() {
		jsonData := httpexpect.NewObject(&httpGomegaFail{}, map[string]interface{}{
			"data": map[string]interface{}{"mutate": nil},
			"errors": []map[string]interface{}{
				{
					"extensions": map[string]interface{}{
						"message": "name is required",
					},
				},
			},
		})

		a := gqlerrors.ErrWithGraphQLMessage("mutate", "name is required")
		ok, err := a.Match(jsonData)
		Expect(ok).To(BeTrue())
		Expect(err).ToNot(HaveOccurred())
	})

	It("should check the return message", func() {
		message := errors.Message("name is required")
		jsonData := httpexpect.NewObject(&httpGomegaFail{}, map[string]interface{}{
			"data": map[string]interface{}{"mutate": nil},
			"errors": []map[string]interface{}{
				{
					"extensions": map[string]interface{}{
						"code": "validation",
					},
					"message": "name is required",
				},
			},
		})

		a := gqlerrors.ErrWithGraphQLMessage("mutate", message)
		ok, err := a.Match(jsonData)
		Expect(ok).To(BeTrue())
		Expect(err).ToNot(HaveOccurred())
	})

	It("should check the return message type error", func() {
		message := errors.New("name is required")
		jsonData := httpexpect.NewObject(&httpGomegaFail{}, map[string]interface{}{
			"data": map[string]interface{}{"mutate": nil},
			"errors": []map[string]interface{}{
				{
					"extensions": map[string]interface{}{
						"code": "validation",
					},
					"message": "name is required",
				},
			},
		})

		a := gqlerrors.ErrWithGraphQLMessage("mutate", message)
		ok, err := a.Match(jsonData)
		Expect(ok).To(BeTrue())
		Expect(err).ToNot(HaveOccurred())
	})

	It("should check the return message type string", func() {
		jsonData := httpexpect.NewObject(&httpGomegaFail{}, map[string]interface{}{
			"data": map[string]interface{}{"mutate": nil},
			"errors": []map[string]interface{}{
				{
					"extensions": map[string]interface{}{
						"code": "validation",
					},
					"message": "email is required",
				},
			},
		})

		a := gqlerrors.ErrWithGraphQLMessage("mutate", "email is required")
		ok, err := a.Match(jsonData)
		Expect(ok).To(BeTrue())
		Expect(err).ToNot(HaveOccurred())
	})

	It("should fail checking the return message when mutate not matcher", func() {
		jsonData := httpexpect.NewObject(&httpGomegaFail{}, map[string]interface{}{
			"data": map[string]interface{}{"mutate": nil},
			"errors": []map[string]interface{}{
				{
					"extensions": map[string]interface{}{
						"message": "validation",
					},
				},
			},
		})

		message := errors.Message("invalid")
		a := gqlerrors.ErrWithGraphQLMessage("mutateInvalid", message)
		ok, err := a.Match(jsonData)
		Expect(ok).To(BeFalse())
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("expected mutate or query name [mutate] not is equal [mutateInvalid]"))
	})

	It("should fail checking the return when error not contains message", func() {
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

		a := gqlerrors.ErrWithGraphQLMessage("mutate", "users")
		ok, err := a.Match(jsonData)
		Expect(ok).To(BeFalse())
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("the field message not found"))
	})

	It("should fail checking the return when error not matcher message option", func() {
		jsonData := httpexpect.NewObject(&httpGomegaFail{}, map[string]interface{}{
			"data": map[string]interface{}{"mutate": nil},
			"errors": []map[string]interface{}{
				{
					"extensions": map[string]interface{}{
						"message": "validation",
					},
				},
			},
		})

		op := errors.Message("new-message")
		a := gqlerrors.ErrWithGraphQLMessage("mutate", op)
		ok, err := a.Match(jsonData)
		Expect(ok).To(BeFalse())
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("message [new-message] not equal [validation]"))
	})

	It("should fail checking the return when error not matcher message text", func() {
		jsonData := httpexpect.NewObject(&httpGomegaFail{}, map[string]interface{}{
			"data": map[string]interface{}{"mutate": nil},
			"errors": []map[string]interface{}{
				{
					"extensions": map[string]interface{}{
						"message": "validation",
					},
				},
			},
		})

		a := gqlerrors.ErrWithGraphQLMessage("mutate", "graphql")
		ok, err := a.Match(jsonData)
		Expect(ok).To(BeFalse())
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("message [graphql] not equal [validation]"))
	})

	It("should checking failure message", func() {
		mutate := "mutate"
		message := "graphql"
		a := gqlerrors.ErrWithGraphQLMessage(mutate, message)
		failureMessage := a.FailureMessage("mutate")
		fMessage := format.Message(mutate, "to have any message equal field", message)
		Expect(failureMessage).To(Equal(fMessage))
	})

	It("should checking negative failure message", func() {
		mutate := "mutate"
		message := "graphql"
		a := gqlerrors.ErrWithGraphQLMessage(mutate, message)
		failureMessage := a.NegatedFailureMessage("mutate")
		fMessage := format.Message(mutate, "to have any message equal field", message)
		Expect(failureMessage).To(Equal(fMessage))
	})

})
