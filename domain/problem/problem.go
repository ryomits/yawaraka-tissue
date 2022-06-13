package problem

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type Problem interface {
	Error() string
	Unwrap() error
	Type() ErrorType
	Title() string
	Detail() string
}

type problem struct {
	cause     error
	errorType ErrorType
	detail    string
}

func (e *problem) Error() string {
	msg := strings.ToLower(e.Title())
	if e.detail != "" {
		msg = fmt.Sprintf("%s, caused by `%s`", msg, strings.ToLower(e.detail))
	}
	if e.cause == nil {
		return msg
	}

	return errors.Wrap(e.cause, msg).Error()
}

func (e *problem) Unwrap() error {
	return e.cause
}

func (e *problem) Type() ErrorType {
	return e.errorType
}

func (e *problem) Title() string {
	return errorTitle[e.errorType]
}

func (e *problem) Detail() string {
	return e.detail
}
