package greatings

import (
	"regexp"
	"testing"
)

// To test this you use "go test" command that search all files with
// end _test.go and execute these unit tests, "go test -v" show the verbose
// execution

// TestHelloName calls greetings. Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) { //t is a pointer to type testing.T
	name := "Michael"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("`+name+`") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
