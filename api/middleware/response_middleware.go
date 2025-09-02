package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var code int
var message string

type AppError struct {
	Status  int
	Message string
}

func (a AppError) Error() string {
	return a.Message
}

func NewError(code int, errMsg string) error {
	return AppError{
		Status:  code,
		Message: errMsg,
	}
}

func ErrorBadRequest(errorMessage string) error {
	return AppError{
		Status:  http.StatusBadRequest,
		Message: errorMessage,
	}
}
func ErrorUnprocessableEntity(errorMessage string) error {
	return AppError{
		Status:  http.StatusUnprocessableEntity,
		Message: errorMessage,
	}
}

func ErrorRequestEntityTooLarge(errorMessage string) error {
	return AppError{
		Status:  http.StatusRequestEntityTooLarge,
		Message: errorMessage,
	}
}
func ErrorExpectationFailed(errorMessage string) error {
	return AppError{
		Status:  http.StatusExpectationFailed,
		Message: errorMessage,
	}
}

func ErrorNotFound(errorMessage string) error {
	return AppError{
		Status:  http.StatusNotFound,
		Message: errorMessage,
	}
}

func ErrorNoContent(errorMessage string) error {
	return AppError{
		Status:  http.StatusNoContent,
		Message: errorMessage,
	}
}
func NewErrorResponses(ctx *fiber.Ctx, err error) error {
	switch e := err.(type) {
	case AppError:
		code = e.Status
		message = e.Message
	case error:
		code = http.StatusUnprocessableEntity
		message = err.Error()
	}
	return ctx.Status(code).JSON(fiber.Map{
		"status": false,
		"error":  message,
	})
}

func NewAppErrorStatusMessage(statusCode int, err error) error {
	return AppError{
		Status:  statusCode,
		Message: err.Error(),
	}
}
func NewErrorMessageResponse(ctx *fiber.Ctx, message interface{}) error {
	return ctx.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
		"status": false,
		"error":  message,
	})
}
func NewErrorErrMsgInternalServerError(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
		"status": false,
		"error":  ErrMsgInternalServerError,
	})
}
func NewErrorErrMsgUnauthorized(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
		"status": false,
		"error":  ErrMsgUnauthorized,
	})
}
func NewErrorErrMsgUnauthorizedErrMsgInvalidToken(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
		"status": false,
		"error":  ErrMsgInvalidAccessToken,
	})
}
func NewErrorBadRequest(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
		"status": false,
		"error":  ErrMsgBadRequest,
	})
}
func NewErrorIDISRequired(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
		"status": false,
		"error":  ErrMsgParamIdIsRequired,
	})
}
func NewErrorUnAuthorizeRole(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusForbidden).JSON(fiber.Map{
		"status": false,
		"error":  YourRoleNotAllowedToAccessThisResource,
	})
}

func NewErrorUnAuthorizePermission(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusForbidden).JSON(fiber.Map{
		"status": false,
		"error":  YourPermissionNotAllowedToAccessThisResource,
	})
}

func NewSuccessResponse(ctx *fiber.Ctx, data interface{}) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"status": true,
		"data":   data,
	})
}

func NewSuccessMessageResponse(ctx *fiber.Ctx, message interface{}) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"status": true,
		"data":   message,
	})
}

func NewErrorUnauthorized(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error":  "Unauthorized",
		"status": false,
	})
}

func NewErrorUnprocessableEntity(errorMessage string) error {
	return AppError{
		Status:  http.StatusUnprocessableEntity,
		Message: errorMessage,
	}
}
