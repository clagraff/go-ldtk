package maybe

// Value is a generic container to holding a possibly-nil
// value.
type Value[T any] interface {
	// Value returns either the non-nil value T, or the
	// zero-value of T if it is nil
	Value() T

	// IsNull returns true if the stored value T is nil.
	IsNull() bool

	// Get returns the value T and true when not nil,
	// or the zero-value of T and false when the underlying value
	// is nil.
	Get() (T, bool)
}

func From[T any](v *T) Value[T] {
	return nullable[T]{
		v: v,
	}
}

type nullable[T any] struct {
	v *T
}

func (n nullable[T]) Value() T {
	if n.v == nil {
		var zeroValue T
		return zeroValue
	}

	return *n.v
}

func (n nullable[T]) IsNull() bool {
	return n.v == nil
}

func (n nullable[T]) Get() (T, bool) {
	var v T
	if n.v != nil {
		v = *n.v
	}

	return v, n.v == nil
}

var _ Value[struct{}] = nullable[struct{}]{}
