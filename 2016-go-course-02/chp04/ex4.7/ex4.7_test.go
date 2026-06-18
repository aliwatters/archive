package main_test

import (
	"testing"

	. "github.com/aliwatters/go-course-02/chp04/ex4.7"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// Needed to fire off the tests
func TestSo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Reverse Unicode byte slices in place")
}

var _ = Describe("Empty", func() {
	It("it works with an empty byte slice", func() {
		in := []byte("")
		out := Reverse(in)
		Expect(out).To(Equal(in))
	})

	It("it reverses a simple string", func() {
		in := []byte("Hello World!")
		out := Reverse(in)
		Expect(out).To(Equal([]byte("!dlroW olleH")))
	})

	It("it reverses a unicode string", func() {
		in := []byte("日本語")
		out := Reverse(in)
		Expect(out).To(Equal([]byte("語本日")))
	})

	It("it reverses a mixed ascii unicode string", func() {
		in := []byte("Hello World! 日本語")
		out := Reverse(in)
		Expect(out).To(Equal([]byte("語本日 !dlroW olleH")))
	})

})
