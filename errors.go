package errors

import "errors"

type Option func(err error) error

// New is an alias for the default `errors.New`.
var New = errors.New

func Http(status int) Option {
	return func(reason error) error {
		return WrapHttp(reason, status)
	}
}

func Module(module string) Option {
	return func(reason error) error {
		return WrapModule(reason, module)
	}
}

func Code(code string) Option {
	return func(reason error) error {
		return WrapCode(reason, code)
	}
}

func Validation() Option {
	return func(err error) error {
		return WrapValidation(err)
	}
}

func Message(message string) Option {
	return func(err error) error {
		return WrapMessage(err, message)
	}
}

func Wrap(reason error, options ...Option) error {
	err := reason
	for _, opt := range options {
		if opt != nil {
			err = opt(err)
		}
	}
	return err
}

func Is(err, target error) bool {
	return false
}
