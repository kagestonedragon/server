// Code generated by git.repo.services.lenvendo.ru/grade-factor/go-kit-service-generator  REMOVE THIS STRING ON EDIT OR DO NOT EDIT.
//go:generate mockgen -destination service_mock.go -package echo  "github.com/kagestonedragon/server/pkg/echo Service
package echo

import (
	"context"

	_ "github.com/golang/mock/mockgen/model"
)

type Service interface {

	// GetEcho returns buid time, last commit and version app
	GetEcho(context.Context, *GetEchoListRequest) (*GetEchoListResponse, error)
}
