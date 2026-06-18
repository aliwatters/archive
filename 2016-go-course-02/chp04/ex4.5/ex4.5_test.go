package main_test

import (
	"testing"

	. "github.com/aliwatters/go-course-02/chp04/ex4.5"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// Needed to fire off the tests
func TestSo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Elimnate Adjacent Dups Tests")
}

var _ = Describe("eliminates adjacent dups in slice of strings", func() {
	It("it works with an empty slice", func() {
		s := []string{}
		resp := EliminateDups(s)
		Expect(resp).To(Equal([]string{}))
	})

	It("it elminates simple dups", func() {
		s := []string{"a", "a"}
		resp := EliminateDups(s)
		Expect(resp).To(Equal([]string{"a"}))
	})

	It("it elminates multiple dups", func() {
		s := []string{"a", "a", "b", "b", "b", "c", "c", "b", "d", "d"}
		resp := EliminateDups(s)
		Expect(resp).To(Equal([]string{"a", "b", "c", "b", "d"}))
	})

	It("it eliminates longer dups", func() {
		s := []string{"ant", "bat", "bat", "cat", "cow", "bat", "dog"}
		resp := EliminateDups(s)
		Expect(resp).To(Equal([]string{"ant", "bat", "cat", "cow", "bat", "dog"}))
	})

})
