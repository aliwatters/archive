package main_test

import (
	"testing"

	. "github.com/aliwatters/go-course-02/chp04/ex4.6"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// Needed to fire off the tests
func TestSo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Convert Adjacent Unicode Whitespaces to Single ASCII")
}

var _ = Describe("Empty", func() {
	It("it works with an empty byte slice", func() {
		in := []byte("")
		out := UnicodeSpaceToASCII(in)
		Expect(out).To(Equal(in))
	})

	It("it converts two ASCII spaces to one", func() {
		in := []byte("  ")
		out := UnicodeSpaceToASCII(in)
		Expect(out).To(Equal([]byte(" ")))
	})

	It("it doesn't break simple ASCII with no space", func() {
		in := []byte("hello")
		out := UnicodeSpaceToASCII(in)
		Expect(out).To(Equal(in))
	})

	It("it doesn't break unicode with no space", func() {
		in := []byte("日本語")
		out := UnicodeSpaceToASCII(in)
		Expect(out).To(Equal(in))
	})

	It("it works for unicode strings with unicode space", func() {
		in := []byte("日\u2005本\u2005\u2005語")
		out := UnicodeSpaceToASCII(in)
		Expect(out).To(Equal([]byte("日 本 語")))
	})

	It("it replaces spaces in mixed unicode ASCII", func() {
		in := []byte("[ ] [ ] [ ] [ ]")
		out := UnicodeSpaceToASCII(in)
		Expect(out).To(Equal([]byte("[ ] [ ] [ ] [ ]")))
	})

	It("it replaces all unicode spaces that I can find", func() {
		in := []byte("\u0020\u00A0\u1680\u2000\u2001\u2002\u2003\u2004\u2005\u2006\u2007\u2008\u2009\u200A\u202F\u205F\u3000")
		out := UnicodeSpaceToASCII(in)
		Expect(out).To(Equal([]byte(" ")))
	})

	// note: https://www.cs.tut.fi/~jkorpela/chars/spaces.html -- used as source
	// several characters didn't get detected - \u180E \u180E \u200B \uFEFF
	// not sure if this is a bug

})
