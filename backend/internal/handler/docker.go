package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luckyss6/cc_server/internal/service"
)

type DockerHandler struct {
	srv *service.DockerService
}

func NewDockerHandler() *DockerHandler {
	return &DockerHandler{
		srv: service.NewDockerService(),
	}
}

// @Summary 获取镜像列表
func (d *DockerHandler) GetImageList(ctx *fiber.Ctx) error {
	imageList, err := d.srv.GetImageList(ctx.Context())
	if err != nil {
		return ctx.JSON(fail(DockerErrCode, DockerErrMsg))
	}
	return ctx.JSON(success(fiber.Map{"data": imageList, "total": len(imageList)}))
}

// @Summary 获取容器列表
func (d *DockerHandler) GetContainerList(ctx *fiber.Ctx) error {
	containerList, err := d.srv.GetContainerList(ctx.Context())
	if err != nil {
		return ctx.JSON(fail(DockerErrCode, DockerErrMsg))
	}
	return ctx.JSON(success(fiber.Map{"data": containerList, "total": len(containerList)}))
}

// @Summary 获取容器信息
func (d *DockerHandler) GetContainerInfo(ctx *fiber.Ctx) error {
	cj, err := d.srv.GetContainerInfo(ctx.Context(), ctx.Params("id"))
	if err != nil {
		return ctx.JSON(fail(DockerErrCode, DockerErrMsg))
	}
	return ctx.JSON(success(fiber.Map{"data": cj}))
}


