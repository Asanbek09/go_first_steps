package strings

import (
	"sync"
	"unicode/utf8"
	"unsafe"
)

type Builder struct {
	addr *Builder // of receiver, to detect copies by value
	buf []byte
}

package strings

func Compare(a, b string) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return +1
}

package strings

import (
	"io"
	"sync"
)

type Replacer struct {
	once sync.Once
	r replacer
	oldnew []string
}