package main

import (
	"fmt"
	"github.com/musienko-maxim/GoFirstBook/chapter14/prose"
	"testing"
)

func TestTwoElements(t *testing.T) {
	list := []string{"apple", "orange"}
	fmt.Println(">>>>>>>>>>>>>>>>: " + prose.JoinWithCommas(list))
	if prose.JoinWithCommas(list) != "apple and orange" {
		t.Error("didn't match expected value")
	}
}

func TestThreeElements(t *testing.T) {
	list := []string{"apple", "orange", "pear"}
	fmt.Println("---------------: " + prose.JoinWithCommas(list))
	if prose.JoinWithCommas(list) != "apple, orange, and pear" {
		t.Error("didn't match expected value")
	}
}
