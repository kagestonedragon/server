package user

import (
	"context"
	"github.com/kagestonedragon/server/internal/user"
)

type userService struct {
	r user.CacheableRepository
}

func NewUserService(r user.CacheableRepository) Service {
	return &userService{
		r: r,
	}
}

func (s *userService) AddUser(ctx context.Context, req *AddUserRequest) (resp *User, err error) {
	u := &user.User{
		Name:   req.Name,
		Active: req.Active,
	}

	if err := s.r.Add(ctx, u); err != nil {
		return nil, err
	}

	return s.convertInternalUserToExternal(u), nil
}

func (s *userService) UpdateUser(ctx context.Context, req *User) (resp *UpdateUserResponse, err error) {
	if err := s.r.Update(ctx, s.convertExternalUserToInternal(req)); err != nil {
		return nil, err
	}

	return &UpdateUserResponse{}, nil
}

func (s *userService) DeleteUserById(ctx context.Context, req *DeleteUserByIdRequest) (resp *DeleteUserByIdResponse, err error) {
	if err := s.r.DeleteById(ctx, req.Id); err != nil {
		return nil, err
	}

	return &DeleteUserByIdResponse{}, nil
}

func (s *userService) GetUserList(ctx context.Context, req *GetUserListRequest) (resp *GetUserListResponse, err error) {
	users, err := s.r.GetList(ctx)
	if err != nil {
		return nil, err
	}

	a := GetUserListResponse{}
	for _, u := range users {
		a = append(a, *s.convertInternalUserToExternal(u))
	}

	return &a, nil
}

func (s *userService) GetUserById(ctx context.Context, req *GetUserByIdRequest) (resp *User, err error) {
	u, err := s.r.GetById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return s.convertInternalUserToExternal(u), nil
}

func (s *userService) convertInternalUserToExternal(u *user.User) *User {
	return &User{
		Id:     u.Id,
		Name:   u.Name,
		Active: u.Active,
	}
}

func (s *userService) convertExternalUserToInternal(u *User) *user.User {
	return &user.User{
		Id:     u.Id,
		Name:   u.Name,
		Active: u.Active,
	}
}
