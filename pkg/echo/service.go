package echo

import (
	"context"
)

type echoService struct {
}

func NewEchoService() Service {
	return &echoService{}
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
