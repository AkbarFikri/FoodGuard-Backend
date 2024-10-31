package token

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strings"
	"time"
)

const (
	AccessTokenSecret = "ACCESS_TOKEN_SECRET"
)

func Sign(Data map[string]interface{}, ExpiredAt time.Duration) (string, error) {
	expiredAt := time.Now().Add(ExpiredAt).Unix()

	JWTSecretKey := os.Getenv(AccessTokenSecret)

	claims := jwt.MapClaims{}
	claims["exp"] = expiredAt
	claims["authorization"] = true

	for i, v := range Data {
		claims[i] = v
	}

	to := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := to.SignedString([]byte(JWTSecretKey))

	if err != nil {
		return accessToken, err
	}

	return accessToken, nil
}

func VerifyTokenHeader(c *fiber.Ctx) (*jwt.Token, error) {
	header := c.Get("Authorization")
	accessToken := strings.SplitAfter(header, "Bearer")[1]
	JWTSecretKey := os.Getenv(AccessTokenSecret)

	token, err := jwt.Parse(strings.Trim(accessToken, " "), func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTSecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
