package greetings;

import (
	"testing"
	"regexp"
)

func TestHelloName(t *testing.T) {
	name := "Orca";
	test := regexp.MustCompile(`.*[` + name + `]{1}.*`);
	msg, err := Hello("Orca");

	if (!test.MatchString(msg) || err != nil) {
		t.Fatalf("Test fail");
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("");
	if (msg != "" || err == nil) {
		t.Fatalf(`Hello("") %q, %v, want "", error`, msg, err);
	}
}
