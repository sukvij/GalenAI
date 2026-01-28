package response

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	ErrBadRequest   = "BAD_REQUEST"
	ErrUnauthorized = "UNAUTHORIZED"
	ErrForbidden    = "FORBIDDEN"
	ErrNotFound     = "NOT_FOUND"
	ErrConflict     = "CONFLICT"
	ErrInternal     = "INTERNAL_ERROR"
	ErrDatabase     = "DATABASE_ERROR"
)

var (
	ErrInvalidInput   = errors.New("invalid input")
	ErrUnauthorizedOp = errors.New("unauthorized operation")
)

type APIResponse struct {
	Success bool        `json:"success"`
	Result  interface{} `json:"result,omitempty"`
	Error   *APIError   `json:"error,omitempty"`
}

type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func MapError(err error) (int, APIResponse) {
	switch {
	case err == nil:
		return 200, APIResponse{Success: true}

	case errors.Is(err, gorm.ErrRecordNotFound):
		return 404, APIResponse{
			Success: false,
			Error: &APIError{
				Code:    ErrNotFound,
				Message: "record not found",
			},
		}

	case errors.Is(err, ErrInvalidInput):
		return 400, APIResponse{
			Success: false,
			Error: &APIError{
				Code:    ErrBadRequest,
				Message: err.Error(),
			},
		}

	case errors.Is(err, ErrUnauthorizedOp):
		return 403, APIResponse{
			Success: false,
			Error: &APIError{
				Code:    ErrForbidden,
				Message: err.Error(),
			},
		}

	default:
		return 500, APIResponse{
			Success: false,
			Error: &APIError{
				Code:    ErrInternal,
				Message: err.Error(),
			},
		}
	}
}

func SendResponse(ctx *gin.Context, result interface{}, err error) {
	status, resp := MapError(err)

	if err == nil {
		resp.Result = result
		resp.Success = true
	}

	ctx.JSON(status, resp)
}

func AbortWithStatus(ctx *gin.Context, result interface{}, err error, status int) {
	ctx.AbortWithStatusJSON(status, err)
}
