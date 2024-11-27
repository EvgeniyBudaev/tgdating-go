package v1

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/entity"
	"github.com/gofiber/fiber/v2"
)

func ResponseError(ctf *fiber.Ctx, err error, httpStatusCode int) error {
	return ctf.Status(httpStatusCode).JSON(err.Error())
}

func ResponseFieldsError(ctf *fiber.Ctx, err *entity.ValidationErrorEntity) error {
	return ctf.Status(fiber.StatusBadRequest).JSON(err)
}

func ResponseOk(ctf *fiber.Ctx, data interface{}) error {
	return ctf.Status(fiber.StatusOK).JSON(data)
}

func ResponseCreated(ctf *fiber.Ctx, data interface{}) error {
	return ctf.Status(fiber.StatusCreated).JSON(data)
}

func ResponseImage(ctf *fiber.Ctx, data []byte) error {
	return ctf.Status(fiber.StatusOK).Send(data)
}
