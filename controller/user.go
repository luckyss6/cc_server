package controller

import (
	"cc_server/model/req"
	"cc_server/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

// 用户登录
func (u *UserController) Login(ctx *gin.Context) {
	var req req.UserLoginReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Fail(InvalidParamCode, err.Error()))
		return
	}

	// 从数据库中查询用户信息
	isSuccess, userInfo, err := service.UserSrv.GetUserInfo(req.Nickname, req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Fail(DBErrorCode, err.Error()))
		return
	}
	if !isSuccess {
		ctx.JSON(http.StatusOK, Fail(LoginFailCode, "用户名或密码错误"))
		return
	}
	fmt.Println(userInfo)
	// 创建token

	ctx.JSON(http.StatusOK, Success(nil))
}
