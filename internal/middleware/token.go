package middleware

import (
	"github.com/AkbarFikri/FoodGuard-Backend/internal/dto"
	"github.com/AkbarFikri/FoodGuard-Backend/internal/pkg/response"
	"github.com/AkbarFikri/FoodGuard-Backend/internal/pkg/token"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

var (
	ErrUnauthorized = response.NewError(http.StatusUnauthorized, "unauthorized, access token invalid or expired")
)

type tokenMiddleware struct {
}

func newTokenMiddleware() *tokenMiddleware {
	return &tokenMiddleware{}
}

func (m *middleware) NewtokenMiddleware(ctx *fiber.Ctx) error {
	clientIP := ctx.IP()

	if ctx.Get("Authorization") == "" {
		m.log.Warnf("authorization header is missing, client ip is %s", clientIP)
		return ErrUnauthorized
	}

	if !strings.Contains(ctx.Get("Authorization"), "Bearer") {
		m.log.Warnf("authorization header is missing, client ip is %s", clientIP)
		return ErrUnauthorized
	}

	userToken, err := token.VerifyTokenHeader(ctx)
	if err != nil {
		return ErrUnauthorized
	} else {
		claims := userToken.Claims.(jwt.MapClaims)
		user := dto.UserClaims{
			Username: claims["username"].(string),
			Email:    claims["email"].(string),
			ID:       claims["id"].(string),
		}
		ctx.Locals("user", user)
		return ctx.Next()
	}
}
