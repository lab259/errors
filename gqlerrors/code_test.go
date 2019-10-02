package gqlerrors_test

import (
	"encoding/json"

	"github.com/99designs/gqlgen/client"

	"github.com/lab259/errors/v2"
	"github.com/lab259/errors/v2/gqlerrors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vektah/gqlparser/gqlerror"
)

var _ = Describe("GraphQL Extensions", func() {

	Describe("HaveCode", func() {
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

})
