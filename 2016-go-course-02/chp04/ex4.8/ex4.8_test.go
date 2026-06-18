package main_test

import (
	"fmt"
	"strings"
	"testing"

	. "github.com/aliwatters/go-course-02/chp04/ex4.8"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// Needed to fire off the tests
func TestSo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Count Unicode Character Category Frequency")
}

var _ = Describe("Empty", func() {
	It("it works with an empty string", func() {
		in := strings.NewReader("")
		res := CharCount(in)
		stat := UnicodeStat{map[rune]int{}, map[int]int{}, map[string]int{}, 0}
		Expect(res).To(Equal(stat))
	})

	It("it counts stats in an simple ascii string", func() {
		in := strings.NewReader("Hello 2 World!")
		res := CharCount(in)
		stat := UnicodeStat{map[rune]int{
			'H': 1, 'e': 1, 'l': 3, 'o': 2, ' ': 2,
			'2': 1, 'W': 1, 'r': 1, 'd': 1, '!': 1,
		},
			map[int]int{1: 14},
			map[string]int{"letter": 10, "space": 2, "number": 1, "punct": 1},
			0}
		Expect(res).To(Equal(stat))
	})
	It("it counts stats in an simple ascii string", func() {
		in := strings.NewReader("語本日")
		res := CharCount(in)
		fmt.Println(res)
		stat := UnicodeStat{map[rune]int{'語': 1, '本': 1, '日': 1},
			map[int]int{3: 3},
			map[string]int{"letter": 3},
			0}
		Expect(res).To(Equal(stat))
	})
})
