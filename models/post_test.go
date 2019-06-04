package models

import "testing"

func TestIncrementURI(t *testing.T) {
	a := IncrementURI("aaa")
	if a != "aab" {
		t.Error("Expected aab, got ", a)
	}
	b := IncrementURI("aaz")
	if b != "aaA" {
		t.Error("Expected aaA, got ", b)
	}
	c := IncrementURI("aaZ")
	if c != "aba" {
		t.Error("Expected aba, got ", c)
	}
	d := IncrementURI("azZ")
	if d != "aAa" {
		t.Error("Expected aAa, got ", d)
	}
	e := IncrementURI("aZZ")
	if e != "baa" {
		t.Error("Expected baa, got ", e)
	}
	f := IncrementURI("aZZ")
	if f != "baa" {
		t.Error("Expected baa, got ", f)
	}
	g := IncrementURI("zZZ")
	if g != "Aaa" {
		t.Error("Expected Aaa, got ", g)
	}
}
