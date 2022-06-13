package problem

type ErrorType string

const (
	TypeUnAuthorized     ErrorType = "/request_unauthorized"
	TypeValidationError  ErrorType = "/validation_error"
	TypeBadRequest       ErrorType = "/bad_request"
	TypeResourceNotFound ErrorType = "/resource_not_found"
	TypeServerError      ErrorType = "/server_error"
)

var errorTitle = map[ErrorType]string{
	TypeUnAuthorized:     "API key authenticatication failed",
	TypeValidationError:  "Your request parameters didn't validate",
	TypeBadRequest:       "Your request is invalid",
	TypeResourceNotFound: "Specified resource does not exist",
	TypeServerError:      "Something wrong occurs",
}
