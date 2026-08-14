package polykit

type PloyKit interface {
	IsNil(v interface{}) bool
	IsEmpty(v interface{}) bool
	ErrIfNotNil(v interface{}, message string) error
}
