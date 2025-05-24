package main



var e = new(ErrorHandler)

func SetErrorHandler(fn func (error)) {
	e.HandleError = fn
}

type ErrorHandler struct {
	HandleError func (error)
}


func Unwrap[T any](val T, err error) T {
	if err != nil {
		e.HandleError(err)
	}

	return val
}

func UnwrapOr[T any](val T, err error) func(T) T {
	if err != nil {
		return func(d T) T {
			return d
		}
	} else {
		return func(_ T) T {
			return val
		}
	}
}

func UnwrapOrElse[T any](val T, err error) func(func() T) T {
	if err != nil {
		return func(fn func() T) T {
			return fn()
		}
	} else {
		return func(_ func() T) T {
			return val
		}
	}

}

func Expect(err error) {
	if err != nil {
		e.HandleError(err)
	}
}

