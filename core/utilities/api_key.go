package utilities

import (
	"fmt"
	"kkl-v2/api/middleware"
	"kkl-v2/bootstrap"
	"kkl-v2/core/utilities/encryption"

	"github.com/gofiber/fiber/v2"
)

func ApiKeyAuth(ctx *fiber.Ctx) error {
	if ctx.Get("x-api-key") == "" {
		return middleware.NewErrorUnauthorized(ctx)
	}
	encrypt := encryption.EncodeB64(bootstrap.GlobalEnv.API.Key)

	if ctx.Get("x-api-key") != encrypt {
		return middleware.NewErrorUnauthorized(ctx)
	}

	return ctx.Next()
}

// ApiKeyAuthFromParam validates API Key from URL query or route params named "api_key".
func ApiKeyAuthFromParam(ctx *fiber.Ctx) error {
	apiKey := ctx.Query("api_key")
	if apiKey == "" {
		apiKey = ctx.Params("api_key")
	}
	fmt.Println(apiKey)
	if apiKey == "" {
		return middleware.NewErrorUnauthorized(ctx)
	}

	encrypt := encryption.EncodeB64(bootstrap.GlobalEnv.API.Key)
	if apiKey != encrypt {
		return middleware.NewErrorUnauthorized(ctx)
	}

	return ctx.Next()
}

// ApiKeyAuthAny validates API Key from header (x-api-key), query (?api_key=), or route param (/:api_key)
func ApiKeyAuthAny(ctx *fiber.Ctx) error {
	apiKey := ctx.Get("x-api-key")
	if apiKey == "" {
		apiKey = ctx.Query("api_key")
	}
	if apiKey == "" {
		apiKey = ctx.Params("api_key")
	}
	if apiKey == "" {
		apiKey = ctx.Cookies("x-api-key")
	}
	if apiKey == "" {
		return middleware.NewErrorUnauthorized(ctx)
	}

	encrypt := encryption.EncodeB64(bootstrap.GlobalEnv.API.Key)
	if apiKey != encrypt {
		return middleware.NewErrorUnauthorized(ctx)
	}

	// Persist the validated key in a cookie so subsequent asset requests under /docs/* succeed in browser
	ctx.Cookie(&fiber.Cookie{
		Name:     "x-api-key",
		Value:    apiKey,
		Path:     "/",
		HTTPOnly: true,
	})

	return ctx.Next()
}
