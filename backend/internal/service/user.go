package service

import (
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/luckyss6/cc_server/internal/repository"
	"github.com/luckyss6/cc_server/pkg/model"
	"github.com/luckyss6/cc_server/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repository.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repository.NewUserRepo(),
	}
}

func (u *UserService) Login(username, password string) (bool, error) {
	userInfo, err := u.userRepo.GetUserInfoByUsername(username)
	if err != nil {
		slog.Error(fmt.Sprintf("get user info error: %s", err.Error()))
		return false, err
	}
	if userInfo == nil {
		return false, nil
	}
	if err := bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(password)); err != nil {
		return false, nil
	}
	return true, nil
}

func (u *UserService) Register(username, password, email string) error {
	// create userID
	userID := uuid.New().String()
	// create password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	err := u.userRepo.CreateUserInfo(&model.UserInfo{
		UserID:   userID,
		Username: username,
		Password: string(hashedPassword),
		Email:    email,
	})
	if err != nil {
		slog.Error(fmt.Sprintf("create user info error: %s", err.Error()))
		return err
	}
	return nil
}

func (u *UserService) CreateToken(userID string) (string, error) {
	token, err := utils.CreateToken(userID)
	if err != nil {
		slog.Error(fmt.Sprintf("create token error: %s", err.Error()))
		return "", err
	}
	return token, nil
}
