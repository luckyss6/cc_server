package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/luckyss6/cc_server/internal/service"
	"github.com/luckyss6/cc_server/pkg/model/request"
)

type UserHandler struct {
	srv *service.UserService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		srv: service.NewUserService(),
	}
}

// @Summary 用户登录
func (u *UserHandler) Login(c *fiber.Ctx) error {
	req := new(request.UserLoginReq)
	if err := c.BodyParser(req); err != nil {
		return c.JSON(fail(BindingParamErrCode, BindingParamErrMsg))
	}

	if err := validator.New().Struct(req); err != nil {
		return c.JSON(fail(BindingParamErrCode, BindingParamErrMsg))
	}

	isSuccess, err := u.srv.Login(req.Username, req.Password)
	if err != nil {
		return c.JSON(fail(DBErrCode, DBErrMsg))
	}

	if !isSuccess {
		return c.JSON(fail(LoginErrCode, LoginErrMsg))
	}

	// create token
	token, err := u.srv.CreateToken(req.Username)
	if err != nil {
		return c.JSON(fail(TokenErrCode, TokenErrMsg))
	}
	return c.JSON(success(map[string]interface{}{"token": token}))
}

// @Summary 用户注册
func (u *UserHandler) Register(c *fiber.Ctx) error {
	req := new(request.UserRegisterReq)
	if err := c.BodyParser(req); err != nil {
		return c.JSON(fail(BindingParamErrCode, BindingParamErrMsg))
	}
	if err := validator.New().Struct(req); err != nil {
		return c.JSON(fail(BindingParamErrCode, BindingParamErrMsg))
	}

	err := u.srv.Register(req.Username, req.Password, req.Email)
	if err != nil {
		return c.JSON(fail(DBErrCode, DBErrMsg))
	}
	return c.JSON(success(nil))
}

// @Summary 用户退出
func (u *UserHandler) Logout() {
}
