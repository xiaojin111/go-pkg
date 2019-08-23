package with

type WithFunc func() error

// With wrap a call for a function with error returns
func With(fn WithFunc) error {
	return fn()
}
