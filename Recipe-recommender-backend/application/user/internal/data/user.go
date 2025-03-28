package data

import (
	"context"
	"errors"
	"fmt"

	"user/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) LoginUser(ctx context.Context, req *biz.LoginRequest) (*biz.LoginResponse, error) {
	// 通过 Casdoor 获取 OAuth 令牌
	if r.data.cs == nil {
        return nil, errors.New("casdoor client is nil") // ✅ 直接返回错误，避免空指针
    }
	code := req.Code
	state := req.State
	token, err := r.data.cs.GetOAuthToken(code, state)
	if err != nil {
		fmt.Println("GetOAuthToken() error", err)
		return nil, errors.New("GetOAuthToken() error:" + err.Error())
	}

	fmt.Println("GetOAuthToken() token", token)
	return &biz.LoginResponse{
		Token: token.AccessToken,
	},nil
}



func (r *userRepo) GetByUsername(ctx context.Context, username string) (*biz.User, error) {
	var user biz.User
	// if err := r.data.DB.WithContext(ctx).
	// 	Where("username = ?", username).
	// 	First(&user).Error; err != nil {
	// 	return nil, err
	// }
	return &user, nil
}
