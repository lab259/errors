package testing

import (
	"github.com/lab259/errors"
)

// errComparatorFnc is the method signature that will be called to "visit" each error and its reason with the
// `errorWithReasonIterator`.
type errComparatorFnc func(err error) bool

// errorWithReasonIterator calls `fnc` for the error. If `err` is a `ErrorWithReason`, it calls `fnc` passing the reason.
// Then, repeats until `fnc` returns `false`, or `err` is not `ErrorWithReason`, or the `Reason` is `nil`.
func errorWithReasonIterator(err error, fnc errComparatorFnc) bool {
	if fnc(err) {
		return true
	}
	for e, ok := err.(errors.ErrorWithReason); ok; {
		if fnc(e) {
			return true
		}
		if e.Reason() == nil {
			return false
		}
		e, ok = e.Reason().(errors.ErrorWithReason)
	}
	return false
}

func reasonIterator(err error, fnc errComparatorFnc) bool {
	if fnc(err) {
		return true
	}
	for e, ok := err.(errors.ErrorWithReason); ok; {
		return reasonIterator(e.Reason(), fnc)
	}
	return false
}
