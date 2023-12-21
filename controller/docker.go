package controller

import (
	"cc_server/model/req"
	"cc_server/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DockerController struct {
}

func (d *DockerController) GetImages(ctx *gin.Context) {
	is, err := service.DockerSrv.GetImages()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Fail(GetDockerInfoFailCode, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, Success(is))
}

func (d *DockerController) GetContainer(ctx *gin.Context) {
	c, err := service.DockerSrv.GetContainer()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Fail(GetDockerInfoFailCode, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, Success(c))
}

func (d *DockerController) CreateContainer(ctx *gin.Context) {
	var req req.CreateContainerReq
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, Fail(InvalidParamCode, err.Error()))
		return
	}
	if err := service.DockerSrv.CreateContainer(req.Host, req.Image); err != nil {
		ctx.JSON(http.StatusInternalServerError, Fail(CreateContainerFailCode, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, Success(nil))
}
