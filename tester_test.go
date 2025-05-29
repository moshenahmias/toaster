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

func TestSkip_Empty(t *testing.T) {
	var called bool

	toaster.Skip().Run(func() {
		called = true
	})

	if called {
		t.Error("expected function not to be called for empty skip")
	}
}

func TestSkip_Chaining(t *testing.T) {
	var results []int

	toaster.Skip(1).Case(2).Skip(3).Case(4).Run(func(x int) {
		results = append(results, x)
	})

	if len(results) != 2 || results[0] != 2 || results[1] != 4 {
		t.Errorf("expected results to be [2, 4], got %v", results)
	}
}

func TestSkip_CaseAfterSkip(t *testing.T) {
	var results []int

	toaster.Skip(1).Case(2).Skip(3).Case(4).Run(func(x int) {
		results = append(results, x)
	})

	if len(results) != 2 || results[0] != 2 || results[1] != 4 {
		t.Errorf("expected results to be [2, 4], got %v", results)
	}
}

func TestSkip_CaseBeforeSkip(t *testing.T) {
	var results []int

	toaster.Case(1).Skip(2).Case(3).Run(func(x int) {
		results = append(results, x)
	})

	if len(results) != 2 || results[0] != 1 || results[1] != 3 {
		t.Errorf("expected results to be [1, 3], got %v", results)
	}
}

func TestSkip_CaseBeforeAndAfterSkip(t *testing.T) {
	var results []int

	toaster.Case(1).Skip(2).Case(3).Run(func(x int) {
		results = append(results, x)
	})

	if len(results) != 2 || results[0] != 1 || results[1] != 3 {
		t.Errorf("expected results to be [1, 3], got %v", results)
	}
}

func TestSkip_EmptyCase(t *testing.T) {
	var called bool

	toaster.Skip().Case().Run(func() {
		called = true
	})

	if called {
		t.Error("expected function not to be called for empty case after skip")
	}
}

func TestSkip_EmptyCaseAfterSkip(t *testing.T) {
	var called bool

	toaster.Skip().Case().Run(func() {
		called = true
	})

	if called {
		t.Error("expected function not to be called for empty case after skip")
	}
}

func TestEvaluator(t *testing.T) {
	var called bool
	var gotA int
	var gotB string

	toaster.Case(42, toaster.EvaluatorFunc(func() any {
		return "hello"
	})).Run(func(a int, b string) {
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

func TestEvaluator2(t *testing.T) {
	var called bool
	var gotA int
	var gotB string

	toaster.Case(42, toaster.EvaluatorFunc(func() any {
		return "hello"
	})).Run(func(a int, b string) {
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
