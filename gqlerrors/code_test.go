package gqlerrors_test

import (
	"encoding/json"

	"github.com/99designs/gqlgen/client"
	"gopkg.in/gavv/httpexpect.v1"

	"github.com/lab259/errors/v2"
	"github.com/lab259/errors/v2/gqlerrors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vektah/gqlparser/gqlerror"
)

var _ = Describe("GraphQL Extensions", func() {

	Describe("HaveCode", func() {
		When("Gqlerror", func() {
			It("should match", func() {
				gqlerr := gqlerror.Error{
					Extensions: map[string]interface{}{
						"code": "validation",
					},
				}

				a := gqlerrors.HaveCode("validation")
				ok, err := a.Match(gqlerr)
				Expect(ok).To(BeTrue())
				Expect(err).ToNot(HaveOccurred())
			})

			It("should match using Option", func() {
				gqlerr := gqlerror.Error{
					Extensions: map[string]interface{}{
						"code": "validation",
					},
				}

				a := gqlerrors.HaveCode(errors.Code("validation"))
				ok, err := a.Match(gqlerr)
				Expect(ok).To(BeTrue())
				Expect(err).ToNot(HaveOccurred())
			})

			It("should match using ErrorWithCode", func() {
				gqlerr := gqlerror.Error{
					Extensions: map[string]interface{}{
						"code": "validation",
					},
				}

				a := gqlerrors.HaveCode(errors.Code("validation")(nil))
				ok, err := a.Match(gqlerr)
				Expect(ok).To(BeTrue())
				Expect(err).ToNot(HaveOccurred())
			})

			It("should match (pointer)", func() {
				gqlerr := &gqlerror.Error{
					Extensions: map[string]interface{}{
						"code": "validation",
					},
				}

				a := gqlerrors.HaveCode("validation")
				ok, err := a.Match(gqlerr)
				Expect(ok).To(BeTrue())
				Expect(err).ToNot(HaveOccurred())
			})

			It("should match (json.RawMessage)", func() {
				gqlerr := json.RawMessage(`[{"extensions": {"code": "validation"}}]`)
				a := gqlerrors.HaveCode("validation")
				ok, err := a.Match(gqlerr)
				Expect(ok).To(BeTrue())
				Expect(err).ToNot(HaveOccurred())
			})

			It("should match (client.RawJsonError)", func() {
				gqlerr := client.RawJsonError{
					RawMessage: json.RawMessage(`[{"extensions": {"code": "validation"}}]`),
				}
				a := gqlerrors.HaveCode("validation")
				ok, err := a.Match(gqlerr)
				Expect(ok).To(BeTrue())
				Expect(err).ToNot(HaveOccurred())
			})

			It("should not match", func() {
				gqlerr := gqlerror.Error{
					Extensions: map[string]interface{}{
						"code": "validation",
					},
				}

				a := gqlerrors.HaveCode("not-found")
				ok, err := a.Match(gqlerr)
				Expect(ok).To(BeFalse())
				Expect(err).ToNot(HaveOccurred())
			})

			It("should fail without extensions", func() {
				gqlerr := gqlerror.Error{
					Extensions: map[string]interface{}{},
				}

				a := gqlerrors.HaveCode("validation")
				ok, err := a.Match(gqlerr)
				Expect(ok).To(BeFalse())
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Code extension not found in"))
			})

			It("should fail with wrong errors.Option", func() {
				gqlerr := gqlerror.Error{
					Extensions: map[string]interface{}{
						"code": "validation",
					},
				}

				a := gqlerrors.HaveCode(errors.Module("validation"))
				ok, err := a.Match(gqlerr)
				Expect(ok).To(BeFalse())
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("HaveCode matcher only support an `errors.Code` option"))
			})

			It("should fail with wrong actual type", func() {
				a := gqlerrors.HaveCode("validation")
				ok, err := a.Match(26)
				Expect(ok).To(BeFalse())
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("HaveCode matcher does not know how to handle int"))
			})

			It("should fail with wrong expected type", func() {
				gqlerr := gqlerror.Error{
					Extensions: map[string]interface{}{
						"code": "validation",
					},
				}

				a := gqlerrors.HaveCode(26)
				ok, err := a.Match(gqlerr)
				Expect(ok).To(BeFalse())
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("HaveCode matcher does not know how to assert int"))
			})
		})

		When("GraphQL", func() {
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

				a := gqlerrors.HaveCode("validation")
				ok, err := a.Match(jsonData)
				Expect(ok).To(BeTrue())
				Expect(err).ToNot(HaveOccurred())
			})

			It("should fail check the return code when input not object", func() {
				a := gqlerrors.HaveCode("validate")
				ok, err := a.Match("it is not an object of the type `*httpexpect.Object`")
				Expect(ok).To(BeFalse())
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("HaveCode matcher does not know how to handle string"))
			})

			It("should fail when not matcher extension error", func() {
				m := gqlerrors.HaveCode("validate")

				ok, err := m.Match(&httpexpect.Object{})
				Expect(ok).To(BeFalse())
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("HaveCode httpexpect.Object not is error *httpexpect.Object"))
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

				a := gqlerrors.HaveCode("users")
				ok, err := a.Match(jsonData)
				Expect(ok).To(BeFalse())
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Code extension not found"))
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
				a := gqlerrors.HaveCode(code)
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
				a := gqlerrors.HaveCode(op)
				ok, err := a.Match(jsonData)
				Expect(ok).To(BeFalse())
				Expect(err).ToNot(HaveOccurred())
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

				a := gqlerrors.HaveCode(nil)
				ok, err := a.Match(jsonData)
				Expect(ok).To(BeFalse())
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("HaveCode matcher does not know how to assert"))
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

				a := gqlerrors.HaveCode("graphql")
				ok, err := a.Match(jsonData)
				Expect(ok).To(BeFalse())
				Expect(err).ToNot(HaveOccurred())
			})

			It("should checking failure message", func() {
				code := "graphql"
				a := gqlerrors.HaveCode(code)
				failureMessage := a.FailureMessage("code")
				Expect(failureMessage).To(ContainSubstring("to have Code extension equals to"))
			})

			It("should checking negative failure message", func() {
				code := "graphql"
				a := gqlerrors.HaveCode(code)
				failureMessage := a.NegatedFailureMessage("mutate")
				Expect(failureMessage).To(ContainSubstring("to not have Code extension equals to"))
			})
		})
	})
})
