package middleware

import (
	"encoding/json"
	"fmt"
	"kkl-v2/bootstrap"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kataras/jwt"
)

type ClaimsToken struct {
	Id        string `json:"id,omitempty"`
	Role      string `json:"role,omitempty"`
	IssuedAt  int64  `json:"iat,omitempty"`
	ExpiresAt int64  `json:"exp,omitempty"`
}

// var (
// 	JWT_ACCESS_TOKEN  = []byte("superdupersecret")
// 	JWT_REFRESH_TOKEN = []byte("superdupersecretrefresher")
// )

type TokenPair struct {
	AccessToken  json.RawMessage `json:"access_token,omitempty"`
	RefreshToken json.RawMessage `json:"refresh_token,omitempty"`
}

func GenerateJWTToken(id string, role string) (*TokenPair, error) {
	standardClaims := ClaimsToken{
		Id:        id,
		Role:      role,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Hour * 24 * 300).Unix(),
	}
	encrypt, _, err := jwt.GCM([]byte(bootstrap.GlobalEnv.JWT.AccessToken), nil)
	if err != nil {

		return nil, err
	}
	token, err := jwt.SignEncrypted(jwt.HS256, []byte(bootstrap.GlobalEnv.JWT.AccessToken), encrypt, standardClaims, jwt.MaxAge(time.Hour*24*7))
	if err != nil {

		return nil, err
	}
	reEncrypt, _, _ := jwt.GCM([]byte(bootstrap.GlobalEnv.JWT.RefreshToken), nil)
	refreshToken, err := jwt.SignEncrypted(jwt.HS256, []byte(bootstrap.GlobalEnv.JWT.RefreshToken), reEncrypt, standardClaims, jwt.MaxAge(time.Hour*24*8))
	if err != nil {

		return nil, err
	}
	tokenPairData := jwt.NewTokenPair(token, refreshToken)
	return &TokenPair{
		AccessToken:  BytesQuote(tokenPairData.AccessToken),
		RefreshToken: BytesQuote(tokenPairData.RefreshToken),
	}, nil
}

func AccessToken(ctx *fiber.Ctx) error {
	_, decrypt, _ := jwt.GCM([]byte(bootstrap.GlobalEnv.JWT.AccessToken), nil)
	auth := ctx.Get("Authorization")
	if auth == "" {
		return NewErrorUnauthorized(ctx)
	}
	jwtFromHeader := strings.TrimSpace(auth[7:])
	verifiedToken, err := jwt.VerifyEncrypted(jwt.HS256, []byte(bootstrap.GlobalEnv.JWT.AccessToken), decrypt, []byte(jwtFromHeader))
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status": false,
			"result": err.Error(),
		})
	}
	var claims ClaimsToken
	err = verifiedToken.Claims(&claims)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status": false,
			"result": err.Error(),
		})
	}
	return ctx.Next()
}

func AccessRefreshToken(ctx *fiber.Ctx) error {
	_, decrypt, _ := jwt.GCM([]byte(bootstrap.GlobalEnv.JWT.RefreshToken), nil)
	auth := ctx.Get("Authorization")
	if auth == "" {
		return NewErrorUnauthorized(ctx)
	}
	jwtFromHeader := strings.TrimSpace(auth[7:])
	verifiedToken, err := jwt.VerifyEncrypted(jwt.HS256, []byte(bootstrap.GlobalEnv.JWT.RefreshToken), decrypt, []byte(jwtFromHeader))
	if err != nil {
		return ctx.Status(fiber.StatusExpectationFailed).JSON(fiber.Map{
			"status": false,
			"result": err.Error(),
		})
	}
	var claims ClaimsToken
	err = verifiedToken.Claims(&claims)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status": false,
			"result": err.Error(),
		})
	}
	return ctx.Next()
}
func GetOwnerAccessToken(c *fiber.Ctx) (string, error) {
	_, decrypt, _ := jwt.GCM([]byte(bootstrap.GlobalEnv.JWT.AccessToken), nil)
	auth := c.Get("Authorization")
	jwtFromHeader := strings.TrimSpace(auth[7:])
	verifiedToken, err := jwt.VerifyEncrypted(jwt.HS256, []byte(bootstrap.GlobalEnv.JWT.AccessToken), decrypt, []byte(jwtFromHeader))
	if err != nil {
		return "", err
	}
	var claims ClaimsToken
	err = verifiedToken.Claims(&claims)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", claims.Id), nil
}
func GetInfoAccessToken(c *fiber.Ctx) (*ClaimsToken, error) {
	_, decrypt, _ := jwt.GCM([]byte(bootstrap.GlobalEnv.JWT.AccessToken), nil)
	auth := c.Get("Authorization")
	jwtFromHeader := strings.TrimSpace(auth[7:])
	verifiedToken, err := jwt.VerifyEncrypted(jwt.HS256, []byte(bootstrap.GlobalEnv.JWT.AccessToken), decrypt, []byte(jwtFromHeader))
	if err != nil {
		return nil, err
	}
	var claims ClaimsToken
	err = verifiedToken.Claims(&claims)
	if err != nil {
		return nil, err
	}
	return &claims, nil
}

func GetOwnerRefresh(c *fiber.Ctx) (*ClaimsToken, error) {
	_, decrypt, _ := jwt.GCM([]byte(bootstrap.GlobalEnv.JWT.RefreshToken), nil)
	auth := c.Get("Authorization")
	if auth == "" {
		return nil, NewErrorUnauthorized(c)
	}
	jwtFromHeader := strings.TrimSpace(auth[7:])

	verifiedToken, err := jwt.VerifyEncrypted(jwt.HS256, []byte(bootstrap.GlobalEnv.JWT.RefreshToken), decrypt, []byte(jwtFromHeader))
	if err != nil {
		return nil, err
	}
	var claims ClaimsToken
	err = verifiedToken.Claims(&claims)
	if err != nil {
		return nil, err
	}
	resp := ClaimsToken{
		Id: claims.Id,
	}
	return &resp, nil
}

// TODO: Refresh token
func GenerateRefreshToken(ctx *fiber.Ctx) (*TokenPair, error) {
	auth := ctx.Get("Authorization")
	if auth == "" {
		return nil, NewErrorUnauthorized(ctx)
	}
	claims, err := GetOwnerRefresh(ctx)
	if err != nil {
		return nil, err
	}
	standardClaims := ClaimsToken{
		Id:        claims.Id,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
	}
	encrypt, _, _ := jwt.GCM([]byte(bootstrap.GlobalEnv.JWT.AccessToken), nil)
	token, err := jwt.SignEncrypted(jwt.HS256, []byte(bootstrap.GlobalEnv.JWT.AccessToken), encrypt, standardClaims, jwt.MaxAge(time.Second*15))
	if err != nil {

		return nil, err
	}
	reEncrypt, _, _ := jwt.GCM([]byte(bootstrap.GlobalEnv.JWT.RefreshToken), nil)
	refreshToken, err := jwt.SignEncrypted(jwt.HS256, []byte(bootstrap.GlobalEnv.JWT.RefreshToken), reEncrypt, standardClaims, jwt.MaxAge(time.Hour*24*8))
	if err != nil {

		return nil, err
	}
	tokenPairData := jwt.NewTokenPair(token, refreshToken)
	return &TokenPair{
		AccessToken:  BytesQuote(tokenPairData.AccessToken),
		RefreshToken: BytesQuote(tokenPairData.RefreshToken),
	}, nil
}

func BytesQuote(b []byte) []byte {
	dst := make([]byte, len(b)+2)
	dst[0] = '"'
	copy(dst[1:], b)
	dst[len(dst)-1] = '"'
	return dst
}
