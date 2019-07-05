package gqlerrors_test

import (
	"encoding/json"

	"github.com/99designs/gqlgen/client"
	"github.com/lab259/errors/v2/gqlerrors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vektah/gqlparser/gqlerror"
)

var _ = Describe("GraphQL Extensions", func() {

	Describe("HaveValidation", func() {
		It("should match", func() {
			gqlerr := gqlerror.Error{
				Extensions: map[string]interface{}{
					"errors": map[string]interface{}{
						"name": []interface{}{
							"required",
						},
					},
				},
			}

			a := gqlerrors.HaveValidation("name", "required")
			ok, err := a.Match(gqlerr)
			Expect(ok).To(BeTrue())
			Expect(err).ToNot(HaveOccurred())
		})

		It("should match multiple rules", func() {
			gqlerr := gqlerror.Error{
				Extensions: map[string]interface{}{
					"errors": map[string]interface{}{
						"email": []interface{}{
							"required",
							"email",
							"min=6",
						},
					},
				},
			}

			a := gqlerrors.HaveValidation("email", "required", "email")
			ok, err := a.Match(gqlerr)
			Expect(ok).To(BeTrue())
			Expect(err).ToNot(HaveOccurred())
		})

		It("should match (pointer)", func() {
			gqlerr := &gqlerror.Error{
				Extensions: map[string]interface{}{
					"errors": map[string]interface{}{
						"name": []interface{}{
							"required",
						},
					},
				},
			}

			a := gqlerrors.HaveValidation("name", "required")
			ok, err := a.Match(gqlerr)
			Expect(ok).To(BeTrue())
			Expect(err).ToNot(HaveOccurred())
		})

		It("should match (json.RawMessage)", func() {
			gqlerr := json.RawMessage(`[{"extensions": {"errors": {"name": ["required"]}}}]`)
			a := gqlerrors.HaveValidation("name", "required")
			ok, err := a.Match(gqlerr)
			Expect(ok).To(BeTrue())
			Expect(err).ToNot(HaveOccurred())
		})

		It("should match (client.RawJsonError)", func() {
			gqlerr := client.RawJsonError{json.RawMessage(`[{"extensions": {"errors": {"name": ["required"]}}}]`)}
			a := gqlerrors.HaveValidation("name", "required")
			ok, err := a.Match(gqlerr)
			Expect(ok).To(BeTrue())
			Expect(err).ToNot(HaveOccurred())
		})

		It("should not match", func() {
			gqlerr := gqlerror.Error{
				Extensions: map[string]interface{}{
					"errors": map[string]interface{}{
						"name": []interface{}{
							"required",
						},
					},
				},
			}

			a := gqlerrors.HaveValidation("name", "min")
			ok, err := a.Match(gqlerr)
			Expect(ok).To(BeFalse())
			Expect(err).ToNot(HaveOccurred())
		})

		It("should not match (field)", func() {
			gqlerr := gqlerror.Error{
				Extensions: map[string]interface{}{
					"errors": map[string]interface{}{
						"name": []interface{}{
							"required",
						},
					},
				},
			}

			a := gqlerrors.HaveValidation("age", "min")
			ok, err := a.Match(gqlerr)
			Expect(ok).To(BeFalse())
			Expect(err).ToNot(HaveOccurred())
		})

		It("should not match (not all rules)", func() {
			gqlerr := gqlerror.Error{
				Extensions: map[string]interface{}{
					"errors": map[string]interface{}{
						"name": []interface{}{
							"required",
						},
					},
				},
			}

			a := gqlerrors.HaveValidation("name", "required", "min")
			ok, err := a.Match(gqlerr)
			Expect(ok).To(BeFalse())
			Expect(err).ToNot(HaveOccurred())
		})

		It("should fail without extensions", func() {
			gqlerr := gqlerror.Error{
				Extensions: map[string]interface{}{},
			}

			a := gqlerrors.HaveValidation("name", "required")
			ok, err := a.Match(gqlerr)
			Expect(ok).To(BeFalse())
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Validation extension not found in"))
		})

		It("should fail with wrong actual type", func() {
			a := gqlerrors.HaveValidation("name", "required")
			ok, err := a.Match(26)
			Expect(ok).To(BeFalse())
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("HaveValidation matcher does not know how to handle int"))
		})
	})
})
