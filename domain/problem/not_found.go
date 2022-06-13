package problem

type ErrNotFound struct {
	*problem
}

func NewNotFound(errType ErrorType) *ErrNotFound {
	return &ErrNotFound{
		problem: &problem{
			errorType: errType,
		},
	}
}

func (e *ErrNotFound) WithDetail(d string) *ErrNotFound {
	e.problem.detail = d
	return e
}

func (e *ErrNotFound) Wrap(err error) *ErrNotFound {
	e.problem.cause = err
	return e
}
