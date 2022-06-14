package build

import (
	"fmt"

	"yawaraka-tissue/domain/problem"
	api "yawaraka-tissue/gen/openapi"
)

const (
	basePath = "https://problems.example.com"
)

func FatalError() *api.Error {
	msg := "Server Error"
	return &api.Error{
		Type:   errorType(problem.TypeServerError),
		Title:  "Something wrong occurs",
		Detail: &msg,
	}
}

func Error(err problem.Problem) *api.Error {
	e := &api.Error{
		Type:  errorType(err.Type()),
		Title: err.Title(),
	}

	if err.Detail() != "" {
		d := err.Detail()
		e.Detail = &d
	}

	return e
}

func errorType(t problem.ErrorType) string {
	return fmt.Sprintf("%s%s", basePath, t)
}

func ValidationError(err *problem.ErrValidationError) *api.ValidationError {
	return &api.ValidationError{
		Type:          errorType(err.Type()),
		Title:         err.Title(),
		InvalidParams: invalidParams(err.InvalidParams),
	}
}

func invalidParams(ip []*problem.InvalidParam) *[]api.InvalidParam {
	if len(ip) == 0 {
		return nil
	}

	ret := make([]api.InvalidParam, 0, len(ip))

	for _, p := range ip {
		ret = append(ret, api.InvalidParam{
			Name:   &p.Name,
			Reason: &p.Reason,
		})
	}

	return &ret
}
