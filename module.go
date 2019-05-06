package errors

import "fmt"

// ModuleError describes an error that belongs to a module.
type ModuleError interface {
	Module() string
}

type moduleError struct {
	reason error
	module string
}

// Unwrap returns the next error in the error chain.
// If there is no next error, Unwrap returns nil.
func (err *moduleError) Unwrap() error {
	return err.reason
}

// Module returns the name of the module related to this error.
func (err *moduleError) Module() string {
	return err.module
}

// Error returns the original reason of the error.
func (err *moduleError) Error() string {
	if err.reason == nil {
		return fmt.Sprintf("%s: unknown error", err.module)
	}
	return fmt.Sprintf("%s: %s", err.module, err.reason.Error())
}

// AppendData adds the module information to the system.
func (err *moduleError) AppendData(response ErrorResponse) {
	response.SetParam("module", err.module)
}

// WrapModule wraps an original error with the module information, returning a
// new instance with the reason set.
func WrapModule(reason error, module string) error {
	return &moduleError{
		reason: reason,
		module: module,
	}
}
