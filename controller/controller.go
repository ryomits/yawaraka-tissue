package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"yawaraka-tissue/controller/build"
	"yawaraka-tissue/domain/problem"

	"github.com/pkg/errors"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func RespondOK(ctx context.Context, w http.ResponseWriter, result interface{}) {
	if result == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(result)

	if err != nil {
		RespondError(ctx, w, errors.Wrap(err, "respond_ok"))

		return
	}
}

func RespondError(ctx context.Context, w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/problem+json")

	var br *problem.ErrBadRequest
	if errors.As(err, &br) {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(build.Error(br))
		return
	}

	var ve *problem.ErrValidationError
	if errors.As(err, &ve) {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(build.ValidationError(ve))
		return
	}

	//var fb *problem.ErrForbidden
	//if errors.As(err, &fb) {
	//	w.WriteHeader(http.StatusForbidden)
	//	_ = json.NewEncoder(w).Encode(build.Error(fb))
	//	return
	//}

	//var nf *problem.ErrNotFound
	//if errors.As(err, &nf) {
	//	w.WriteHeader(http.StatusNotFound)
	//	_ = json.NewEncoder(w).Encode(build.Error(nf))
	//	return
	//}

	//var se *problem.ErrInternalServerError
	//if errors.As(err, &se) {
	//	w.WriteHeader(http.StatusInternalServerError)
	//	_ = json.NewEncoder(w).Encode(build.Error(se))
	//	return
	//}

	w.WriteHeader(http.StatusInternalServerError)
	_ = json.NewEncoder(w).Encode(build.FatalError())
}
