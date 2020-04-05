package datasection_test

import (
	"testing"

	"github.com/skaji/go-datasection/pkg/datasection"
)

func TestBasic(t *testing.T) {
	sections := datasection.Parse(`
@@  ba z          
1

@@ foo
a
b

c
`)
	if len(sections) != 2 {
		t.Fail()
	}
	if sections["foo"] != "a\nb\n\nc\n" {
		t.Fail()
	}
	if sections["ba z"] != "1\n" {
		t.Fail()
	}
}
