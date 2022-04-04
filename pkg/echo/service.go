package echo

import (
	"context"
	"github.com/kagestonedragon/server/internal/repository/echo"
)

type echoService struct {
	e echo.Echo
}

func NewEchoService(e echo.Echo) Service {
	return &echoService{e}
}

func (s *echoService) GetEcho(ctx context.Context, req *GetEchoListRequest) (resp *GetEchoListResponse, err error) {
	a := GetEchoListResponse{
		Echos: []Echo{
			{
				Id:       uint32(1),
				Title:    "title",
				Reminder: "Reminder",
			},
		},
	}
	if err != nil {
		return &a, err
	}
	return &a, nil
}
