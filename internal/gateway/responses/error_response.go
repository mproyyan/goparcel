package responses

import (
	"net/http"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ErrorResponse struct {
	Error      string       `json:"error"`
	Message    string       `json:"message"`
	StackTrace []StackTrace `json:"stack_trace"`
}

type StackTrace struct {
	Error string `json:"error"`
	Cause string `json:"cause"`
}

func NewErrorResponse(err error) (int, ErrorResponse) {
	st, ok := status.FromError(err)
	if !ok {
		return http.StatusInternalServerError, ErrorResponse{
			Error:   "INTERNAL",
			Message: "An unexpected error occurred",
		}
	}

	// Mapping grpc error to http status code
	var httpStatus int
	switch st.Code() {
	case codes.OK:
		httpStatus = http.StatusOK
	case codes.Canceled:
		httpStatus = http.StatusRequestTimeout
	case codes.Unknown, codes.Internal:
		httpStatus = http.StatusInternalServerError
	case codes.InvalidArgument:
		httpStatus = http.StatusBadRequest
	case codes.DeadlineExceeded:
		httpStatus = http.StatusGatewayTimeout
	case codes.NotFound:
		httpStatus = http.StatusNotFound
	case codes.AlreadyExists:
		httpStatus = http.StatusConflict
	case codes.PermissionDenied:
		httpStatus = http.StatusForbidden
	case codes.Unauthenticated:
		httpStatus = http.StatusUnauthorized
	case codes.ResourceExhausted:
		httpStatus = http.StatusTooManyRequests
	case codes.FailedPrecondition, codes.Aborted:
		httpStatus = http.StatusPreconditionFailed
	case codes.OutOfRange:
		httpStatus = http.StatusBadRequest
	case codes.Unimplemented:
		httpStatus = http.StatusNotImplemented
	case codes.Unavailable:
		httpStatus = http.StatusServiceUnavailable
	case codes.DataLoss:
		httpStatus = http.StatusInternalServerError
	default:
		httpStatus = http.StatusInternalServerError
	}

	return httpStatus, parseErrorMessage(httpStatus, st.Message())
}

func parseErrorMessage(code int, errorMessage string) ErrorResponse {
	parts := strings.Split(errorMessage, ", cause: ")
	stackTrace := []StackTrace{}

	// Iterate to create stack trace
	for i := 0; i < len(parts)-1; i++ {
		stackTrace = append(stackTrace, StackTrace{
			Error: parts[i],
			Cause: parts[i+1],
		})
	}

	// The last message is the root cause
	message := parts[len(parts)-1]

	// Return error response
	return ErrorResponse{
		Error:      http.StatusText(code),
		Message:    message,
		StackTrace: stackTrace,
	}
}

func NewInvalidRequestBodyErrorResponse(err error) ErrorResponse {
	return ErrorResponse{
		Error:   "Invalid request body",
		Message: err.Error(),
	}
}
