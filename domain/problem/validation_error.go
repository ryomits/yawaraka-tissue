package problem

type ErrValidationError struct {
	*problem
	InvalidParams []*InvalidParam
}

type InvalidParam struct {
	Name   string
	Reason string
}

func NewValidationError(e ErrorType) *ErrValidationError {
	return &ErrValidationError{
		problem: &problem{
			errorType: e,
		},
		InvalidParams: make([]*InvalidParam, 0),
	}
}

func (e *ErrValidationError) WithDetaul(d string) *ErrValidationError {
	e.detail = d
	return e
}

func (e *ErrValidationError) Wrap(err error) *ErrValidationError {
	e.cause = err
	return e
}

func (e *ErrValidationError) WithInvalidParams(ps []*InvalidParam) *ErrValidationError {
	e.InvalidParams = append(e.InvalidParams, ps...)
	return e
}
