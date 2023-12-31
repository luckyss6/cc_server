package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luckyss6/cc_server/internal/service"
)

type MachineHandler struct {
	srv *service.MachineService
}

func NewMachineHandler() *MachineHandler {
	return &MachineHandler{
		srv: service.NewMachineService(),
	}
}

func (h *MachineHandler) GetMachineInfo(ctx *fiber.Ctx) error {
	return nil

}
