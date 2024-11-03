package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type middleware struct {
	token        *tokenMiddleware
	rateLimitter *rateLimiter
	log          *logrus.Logger
}

type Middleware interface {
	NewRateLimitter(ctx *fiber.Ctx) error
	NewtokenMiddleware(ctx *fiber.Ctx) error
}

func New(logger *logrus.Logger) Middleware {
	rateLimit := newRateLimiter(50, 100)
	token := newTokenMiddleware()

	return &middleware{
		token:        token,
		rateLimitter: rateLimit,
		log:          logger,
	}
}