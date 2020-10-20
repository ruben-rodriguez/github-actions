package hello

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

)

var _ = Describe("Greeter Test", func() {
	var (
		greeter *Greeter
	)

	BeforeEach(func() {
		greeter = &Greeter{
			Template: "Hello, %s",
		}
	})

	// Test will fail
	It("Should format my name", func() {
		out, err := greeter.Greet(&GreeterInput{
			Name: "Rubenn",
		})
		Expect(err).To(BeNil())
		Expect(out).To(Equal(&GreeterOutput{
			Message: "Hello, Ruben",
		}))
	})
})