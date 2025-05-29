# Toaster

**Toaster** is a lightweight Go package for **parameterized testing**, allowing dynamic test cases to be executed with different sets of parameters.

## Features
- Define multiple test cases dynamically.
- Execute functions with varied parameters effortlessly.
- Simple and scalable API for clean testing.

## Installation

```sh
go get github.com/moshenahmias/toaster
```

## Usage

```go
func Add(a, b int) int {
	return a + b
}

func TestAddFunction(t *testing.T) {
	toaster.
		Case(1, 2, 3).
		Case(3, 4, 7).
		Skip(5, 6, 200). // Skipped
		Case(7, 8, toaster.EvaluatorFunc(func() any { return 15 })).
		Run(func(a, b, expected int) {
			if result := Add(a, b); result != expected {
				t.Errorf("expected %d + %d = %d, got %d", a, b, expected, result)
			}
		})
}
```
