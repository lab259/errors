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
	Describe("HaveModule", func() {
		It("should match", func() {
			gqlerr := gqlerror.Error{
				Extensions: map[string]interface{}{
					"module": "api",
				},
			}

			a := gqlerrors.HaveModule("api")
			ok, err := a.Match(gqlerr)
			Expect(ok).To(BeTrue())
			Expect(err).ToNot(HaveOccurred())
		})

		It("should match using Option", func() {
			gqlerr := gqlerror.Error{
				Extensions: map[string]interface{}{
					"module": "api",
				},
			}

			a := gqlerrors.HaveModule(errors.Module("api"))
			ok, err := a.Match(gqlerr)
			Expect(ok).To(BeTrue())
			Expect(err).ToNot(HaveOccurred())
		})

		It("should match using ErrorWithModule", func() {
			gqlerr := gqlerror.Error{
				Extensions: map[string]interface{}{
					"module": "api",
				},
			}

			a := gqlerrors.HaveModule(errors.Module("api")(nil))
			ok, err := a.Match(gqlerr)
			Expect(ok).To(BeTrue())
			Expect(err).ToNot(HaveOccurred())
		})

		It("should match (pointer)", func() {
			gqlerr := &gqlerror.Error{
				Extensions: map[string]interface{}{
					"module": "api",
				},
			}

			a := gqlerrors.HaveModule("api")
			ok, err := a.Match(gqlerr)
			Expect(ok).To(BeTrue())
			Expect(err).ToNot(HaveOccurred())
		})

		It("should match (json.RawMessage)", func() {
			gqlerr := json.RawMessage(`[{"extensions": {"module": "api"}}]`)
			a := gqlerrors.HaveModule("api")
			ok, err := a.Match(gqlerr)
			Expect(ok).To(BeTrue())
			Expect(err).ToNot(HaveOccurred())
		})

		It("should match (client.RawJsonError)", func() {
			gqlerr := client.RawJsonError{json.RawMessage(`[{"extensions": {"module": "api"}}]`)}
			a := gqlerrors.HaveModule("api")
			ok, err := a.Match(gqlerr)
			Expect(ok).To(BeTrue())
			Expect(err).ToNot(HaveOccurred())
		})

		It("should not match", func() {
			gqlerr := gqlerror.Error{
				Extensions: map[string]interface{}{
					"module": "api",
				},
			}

			a := gqlerrors.HaveModule("not-found")
			ok, err := a.Match(gqlerr)
			Expect(ok).To(BeFalse())
			Expect(err).ToNot(HaveOccurred())
		})

		It("should fail without extensions", func() {
			gqlerr := gqlerror.Error{
				Extensions: map[string]interface{}{},
			}

			a := gqlerrors.HaveModule("api")
			ok, err := a.Match(gqlerr)
			Expect(ok).To(BeFalse())
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Module extension not found in"))
		})

		It("should fail with wrong errors.Option", func() {
			gqlerr := gqlerror.Error{
				Extensions: map[string]interface{}{
					"module": "api",
				},
			}

			a := gqlerrors.HaveModule(errors.Code("api"))
			ok, err := a.Match(gqlerr)
			Expect(ok).To(BeFalse())
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("HaveModule matcher only support an `errors.Module` option"))
		})

		It("should fail with wrong actual type", func() {
			a := gqlerrors.HaveModule("api")
			ok, err := a.Match(26)
			Expect(ok).To(BeFalse())
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("HaveModule matcher does not know how to handle int"))
		})

		It("should fail with wrong expected type", func() {
			gqlerr := gqlerror.Error{
				Extensions: map[string]interface{}{
					"module": "api",
				},
			}

			a := gqlerrors.HaveModule(26)
			ok, err := a.Match(gqlerr)
			Expect(ok).To(BeFalse())
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("HaveModule matcher does not know how to assert int"))
		})
	})

})
