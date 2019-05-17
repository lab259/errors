package gqlerrors_test

import (
	"github.com/gofrs/uuid"
	"github.com/lab259/errors/gqlerrors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GraphQL Testing Utils", func() {
	Describe("Extensions", func() {
		Describe("Decode", func() {

			It("should decode map structure to struct", func() {
				id := uuid.Must(uuid.NewV4())
				input := map[string]interface{}{
					"id":   id,
					"name": "Chico Bento",
				}

				var person Person
				err := gqlerrors.Decode(input, &person)
				Expect(err).ToNot(HaveOccurred())

				Expect(person.ID).To(Equal(id))
				Expect(person.Name).To(Equal("Chico Bento"))
				Expect(person.Age).To(Equal(0))
			})

			It("should decode map structure to struct id text", func() {
				id := uuid.Must(uuid.NewV4())
				input := map[string]interface{}{
					"id":   id.String(),
					"name": "Chico Bento",
				}

				var person Person
				err := gqlerrors.Decode(input, &person)
				Expect(err).ToNot(HaveOccurred())

				Expect(person.ID).To(Equal(id))
				Expect(person.Name).To(Equal("Chico Bento"))
				Expect(person.Age).To(Equal(0))
			})

			It("should fail decoding map structure result must be a pointer", func() {
				input := map[string]interface{}{
					"name": "Francisco Bento",
				}

				var person Person
				err := gqlerrors.Decode(input, person)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("result must be a pointer"))

				Expect(person.ID).Should(Equal(uuid.Nil))
				Expect(person.Name).Should(Equal(""))
				Expect(person.Age).Should(Equal(0))
			})

			It("should fail decoding map structure result must be a pointer", func() {
				input := map[string]interface{}{
					"id":   12345,
					"name": "Francisco Bento",
				}

				var person Person
				err := gqlerrors.Decode(input, &person)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).Should(ContainSubstring("error(s) decoding"))

				Expect(person.ID).Should(Equal(uuid.Nil))
				Expect(person.Name).Should(Equal("Francisco Bento"))
				Expect(person.Age).Should(Equal(0))
			})
		})
	})
})

type Person struct {
	ID   uuid.UUID
	Name string
	Age  int
}
