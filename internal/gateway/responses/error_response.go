package responses

import (
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func NewErrorResponse(err error) (int, ErrorResponse) {
	st, ok := status.FromError(err)
	if !ok {
		return http.StatusInternalServerError, ErrorResponse{
			Error:   "INTERNAL",
			Message: "An unexpected error occurred",
		}
	}

	// Mapping gRPC codes ke HTTP status codes
	var httpStatus int
	switch st.Code() {
	case codes.NotFound:
		httpStatus = http.StatusNotFound
	case codes.InvalidArgument:
		httpStatus = http.StatusBadRequest
	case codes.AlreadyExists:
		httpStatus = http.StatusConflict
	case codes.PermissionDenied:
		httpStatus = http.StatusForbidden
	case codes.Unauthenticated:
		httpStatus = http.StatusUnauthorized
	case codes.DeadlineExceeded:
		httpStatus = http.StatusGatewayTimeout
	case codes.Unavailable:
		httpStatus = http.StatusServiceUnavailable
	case codes.Internal:
		httpStatus = http.StatusInternalServerError
	default:
		httpStatus = http.StatusInternalServerError
	}

	return httpStatus, ErrorResponse{
		Error:   st.Code().String(),
		Message: st.Message(),
	}
}

func NewInvalidRequestBodyErrorResponse(err error) ErrorResponse {
	return ErrorResponse{
		Error:   "Invalid request body",
		Message: err.Error(),
	}
}
