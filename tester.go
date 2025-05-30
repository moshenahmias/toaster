package toaster

import (
	"fmt"
	"reflect"
)

// Tester is an interface for creating and running parameterized tests.
type Tester interface {
	// Case adds a test case with the provided arguments to the Tester.
	Case(args ...any) Tester

	// Skip is a no-op method that can be used to skip the test case.
	Skip(args ...any) Tester

	// Bind binds the provided arguments to the Tester, which will be used in each test case.
	Bind(args ...any) Tester

	// Run executes the provided function with each set of arguments.
	Run(f any)
}

// Evaluator is an interface for evaluating a function with no parameters.
type Evaluator interface {
	// Evaluate runs the function and returns the result.
	Evaluate() any
}

// EvaluatorFunc is a function type that implements the Evaluator interface.
type EvaluatorFunc func() any

func (f EvaluatorFunc) Evaluate() any {
	return f()
}

type tester struct {
	cases [][]any
	bind  []any
}

// Case creates a new Tester instance and adds the first test case with the provided arguments.
func Case(args ...any) Tester {
	return new(tester).Case(args...)
}

// Skip creates a new Tester instance that skips the test case with the provided arguments.
func Skip(args ...any) Tester {
	return new(tester)
}

// Bind creates a new Tester instance and binds the provided arguments to it.
func Bind(args ...any) Tester {
	return new(tester).Bind(args...)
}

func (t *tester) Skip(args ...any) Tester {
	return t
}

func (t *tester) Case(args ...any) Tester {
	if len(args) > 0 {
		t.cases = append(t.cases, args)
	}

	return t
}

func (t *tester) Bind(args ...any) Tester {
	t.bind = append(t.bind, args...)

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

	bind := make([]reflect.Value, len(t.bind))

	for i, param := range t.bind {
		bind[i] = reflect.ValueOf(param)
	}

	for i, testCase := range t.cases {
		if len(testCase)+len(bind) != fn.Type().NumIn() {
			panic(fmt.Sprintf("case %d: expected %d arguments, got %d", i, fn.Type().NumIn(), len(testCase)+len(bind)))
		}

		t.runCase(fn, bind, testCase)
	}
}

func (t *tester) runCase(fn reflect.Value, bind []reflect.Value, testCase []any) {
	in := make([]reflect.Value, len(testCase)+len(bind))

	copy(in, bind)

	for i, param := range testCase {

		if ev, ok := param.(Evaluator); ok {
			param = ev.Evaluate()
		}

		in[i+len(bind)] = reflect.ValueOf(param)
	}

	_ = fn.Call(in)
}
