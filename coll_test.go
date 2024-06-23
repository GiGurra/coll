package coll

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestToSet(t *testing.T) {
	items1 := []string{"a", "b", "c"}
	items2 := []string{"a", "b", "c", "d"}

	if diff := cmp.Diff(ToSet(items1), map[string]bool{
		"a": true, "b": true, "c": true,
	}); diff != "" {
		t.Errorf("ToSet() mismatch (-want +got):\n%s", diff)
	}

	if diff := cmp.Diff(ToSet(items2), map[string]bool{
		"a": true, "b": true, "c": true, "d": true,
	}); diff != "" {
		t.Errorf("ToSet() mismatch (-want +got):\n%s", diff)
	}
}

type FancyStruct struct {
	A string
	B int
}

func TestToSetF(t *testing.T) {
	items := []FancyStruct{
		{"a", 1},
		{"b", 2},
		{"c", 3},
	}

	if diff := cmp.Diff(ToSetF(items, func(item FancyStruct) string {
		return item.A
	}), map[string]bool{
		"a": true, "b": true, "c": true,
	}); diff != "" {
		t.Errorf("ToSetF() mismatch (-want +got):\n%s", diff)
	}
}

func TestReversed(t *testing.T) {
	items := []string{"a", "b", "c"}

	if diff := cmp.Diff(Reversed(items), []string{"c", "b", "a"}); diff != "" {
		t.Errorf("Reversed() mismatch (-want +got):\n%s", diff)
	}

	// check that the original slice is not modified
	if diff := cmp.Diff(items, []string{"a", "b", "c"}); diff != "" {
		t.Errorf("Reversed() mismatch (-want +got):\n%s", diff)
	}
}
