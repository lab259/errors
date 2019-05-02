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

func Combine(options ...interface{}) Option {
	return func(err error) error {
		for _, opt := range options {
			switch act := opt.(type) {
			case Option:
				err = act(err)
			case string:
				err = WrapMessage(err, act)
			case int:
				err = WrapHttp(err, act)
			default:
				continue
			}
		}
		return err
	}
}

func Wrap(reason error, options ...interface{}) error {
	return Combine(options...)(reason)
}

func Is(err, target error) bool {
	for {
		if err == target {
			return true
		}
		wrapper, ok := err.(Wrapper)
		if !ok {
			return false
		}
		err = wrapper.Unwrap()
		if err == nil {
			return false
		}
	}
}

func Reason(err error) error {
	for {
		wrapper, ok := err.(Wrapper)
		if !ok {
			return err
		}
		err = wrapper.Unwrap()
	}
}
