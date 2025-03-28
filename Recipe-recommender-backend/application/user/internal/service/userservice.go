package service

import (
	"context"
	"fmt"
	pb "user/api/user/v1"
	"user/internal/biz"
)

type UserServiceService struct {
	pb.UnimplementedUserServiceServer
	ac *biz.UserUseCase
}

func NewUserServiceService(ac *biz.UserUseCase) *UserServiceService {
	return &UserServiceService{ac: ac}
}

func (s *UserServiceService) RegisterUser(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return &pb.RegisterResponse{}, nil
}
func (s *UserServiceService) LoginUser(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	result, err := s.ac.LoginUser(ctx, &biz.LoginRequest{
		State: req.State,
		Code:  req.Code,
	})
	fmt.Printf("result:%+v", result)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{Token: result.Token}, nil
}
func (s *UserServiceService) GetUserProfile(ctx context.Context, req *pb.GetUserProfileRequest) (*pb.UserProfileResponse, error) {
	return &pb.UserProfileResponse{}, nil
}
func (s *UserServiceService) UpdateUserProfile(ctx context.Context, req *pb.UpdateUserProfileRequest) (*pb.UserProfileResponse, error) {
	return &pb.UserProfileResponse{}, nil
}
func (s *UserServiceService) GetHealthRecord(ctx context.Context, req *pb.GetHealthRecordRequest) (*pb.HealthRecordResponse, error) {
	return &pb.HealthRecordResponse{}, nil
}
func (s *UserServiceService) UpdateHealthRecord(ctx context.Context, req *pb.UpdateHealthRecordRequest) (*pb.HealthRecordResponse, error) {
	return &pb.HealthRecordResponse{}, nil
}
func (s *UserServiceService) GetUserPreferences(ctx context.Context, req *pb.GetUserPreferencesRequest) (*pb.UserPreferencesResponse, error) {
	return &pb.UserPreferencesResponse{}, nil
}
func (s *UserServiceService) UpdateUserPreferences(ctx context.Context, req *pb.UpdateUserPreferencesRequest) (*pb.UserPreferencesResponse, error) {
	return &pb.UserPreferencesResponse{}, nil
}
func (s *UserServiceService) LogoutUser(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	return &pb.LogoutResponse{}, nil
}
