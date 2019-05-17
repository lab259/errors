package gqlerrors_test

import (
	"github.com/lab259/errors/gqlerrors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
	"gopkg.in/gavv/httpexpect.v1"
)

var _ = Describe("GraphQL Testing Utils", func() {
	Describe("Extensions", func() {
		Describe("Validate", func() {

			It("should check the return extension validate", func() {
				jsonData := httpexpect.NewObject(&gqlerrors.HttpGomegaFail{}, map[string]interface{}{
					"data": map[string]interface{}{"mutate": nil},
					"errors": []map[string]interface{}{
						{
							"extensions": map[string]interface{}{
								"module": "users",
								"errors": map[string]interface{}{
									"name": []interface{}{
										"required",
									},
								},
							},
						},
					},
				})

				a := gqlerrors.ErrWithGraphQLValidate("mutate", "name", "required")
				ok, err := a.Match(jsonData)
				Expect(ok).To(BeTrue())
				Expect(err).ToNot(HaveOccurred())
			})

			It("should fail checking the return validate when mutate not matcher", func() {
				jsonData := httpexpect.NewObject(&gqlerrors.HttpGomegaFail{}, map[string]interface{}{
					"data": map[string]interface{}{"mutate": nil},
					"errors": []map[string]interface{}{
						{
							"extensions": map[string]interface{}{
								"module": "users",
								"errors": map[string]interface{}{
									"name": []interface{}{
										"required",
									},
								},
							},
						},
					},
				})

				a := gqlerrors.ErrWithGraphQLValidate("mutateInvalid","email", "email")
				ok, err := a.Match(jsonData)
				Expect(ok).To(BeFalse())
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("expected mutate or query name [mutate] not is equal [mutateInvalid]"))
			})

			It("should fail checking the return when error not matcher validate", func() {
				jsonData := httpexpect.NewObject(&gqlerrors.HttpGomegaFail{}, map[string]interface{}{
					"data": map[string]interface{}{"mutate": nil},
					"errors": []map[string]interface{}{
						{
							"extensions": map[string]interface{}{
								"module": "users",
								"errors": map[string]interface{}{
									"email": []interface{}{
										"required",
									},
								},
							},
						},
					},
				})

				a := gqlerrors.ErrWithGraphQLValidate("mutate", "email", "email")
				ok, err := a.Match(jsonData)
				Expect(ok).To(BeFalse())
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("errors not containing key [email] on the [\"\"]"))
			})

			It("should fail checking the return when error not matcher errors", func() {
				jsonData := httpexpect.NewObject(&gqlerrors.HttpGomegaFail{}, map[string]interface{}{
					"data": map[string]interface{}{"mutate": nil},
					"errors": []map[string]interface{}{
						{
							"extensions": map[string]interface{}{
								"module": "users",
							},
						},
					},
				})

				a := gqlerrors.ErrWithGraphQLValidate("mutate", "name", "required")
				ok, err := a.Match(jsonData)
				Expect(ok).To(BeFalse())
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("couldn't have key `validation`"))
			})

			It("should checking failure validate", func() {
				mutate := "mutate"
				field := "name"
				rule := "required"
				a := gqlerrors.ErrWithGraphQLValidate(mutate, "name", "required")
				failureValidate := a.FailureMessage("mutate")
				fValidate := format.Message(mutate, "to have any validation equal field [", field, "] and rules [", rule, "]")
				Expect(failureValidate).To(Equal(fValidate))
			})

			It("should checking negative failure validate", func() {
				mutate := "mutate"
				field := "name"
				rule := "required"
				a := gqlerrors.ErrWithGraphQLValidate(mutate, field, rule)
				failureValidate := a.NegatedFailureMessage("mutate")
				fValidate := format.Message(mutate, "to have any validation equal field [", field, "] and rules [", rule, "]")
				Expect(failureValidate).To(Equal(fValidate))
			})

		})
	})
})
