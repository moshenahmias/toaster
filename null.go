package toaster

// NullTester is a no-op implementation of the Tester interface.
type NullTester struct{}

func (t *NullTester) Case(params ...any) Tester {
	return t
}

func (t *NullTester) Skip(params ...any) Tester {
	return t
}

func (t *NullTester) Bind(args ...any) Tester {
	return t
}

func (*NullTester) Run(f any) {}

func (*NullTester) Go(f any) {}

// SkipAll returns a Tester that does nothing, effectively skipping all test cases.
// This can be used when you want to skip all tests in a suite.
func SkipAll(reason string) *NullTester {
	return new(NullTester)
}
