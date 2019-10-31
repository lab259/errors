package gqlerrors_test

import (
	"encoding/json"

	"github.com/99designs/gqlgen/client"
	"github.com/lab259/errors/v2"
	"github.com/lab259/errors/v2/gqlerrors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vektah/gqlparser/gqlerror"
	"gopkg.in/gavv/httpexpect.v1"
)

var _ = Describe("GraphQL Extensions", func() {
	Describe("HaveMessage", func() {
		When("Gqlerror", func() {

			It("should match", func() {
				gqlerr := gqlerror.Error{
					Extensions: map[string]interface{}{
						"message": "name is required",
					},
				}

				a := gqlerrors.HaveMessage("name is required")
				ok, err := a.Match(gqlerr)
				Expect(ok).To(BeTrue())
				Expect(err).ToNot(HaveOccurred())
			})

			It("should match using Option", func() {
				gqlerr := gqlerror.Error{
					Extensions: map[string]interface{}{
						"message": "name is required",
					},
				}

				a := gqlerrors.HaveMessage(errors.Message("name is required"))
				ok, err := a.Match(gqlerr)
				Expect(ok).To(BeTrue())
				Expect(err).ToNot(HaveOccurred())
			})

			It("should match using ErrorWithMessage", func() {
				gqlerr := gqlerror.Error{
					Extensions: map[string]interface{}{
						"message": "name is required",
					},
				}

				a := gqlerrors.HaveMessage(errors.Message("name is required")(nil))
				ok, err := a.Match(gqlerr)
				Expect(ok).To(BeTrue())
				Expect(err).ToNot(HaveOccurred())
			})

			It("should match (pointer)", func() {
				gqlerr := &gqlerror.Error{
					Extensions: map[string]interface{}{
						"message": "name is required",
					},
				}

				a := gqlerrors.HaveMessage("name is required")
				ok, err := a.Match(gqlerr)
				Expect(ok).To(BeTrue())
				Expect(err).ToNot(HaveOccurred())
			})

			It("should match (json.RawMessage)", func() {
				gqlerr := json.RawMessage(`[{"extensions": {"message": "name is required"}}]`)
				a := gqlerrors.HaveMessage("name is required")
				ok, err := a.Match(gqlerr)
				Expect(ok).To(BeTrue())
				Expect(err).ToNot(HaveOccurred())
			})

			It("should match (client.RawJsonError)", func() {
				gqlerr := client.RawJsonError{
					RawMessage: json.RawMessage(`[{"extensions": {"message": "name is required"}}]`),
				}
				a := gqlerrors.HaveMessage("name is required")
				ok, err := a.Match(gqlerr)
				Expect(ok).To(BeTrue())
				Expect(err).ToNot(HaveOccurred())
			})

			It("should not match", func() {
				gqlerr := gqlerror.Error{
					Extensions: map[string]interface{}{
						"message": "name is required",
					},
				}

				a := gqlerrors.HaveMessage("not-found")
				ok, err := a.Match(gqlerr)
				Expect(ok).To(BeFalse())
				Expect(err).ToNot(HaveOccurred())
			})

			It("should fail without extensions", func() {
				gqlerr := gqlerror.Error{
					Extensions: map[string]interface{}{},
				}

				a := gqlerrors.HaveMessage("name is required")
				ok, err := a.Match(gqlerr)
				Expect(ok).To(BeFalse())
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Message extension not found in"))
			})

			It("should fail with wrong errors.Option", func() {
				gqlerr := gqlerror.Error{
					Extensions: map[string]interface{}{
						"message": "name is required",
					},
				}

				a := gqlerrors.HaveMessage(errors.Module("name is required"))
				ok, err := a.Match(gqlerr)
				Expect(ok).To(BeFalse())
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("HaveMessage matcher only support an `errors.Message` option"))
			})

			It("should fail with wrong actual type", func() {
				a := gqlerrors.HaveMessage("name is required")
				ok, err := a.Match(26)
				Expect(ok).To(BeFalse())
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("HaveMessage matcher does not know how to handle int"))
			})

			It("should fail with wrong expected type", func() {
				gqlerr := gqlerror.Error{
					Extensions: map[string]interface{}{
						"message": "name is required",
					},
				}

				a := gqlerrors.HaveMessage(26)
				ok, err := a.Match(gqlerr)
				Expect(ok).To(BeFalse())
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("HaveMessage matcher does not know how to assert int"))
			})
		})

		Describe("ContainMessage", func() {

			When("using error's message", func() {
				It("should match", func() {
					gqlerr := gqlerror.Error{
						Message: "error: name is required",
					}

					a := gqlerrors.ContainMessage("name is required")
					ok, err := a.Match(gqlerr)
					Expect(ok).To(BeTrue())
					Expect(err).ToNot(HaveOccurred())
				})

				It("should match using Option", func() {
					gqlerr := gqlerror.Error{
						Message: "error: name is required",
					}

					a := gqlerrors.ContainMessage(errors.Message("name is required"))
					ok, err := a.Match(gqlerr)
					Expect(ok).To(BeTrue())
					Expect(err).ToNot(HaveOccurred())
				})

				It("should match using ErrorWithMessage", func() {
					gqlerr := gqlerror.Error{
						Message: "error: name is required",
					}

					a := gqlerrors.ContainMessage(errors.Message("name is required")(nil))
					ok, err := a.Match(gqlerr)
					Expect(ok).To(BeTrue())
					Expect(err).ToNot(HaveOccurred())
				})

				It("should match (pointer)", func() {
					gqlerr := &gqlerror.Error{
						Message: "error: name is required",
					}

					a := gqlerrors.ContainMessage("name is required")
					ok, err := a.Match(gqlerr)
					Expect(ok).To(BeTrue())
					Expect(err).ToNot(HaveOccurred())
				})

				It("should match (json.RawMessage)", func() {
					gqlerr := json.RawMessage(`[{"message": "name is required"}]`)
					a := gqlerrors.ContainMessage("name is required")
					ok, err := a.Match(gqlerr)
					Expect(ok).To(BeTrue())
					Expect(err).ToNot(HaveOccurred())
				})

				It("should match (client.RawJsonError)", func() {
					gqlerr := client.RawJsonError{
						RawMessage: json.RawMessage(`[{"message": "name is required"}]`),
					}
					a := gqlerrors.ContainMessage("name is required")
					ok, err := a.Match(gqlerr)
					Expect(ok).To(BeTrue())
					Expect(err).ToNot(HaveOccurred())
				})

				It("should not match", func() {
					gqlerr := gqlerror.Error{
						Message: "error: name is required",
					}

					a := gqlerrors.ContainMessage("not-found")
					ok, err := a.Match(gqlerr)
					Expect(ok).To(BeFalse())
					Expect(err).ToNot(HaveOccurred())
				})

				It("should fail with wrong errors.Option", func() {
					gqlerr := gqlerror.Error{
						Message: "error: name is required",
					}

					a := gqlerrors.ContainMessage(errors.Module("name is required"))
					ok, err := a.Match(gqlerr)
					Expect(ok).To(BeFalse())
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("ContainMessage matcher only support an `errors.Message` option"))
				})

				It("should fail with wrong actual type", func() {
					a := gqlerrors.ContainMessage("name is required")
					ok, err := a.Match(26)
					Expect(ok).To(BeFalse())
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("ContainMessage matcher does not know how to handle int"))
				})

				It("should fail with wrong expected type", func() {
					gqlerr := gqlerror.Error{
						Message: "error: name is required",
					}

					a := gqlerrors.ContainMessage(26)
					ok, err := a.Match(gqlerr)
					Expect(ok).To(BeFalse())
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("ContainMessage matcher does not know how to assert int"))
				})
			})

			When("using error's extensions", func() {
				It("should match", func() {
					gqlerr := gqlerror.Error{
						Extensions: map[string]interface{}{
							"message": "name is required",
						},
					}

					a := gqlerrors.ContainMessage("name is required")
					ok, err := a.Match(gqlerr)
					Expect(ok).To(BeTrue())
					Expect(err).ToNot(HaveOccurred())
				})

				It("should match using Option", func() {
					gqlerr := gqlerror.Error{
						Extensions: map[string]interface{}{
							"message": "name is required",
						},
					}

					a := gqlerrors.ContainMessage(errors.Message("name is required"))
					ok, err := a.Match(gqlerr)
					Expect(ok).To(BeTrue())
					Expect(err).ToNot(HaveOccurred())
				})

				It("should match using ErrorWithMessage", func() {
					gqlerr := gqlerror.Error{
						Extensions: map[string]interface{}{
							"message": "name is required",
						},
					}

					a := gqlerrors.ContainMessage(errors.Message("name is required")(nil))
					ok, err := a.Match(gqlerr)
					Expect(ok).To(BeTrue())
					Expect(err).ToNot(HaveOccurred())
				})

				It("should match (pointer)", func() {
					gqlerr := &gqlerror.Error{
						Extensions: map[string]interface{}{
							"message": "name is required",
						},
					}

					a := gqlerrors.ContainMessage("name is required")
					ok, err := a.Match(gqlerr)
					Expect(ok).To(BeTrue())
					Expect(err).ToNot(HaveOccurred())
				})

				It("should match (json.RawMessage)", func() {
					gqlerr := json.RawMessage(`[{"extensions": {"message": "name is required"}}]`)
					a := gqlerrors.ContainMessage("name is required")
					ok, err := a.Match(gqlerr)
					Expect(ok).To(BeTrue())
					Expect(err).ToNot(HaveOccurred())
				})

				It("should match (client.RawJsonError)", func() {
					gqlerr := client.RawJsonError{
						RawMessage: json.RawMessage(`[{"extensions": {"message": "name is required"}}]`),
					}
					a := gqlerrors.ContainMessage("name is required")
					ok, err := a.Match(gqlerr)
					Expect(ok).To(BeTrue())
					Expect(err).ToNot(HaveOccurred())
				})

				It("should not match", func() {
					gqlerr := gqlerror.Error{
						Extensions: map[string]interface{}{
							"message": "name is required",
						},
					}

					a := gqlerrors.ContainMessage("not-found")
					ok, err := a.Match(gqlerr)
					Expect(ok).To(BeFalse())
					Expect(err).ToNot(HaveOccurred())
				})

				It("should fail without extensions", func() {
					gqlerr := gqlerror.Error{
						Extensions: map[string]interface{}{},
					}

					a := gqlerrors.ContainMessage("name is required")
					ok, err := a.Match(gqlerr)
					Expect(ok).To(BeFalse())
					Expect(err).ToNot(HaveOccurred())
				})

				It("should fail with wrong errors.Option", func() {
					gqlerr := gqlerror.Error{
						Extensions: map[string]interface{}{
							"message": "name is required",
						},
					}

					a := gqlerrors.ContainMessage(errors.Module("name is required"))
					ok, err := a.Match(gqlerr)
					Expect(ok).To(BeFalse())
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("ContainMessage matcher only support an `errors.Message` option"))
				})

				It("should fail with wrong actual type", func() {
					a := gqlerrors.ContainMessage("name is required")
					ok, err := a.Match(26)
					Expect(ok).To(BeFalse())
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("ContainMessage matcher does not know how to handle int"))
				})

				It("should fail with wrong expected type", func() {
					gqlerr := gqlerror.Error{
						Extensions: map[string]interface{}{
							"message": "name is required",
						},
					}

					a := gqlerrors.ContainMessage(26)
					ok, err := a.Match(gqlerr)
					Expect(ok).To(BeFalse())
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("ContainMessage matcher does not know how to assert int"))
				})
			})
		})

		When("GraphQL", func() {
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

				a := gqlerrors.HaveMessage("name is required")
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

				a := gqlerrors.ContainMessage(message)
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

				a := gqlerrors.ContainMessage(message)
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

				a := gqlerrors.ContainMessage("email is required")
				ok, err := a.Match(jsonData)
				Expect(ok).To(BeTrue())
				Expect(err).ToNot(HaveOccurred())
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

				a := gqlerrors.HaveMessage("users")
				ok, err := a.Match(jsonData)
				Expect(ok).To(BeFalse())
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Message extension not found"))
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
				a := gqlerrors.HaveMessage(op)
				ok, err := a.Match(jsonData)
				Expect(ok).To(BeFalse())
				Expect(err).ToNot(HaveOccurred())
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

				a := gqlerrors.HaveMessage("graphql")
				ok, err := a.Match(jsonData)
				Expect(ok).To(BeFalse())
				Expect(err).ToNot(HaveOccurred())
			})

			It("should checking failure message", func() {
				message := "graphql"
				a := gqlerrors.HaveMessage(message)
				failureMessage := a.FailureMessage("mutate")
				Expect(failureMessage).To(ContainSubstring("to have Message extension equals to"))
			})

			It("should checking negative failure message", func() {
				message := "graphql"
				a := gqlerrors.HaveMessage(message)
				failureMessage := a.NegatedFailureMessage("mutate")
				Expect(failureMessage).To(ContainSubstring("to not have Message extension equals to"))
			})
		})
	})
})
