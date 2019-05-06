package testing

import (
	"github.com/lab259/errors"
)

// errComparatorFnc is the method signature that will be called to "visit" each error and its reason with the
// `errorWithReasonIterator`.
type errComparatorFnc func(err error) bool

// errorWithReasonIterator calls `fnc` for the error. If `err` is a `Wrapper`, it calls `fnc` passing the reason.
// Then, repeats until `fnc` returns `false`, or `err` is not `Wrapper`, or the `Reason` is `nil`.
func errorWithReasonIterator(err error, fnc errComparatorFnc) bool {
	for {
		if fnc(err) {
			return true
		}
		wrapper, ok := err.(errors.Wrapper)
		if !ok {
			return false
		}
		err = wrapper.Unwrap()
		if err == nil {
			return false
		}
	}
}