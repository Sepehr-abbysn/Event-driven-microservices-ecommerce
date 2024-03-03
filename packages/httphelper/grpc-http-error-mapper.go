package httphelper

import (
	"encoding/json"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ErrorMessage struct {
	Error string `json:"error"`
}

func ParseGrpcToHttpError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")

	if serr, ok := status.FromError(err); ok {
		switch serr.Code() {
		case codes.InvalidArgument:
			w.WriteHeader(http.StatusBadRequest)

		case codes.NotFound:
			w.WriteHeader(http.StatusNotFound)

		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(ErrorMessage{Error: serr.Message()})
		return
	}
}
