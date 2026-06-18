package main_test

import (
	"strings"
	"testing"

	. "github.com/aliwatters/go-course-02/chp04/ex4.9"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// Needed to fire off the tests
func TestSo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Word Frequency")
}

var _ = Describe("From Strings", func() {
	It("it works with an empty string", func() {
		in := strings.NewReader("")
		res := WordFreq(in)
		Expect(res).To(Equal(map[string]int{}))
	})

	It("it counts words in an simple ascii string", func() {
		in := strings.NewReader("Hello World!")
		res := WordFreq(in)
		eg := map[string]int{"hello": 1, "world": 1}
		Expect(res).To(Equal(eg))
	})

	It("it counts words in a unicode string", func() {
		in := strings.NewReader("Hello \"World\"! 語本日")
		res := WordFreq(in)
		eg := map[string]int{"hello": 1, "world": 1, "語本日": 1}
		Expect(res).To(Equal(eg))
	})

	It("it counts words in a multi-line string", func() {
		sample := `hello hello
		hello world!
		語本日 語本日`
		in := strings.NewReader(sample)
		res := WordFreq(in)
		eg := map[string]int{"hello": 3, "world": 1, "語本日": 2}
		Expect(res).To(Equal(eg))
	})

})
