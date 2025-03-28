package biz

import (
	"context"
)

type User struct {
	ID           int64
	Username     string
	PasswordHash string
}

type LoginRequest struct {
	Code string	`json:"code,omitempty"` //结构体标签
	State string	`json:"state,omitempty"`
}

type LoginResponse struct {
	Token string	`json:"token,omitempty"`
}


func (uc *UserUseCase) LoginUser(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	uc.log.WithContext(ctx).Infof("LoginUser request:%+v",req)
	return uc.repo.LoginUser(ctx, req)
}

func (uc *UserUseCase) GetUserByUsername(ctx context.Context, username string) (*User, error) {
	return uc.repo.GetByUsername(ctx, username)
}