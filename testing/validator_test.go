package testing_test

import (
	"github.com/lab259/errors"
	"github.com/lab259/errors/testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/go-playground/validator.v9"
	"reflect"
	"strings"
)

var _ = Describe("ErrorWithValidator Test Suite", func() {

	When("should initialize the match validation", func() {
		It("should initialize the matcher validation", func() {
			type Person struct {
				Name   string `json:"name"   validate:"required"`
				Age    int    `json:"age"    validate:"min=0"`
				Status bool   `json:"status" validate:"-"`
			}

			person := Person{
				Name:   "Chico Bento",
				Age:    -1,
				Status: true,
			}

			err := validate.Struct(person)
			Expect(err).To(HaveOccurred())

			m := testing.ErrorWithValidation("Age")

			result, err := m.Match(err)
			Expect(result).To(BeTrue())
			Expect(err).ToNot(HaveOccurred())
		})

		It("should initialize the matcher two field", func() {
			type Person struct {
				Name   string `json:"name"   validate:"required"`
				Age    int    `json:"age"    validate:"min=0"`
				Status bool   `json:"status" validate:"-"`
			}

			person := Person{
				Age:    -1,
				Status: true,
			}

			err := validate.Struct(person)
			Expect(err).To(HaveOccurred())

			m := testing.ErrorWithValidation("Age", "Name")

			result, err := m.Match(err)
			Expect(result).To(BeTrue())
			Expect(err).ToNot(HaveOccurred())
		})

		It("should initialize the matcher fail field pass not matcher", func() {
			type Person struct {
				Name   string `json:"name"   validate:"required"`
				Age    int    `json:"age"    validate:"min=0"`
				Status bool   `json:"status" validate:"-"`
			}

			person := Person{
				Age:    -1,
				Status: true,
			}

			err := validate.Struct(person)
			Expect(err).To(HaveOccurred())

			m := testing.ErrorWithValidation("Age", "Status")

			result, err := m.Match(err)
			Expect(result).To(BeFalse())
			Expect(err).To(HaveOccurred())
		})
	})

	When("should initialize the matcher validation to reason", func() {
		It("should initialize the matcher validation wrap validation", func() {
			type Person struct {
				Name   string `json:"name"   validate:"required"`
				Age    int    `json:"age"    validate:"min=0"`
				Status bool   `json:"status" validate:"-"`
			}

			person := Person{
				Name:   "Chico Bento",
				Age:    -1,
				Status: true,
			}

			err := validate.Struct(person)
			Expect(err).To(HaveOccurred())

			wrapError := errors.WrapValidation(err)

			m := testing.ErrorWithValidation("Age")

			result, err := m.Match(wrapError)
			Expect(result).To(BeTrue())
			Expect(err).ToNot(HaveOccurred())
		})
	})
})

var validate validator.Validate

func init() {

	validate = *validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}
