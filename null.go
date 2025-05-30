package toaster

type nullTester struct{}

func (t *nullTester) Case(params ...any) Tester {
	return t
}

func (t *nullTester) Skip(params ...any) Tester {
	return t
}

func (t *nullTester) Bind(params ...any) Tester {
	return t
}

func (*nullTester) Run(f any) {}

// SkipAll returns a Tester that does nothing, effectively skipping all test cases.
// This can be used when you want to skip all tests in a suite.
func SkipAll(reason string) Tester {
	return new(nullTester)
}
