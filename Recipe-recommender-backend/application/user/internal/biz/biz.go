package biz

import (
	"context"
	"github.com/google/wire"
	"github.com/go-kratos/kratos/v2/log"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewUserUseCase)

// 对 User 定义操作接口 UserRepo
type UserRepo interface {
	LoginUser(ctx context.Context, req *LoginRequest) (*LoginResponse, error)
	GetByUsername(ctx context.Context, username string) (*User, error)
}

// 操作加上日志
type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}