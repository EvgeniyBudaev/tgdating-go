package v1

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/entity"
	"github.com/gofiber/fiber/v2"
)

func ResponseError(ctx *fiber.Ctx, err error, httpStatusCode int) error {
	return ctx.Status(httpStatusCode).JSON(err.Error())
}

func ResponseFieldsError(ctx *fiber.Ctx, err *entity.ValidationErrorEntity) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(err)
}

func ResponseOk(ctx *fiber.Ctx, data interface{}) error {
	return ctx.Status(fiber.StatusOK).JSON(data)
}

func ResponseCreated(ctx *fiber.Ctx, data interface{}) error {
	return ctx.Status(fiber.StatusCreated).JSON(data)
}

func ResponseImage(ctx *fiber.Ctx, data []byte) error {
	return ctx.Status(fiber.StatusOK).Send(data)
}
