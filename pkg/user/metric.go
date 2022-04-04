// Code generated by git.repo.services.lenvendo.ru/grade-factor/go-kit-service-generator  REMOVE THIS STRING ON EDIT OR DO NOT EDIT.
package user

import (
	"context"
	"strconv"
	"time"

	tool "github.com/kagestonedragon/server/tools/metrics"
	"github.com/go-kit/kit/metrics"
)

// NewMetricService returns an instance of an instrumenting Service.
func NewMetricsService(ctx context.Context, s Service) Service {
	counter, latency := tool.FromContext(ctx)
	return &metricService{counter, latency, s}
}

type metricService struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	Service
}

func (s *metricService) AddUser(ctx context.Context, req *AddUserRequest) (resp *User, err error) {
	defer func(begin time.Time) {
		go func() {
			s.requestCount.With("service", "user", "handler", "AddUser", "code", strconv.Itoa(getHTTPStatusCode(err))).Add(1)
			s.requestLatency.With("service", "user", "handler", "AddUser", "code", strconv.Itoa(getHTTPStatusCode(err))).Observe(time.Since(begin).Seconds())
		}()
	}(time.Now())
	return s.Service.AddUser(ctx, req)
}

func (s *metricService) GetUserList(ctx context.Context, req *GetUserListRequest) (resp *GetUserListResponse, err error) {
	defer func(begin time.Time) {
		go func() {
			s.requestCount.With("service", "user", "handler", "GetUserList", "code", strconv.Itoa(getHTTPStatusCode(err))).Add(1)
			s.requestLatency.With("service", "user", "handler", "GetUserList", "code", strconv.Itoa(getHTTPStatusCode(err))).Observe(time.Since(begin).Seconds())
		}()
	}(time.Now())
	return s.Service.GetUserList(ctx, req)
}

func (s *metricService) UpdateUser(ctx context.Context, req *User) (resp *UpdateUserResponse, err error) {
	defer func(begin time.Time) {
		go func() {
			s.requestCount.With("service", "user", "handler", "UpdateUser", "code", strconv.Itoa(getHTTPStatusCode(err))).Add(1)
			s.requestLatency.With("service", "user", "handler", "UpdateUser", "code", strconv.Itoa(getHTTPStatusCode(err))).Observe(time.Since(begin).Seconds())
		}()
	}(time.Now())
	return s.Service.UpdateUser(ctx, req)
}

func (s *metricService) DeleteUserById(ctx context.Context, req *DeleteUserByIdRequest) (resp *DeleteUserByIdResponse, err error) {
	defer func(begin time.Time) {
		go func() {
			s.requestCount.With("service", "user", "handler", "DeleteUserById", "code", strconv.Itoa(getHTTPStatusCode(err))).Add(1)
			s.requestLatency.With("service", "user", "handler", "DeleteUserById", "code", strconv.Itoa(getHTTPStatusCode(err))).Observe(time.Since(begin).Seconds())
		}()
	}(time.Now())
	return s.Service.DeleteUserById(ctx, req)
}

func (s *metricService) GetUserById(ctx context.Context, req *GetUserByIdRequest) (resp *User, err error) {
	defer func(begin time.Time) {
		go func() {
			s.requestCount.With("service", "user", "handler", "GetUserById", "code", strconv.Itoa(getHTTPStatusCode(err))).Add(1)
			s.requestLatency.With("service", "user", "handler", "GetUserById", "code", strconv.Itoa(getHTTPStatusCode(err))).Observe(time.Since(begin).Seconds())
		}()
	}(time.Now())
	return s.Service.GetUserById(ctx, req)
}
