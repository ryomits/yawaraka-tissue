package problem

type ErrUnauthorized struct {
	*problem
}

func NewUnauthorized(e ErrorType) *ErrUnauthorized {
	return &ErrUnauthorized{
		problem: &problem{
			errorType: e,
		},
	}
}

func (e *ErrUnauthorized) WithDetail(d string) *ErrUnauthorized {
	e.detail = d
	return e
}

func (e *ErrUnauthorized) Wrap(err error) *ErrUnauthorized {
	e.cause = err
	return e
}
