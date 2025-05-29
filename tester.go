package toaster

import (
	"fmt"
	"reflect"
)

// Tester is an interface for creating and running parameterized tests.
type Tester interface {
	// Case adds a test case with the provided parameters to the Tester.
	Case(params ...any) Tester

	// Run executes the provided function with each set of parameters.
	Run(f any)
}

type tester struct {
	cases [][]any
}

// Case creates a new Tester instance and adds the first test case with the provided parameters.
func Case(params ...any) Tester {
	return new(tester).Case(params...)
}

func (t *tester) Case(params ...any) Tester {
	if len(params) == 0 {
		return t
	}

	t.cases = append(t.cases, params)
	return t
}

func (t *tester) Run(f any) {
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
		in[i] = reflect.ValueOf(param)
	}

	_ = fn.Call(in)
}
