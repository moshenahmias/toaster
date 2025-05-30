package toaster

import (
	"fmt"
	"reflect"
)

// Tester is an interface for creating and running parameterized tests.
type Tester interface {
	// Case adds a test case with the provided parameters to the Tester.
	Case(params ...any) Tester

	// Skip is a no-op method that can be used to skip the test case.
	Skip(params ...any) Tester

	// Run executes the provided function with each set of parameters.
	// The function must accept zero parameters and may return a boolean value.
	// If the function returns false, the test execution stops.
	Run(f any)
}

// Evaluator is an interface for evaluating a function with no parameters.
type Evaluator interface {
	// Evaluate runs the function with the provided parameters and returns the result.
	Evaluate() any
}

// EvaluatorFunc is a function type that implements the Evaluator interface.
type EvaluatorFunc func() any

func (f EvaluatorFunc) Evaluate() any {
	return f()
}

type tester struct {
	cases [][]any
}

// Case creates a new Tester instance and adds the first test case with the provided parameters.
func Case(params ...any) Tester {
	return new(tester).Case(params...)
}

// Skip creates a new Tester instance that skips the test case with the provided parameters.
func Skip(params ...any) Tester {
	return new(tester)
}

func (t *tester) Skip(params ...any) Tester {
	return t
}

func (t *tester) Case(params ...any) Tester {
	if len(params) > 0 {
		t.cases = append(t.cases, params)
	}

	return t
}

func (t *tester) Run(f any) {
	if f == nil {
		panic("f must not be nil")
	}

	fn := reflect.ValueOf(f)

	if fn.Kind() != reflect.Func {
		panic("f must be a function")
	}

	for i, testCase := range t.cases {
		if len(testCase) != fn.Type().NumIn() {
			panic(fmt.Sprintf("case %d: expected %d parameters, got %d", i, fn.Type().NumIn(), len(testCase)))
		}

		t.runCase(fn, testCase)
	}
}

func (t *tester) runCase(fn reflect.Value, testCase []any) {
	in := make([]reflect.Value, len(testCase))

	for i, param := range testCase {

		if ev, ok := param.(Evaluator); ok {
			param = ev.Evaluate()
		}

		in[i] = reflect.ValueOf(param)
	}

	_ = fn.Call(in)
}
