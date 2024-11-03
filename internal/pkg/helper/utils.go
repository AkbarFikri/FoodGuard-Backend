package helper

import (
	"fmt"
	"github.com/AkbarFikri/FoodGuard-Backend/internal/dto"
	"github.com/gofiber/fiber/v2"
)

func GetUserFromContext(ctx *fiber.Ctx) (dto.UserClaims, error) {
	userCtx := ctx.Locals("user")
	if userCtx == nil {
		return dto.UserClaims{}, fmt.Errorf("user not found in context")
	}

	return userCtx.(dto.UserClaims), nil
}
