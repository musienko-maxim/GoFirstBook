package main

import (
	"fmt"
	"github.com/musienko-maxim/GoFirstBook/chapter14/prose"
	"testing"
)

func errorString(list []string, got string, want string) string {
	return fmt.Sprintf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", list, got, want)
}

func TestOneElement(t *testing.T) {
	list := []string{"apple"}
	want := "apple"
	got := prose.JoinWithCommas(list)
	if got != want {
		t.Errorf(errorString(list, got, want))
	}
}

func TestTwoElements(t *testing.T) {
	list := []string{"apple", "orange"}
	want := "apple and orange"
	got := prose.JoinWithCommas(list)
	if got != want {
		t.Errorf(errorString(list, got, want))
	}
}

func TestThreeElements(t *testing.T) {
	list := []string{"apple", "orange", "pear"}
	want := "apple, orange, and pear"
	got := prose.JoinWithCommas(list)
	if got != want {
		t.Errorf(errorString(list, got, want))
	}
}
