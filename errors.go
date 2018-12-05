package errors

type Option func(err error) error

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

func Wrap(reason error, options ...Option) error {
	err := reason
	for _, opt := range options {
		if opt != nil {
			err = opt(err)
		}
	}
	return err
}
