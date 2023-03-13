package handler

import (
	"strconv"
	"warehouse/model"
	"warehouse/repository"

	// "warehousesystem/entity"
	// "warehousesystem/repository"

	"github.com/gofiber/fiber/v2"
)

type OutboundHandler struct {
	Repository  *repository.OutboundRepository
	App			*fiber.App
}

func NewOutboundHandler(repository *repository.OutboundRepository, app *fiber.App) *OutboundHandler  {
	return &OutboundHandler{Repository: repository, App: app}
}

func (h *OutboundHandler) SetupRoutes() {
	h.App.Get("/outbound/api", h.FindAll)
	h.App.Get("/outbound", h.FindById)
	h.App.Post("/outbound", h.Save)
	h.App.Delete("/outbound", h.Delete)
	h.App.Put("/outbound", h.Update)
}

func (h *OutboundHandler) FindAll(c *fiber.Ctx) error  {
	outbounds, err := h.Repository.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(outbounds)
}

func (h *OutboundHandler) FindById(c *fiber.Ctx) error  {
	id := c.Params("id")
	idInuint, err := strconv.ParseUint(id, 10, 64)

	outbound, err := h.Repository.FindById(uint(idInuint))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(outbound)
}

func (h *OutboundHandler) Save(c *fiber.Ctx) error  {
	outbound := new(model.Outbound)
	if err := c.BodyParser(outbound); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(outbound)
}

func (h *OutboundHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	idInUint, err := strconv.ParseUint(id, 10, 64)

	outbound, err := h.Repository.FindById(uint(idInUint))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if err := c.BodyParser(&outbound); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if err := h.Repository.Update(&outbound); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(outbound)	
}

func (h *OutboundHandler) Delete(c *fiber.Ctx) error  {
	id := c. Params("id")
	idInUint, err := strconv.ParseUint(id, 10, 64)

	outbound, err := h.Repository.FindById(uint(idInUint))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if err := h.Repository.Delete(&outbound); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
}