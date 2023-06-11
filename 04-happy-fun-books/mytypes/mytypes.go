package mytypes

import "strings"

type MyInt int

func (i MyInt) Twice() MyInt {
	return i * 2
}

type MyString string

func (s MyString) Len() int {
	return len(s)
}

// inherits the strings packages methods
// if you don't need the underlying methods
// dont wrap the Contents type in a struct
// ie: type Contents strings.Builder
type MyBuilder struct {
	Contents strings.Builder
}

func (s MyBuilder) Hello() string {
	return "Hello, Gophers!"
}

type StringUpperCaser struct {
	Contents strings.Builder
}

func (su StringUpperCaser) ToUpper() string {
	return strings.ToUpper(su.Contents.String())
}
