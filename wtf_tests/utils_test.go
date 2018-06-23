package wtf_tests

import (
	"testing"

	"github.com/go-test/deep"
	. "github.com/senorprogrammer/wtf/wtf"
)

/* -------------------- Exclude() -------------------- */

func TestExcludeWhenTrue(t *testing.T) {
	if Exclude([]string{"cat", "dog", "rat"}, "bat") != true {
		t.Fatalf("Expected true but got false")
	}
}

func TestExcludeWhenFalse(t *testing.T) {
	if Exclude([]string{"cat", "dog", "rat"}, "dog") != false {
		t.Fatalf("Expected false but got true")
	}
}

/* -------------------- NameFromEmail() -------------------- */

func TestNameFromEmailWhenEmpty(t *testing.T) {
	expected := ""
	actual := NameFromEmail("")

	if expected != actual {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}

func TestNameFromEmailWithEmail(t *testing.T) {
	expected := "Chris Cummer"
	actual := NameFromEmail("chris.cummer@me.com")

	if expected != actual {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}

/* -------------------- NamesFromEmails() -------------------- */

func TestNamesFromEmailsWhenEmpty(t *testing.T) {
	expected := []string{}
	actual := NamesFromEmails([]string{})

	if diff := deep.Equal(expected, actual); diff != nil {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}

func TestNamesFromEmailsWithEmails(t *testing.T) {
	expected := []string{"Chris Cummer", "Chriscummer"}
	actual := NamesFromEmails([]string{"chris.cummer@me.com", "chriscummer@me.com"})

	if diff := deep.Equal(expected, actual); diff != nil {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}

/* -------------------- ToInts() -------------------- */

func TestToInts(t *testing.T) {
	expected := []int{1, 2, 3}

	source := make([]interface{}, len(expected))
	for idx, val := range expected {
		source[idx] = val
	}

	actual := ToInts(source)

	if diff := deep.Equal(expected, actual); diff != nil {
		t.Fatalf("Expected %v but got %v", expected, actual)
	}
}

/* -------------------- ToStrs() -------------------- */

func TestToStrs(t *testing.T) {
	expected := []string{"cat", "dog", "rat"}

	source := make([]interface{}, len(expected))
	for idx, val := range expected {
		source[idx] = val
	}

	actual := ToStrs(source)

	if diff := deep.Equal(expected, actual); diff != nil {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}

/* -------------------- PrettyDate() -------------------- */

func TestPrettyDate(t *testing.T) {
	expected := "Oct 21, 1999"
	actual := PrettyDate("1999-10-21")

	if expected != actual {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}
