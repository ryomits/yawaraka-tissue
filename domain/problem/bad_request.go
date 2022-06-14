package problem

type ErrBadRequest struct {
	*problem
}

func NewBadRequest(e ErrorType) *ErrBadRequest {
	return &ErrBadRequest{
		problem: &problem{
			errorType: e,
		},
	}
}

func (e *ErrBadRequest) WithDetail(d string) *ErrBadRequest {
	e.problem.detail = d
	return e
}

func (e *ErrBadRequest) Wrap(err error) *ErrBadRequest {
	e.cause = err
	return e
}
