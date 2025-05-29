package toaster_test

import (
	"testing"

	"github.com/moshenahmias/toaster"
)

func TestCaseAndRun_SingleCase(t *testing.T) {
	var called bool
	var gotA int
	var gotB string

	toaster.Case(42, "hello").Run(func(a int, b string) {
		called = true
		gotA = a
		gotB = b
	})

	if !called {
		t.Error("expected function to be called")
	}
	if gotA != 42 || gotB != "hello" {
		t.Errorf("unexpected arguments: gotA=%v, gotB=%v", gotA, gotB)
	}
}

func TestCaseAndRun_MultipleCases(t *testing.T) {
	var results []string

	toaster.Case(1, "a").Case(2, "b").Case(3, "c").Run(func(a int, b string) {
		results = append(results, b)
	})

	if len(results) != 3 {
		t.Errorf("expected 3 calls, got %d", len(results))
	}
	if results[0] != "a" || results[1] != "b" || results[2] != "c" {
		t.Errorf("unexpected results: %v", results)
	}
}

func TestRun_PanicOnNonFunc(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic for non-function argument")
		}
	}()
	toaster.Case(1).Run(123)
}

func TestRun_PanicOnWrongParamCount(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic for wrong parameter count")
		}
	}()
	toaster.Case(1, 2).Run(func(a int) {})
}

func TestCase_Chaining(t *testing.T) {
	var count int
	c := toaster.Case(1).Case(2).Case(3)
	c.Run(func(a int) {
		count += a
	})
	if count != 6 {
		t.Errorf("expected sum 6, got %d", count)
	}
}
func TestCase_Empty(t *testing.T) {
	var called bool

	toaster.Case().Run(func() {
		called = true
	})

	if called {
		t.Error("expected function not to be called for empty case")
	}
}

func TestSkip(t *testing.T) {
	var s []int

	toaster.Skip(1).Case(2).Skip(3).Run(func(x int) {
		s = append(s, x)
	})

	if len(s) != 1 || s[0] != 2 {
		t.Errorf("expected slice to contain only 2, got %v", s)
	}
}
