package testing_test

import (
	"reflect"
	"strings"

	"github.com/lab259/errors/v2"
	"github.com/lab259/errors/v2/testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	validator "gopkg.in/go-playground/validator.v9"
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

			m := testing.ErrorWithValidation("Age", "min")

			result, err := m.Match(errors.Wrap(err, errors.Validation()))
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

			m := testing.ErrorWithValidation("Name", "required")

			result, errMatch := m.Match(errors.Wrap(err, errors.Validation()))
			Expect(result).To(BeTrue())
			Expect(errMatch).ToNot(HaveOccurred())

			m = testing.ErrorWithValidation("Age", "min")

			result, errMatch = m.Match(errors.Wrap(err, errors.Validation()))
			Expect(result).To(BeTrue())
			Expect(errMatch).ToNot(HaveOccurred())
		})

		It("should initialize the matcher two field and many rules", func() {
			type Person struct {
				Name    string `json:"name"   validate:"required"`
				Age     int    `json:"age"    validate:"min=0,max=21"`
				Status  bool   `json:"status" validate:"-"`
				Tagline string `validate:"required,lt=10"`
			}

			person := Person{
				Name:    "Chico Bento",
				Age:     -1,
				Status:  true,
				Tagline: "This tagline is way too long.",
			}

			err := validate.Struct(person)
			Expect(err).To(HaveOccurred())

			m := testing.ErrorWithValidation("Age", "min")

			result, errMatch := m.Match(errors.Wrap(err, errors.Validation()))
			Expect(result).To(BeTrue())
			Expect(errMatch).ToNot(HaveOccurred())

			person.Age = 22

			err = validate.Struct(person)
			Expect(err).To(HaveOccurred())

			m = testing.ErrorWithValidation("Age", "max")

			result, errMatch = m.Match(errors.Wrap(err, errors.Validation()))
			Expect(result).To(BeTrue())
			Expect(errMatch).ToNot(HaveOccurred())

			m = testing.ErrorWithValidation("Tagline", "lt")

			result, errMatch = m.Match(errors.Wrap(err, errors.Validation()))
			Expect(result).To(BeTrue())
			Expect(errMatch).ToNot(HaveOccurred())
		})

		It("should initialize the matcher email invalid", func() {
			type Person struct {
				Name  string `json:"name"   validate:"required"`
				Age   int    `json:"age"    validate:"min=0,max=21"`
				Email string `json:"status" validate:"required,email"`
			}

			person := Person{
				Name:  "Chico Bento",
				Age:   20,
				Email: " ",
			}

			err := validate.Struct(person)
			Expect(err).To(HaveOccurred())

			m := testing.ErrorWithValidation("Email", "email")

			result, errMatch := m.Match(errors.Wrap(err, errors.Validation()))
			Expect(result).To(BeTrue())
			Expect(errMatch).ToNot(HaveOccurred())
		})

		It("should initialize the matcher email valid", func() {
			type Person struct {
				Name  string `json:"name"   validate:"required"`
				Age   int    `json:"age"    validate:"min=0,max=21"`
				Email string `json:"status" validate:"required,email"`
			}

			person := Person{
				Name:  "Chico Bento",
				Age:   20,
				Email: "not is email.",
			}

			err := validate.Struct(person)
			Expect(err).To(HaveOccurred())

			m := testing.ErrorWithValidation("Email", "email")

			result, errMatch := m.Match(errors.Wrap(err, errors.Validation()))
			Expect(result).To(BeTrue())
			Expect(errMatch).ToNot(HaveOccurred())
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

			result, err := m.Match(errors.Wrap(err, errors.Validation()))
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

			m := testing.ErrorWithValidation("Age", "min")

			result, err := m.Match(wrapError)
			Expect(result).To(BeTrue())
			Expect(err).ToNot(HaveOccurred())
		})

		It("should initialize the matcher validation wrap module", func() {
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

			errModule := errors.Module("test")

			wrapError := errors.Wrap(err, errModule, errors.Validation())

			m := testing.ErrorWithValidation("Age", "min")

			result, err := m.Match(wrapError)
			Expect(result).To(BeTrue())
			Expect(err).ToNot(HaveOccurred())
		})

		It("should fail initialize the matcher validation wrap module", func() {
			type Person struct {
				Name   string `json:"name"   validate:"required"`
				Age    int    `json:"age"    validate:"max=100"`
				Status bool   `json:"status" validate:"-"`
			}

			person := Person{
				Name:   "Chico Bento",
				Age:    101,
				Status: true,
			}

			err := validate.Struct(person)
			Expect(err).To(HaveOccurred())

			errModule := errors.Module("test")
			wrapError := errModule(errors.Wrap(err, errors.Validation()))

			// Case invalid
			m := testing.ErrorWithValidation("Name", "required")

			result, err := m.Match(wrapError)
			Expect(result).To(BeFalse())
			Expect(err).To(HaveOccurred())

			// Case valid
			m = testing.ErrorWithValidation("Age", "max")

			result, err = m.Match(wrapError)
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
