package api

import (
	"context"
	apiv1 "project/gen/api/v1"
	"project/gen/api/v1/apiv1connect"
	"project/gen/common"

	"connectrpc.com/connect"
)

type service struct{}

func NewService() *service {
	return &service{}
}

func (s *service) Hello(context.Context, *connect.Request[common.NoRequest]) (*connect.Response[apiv1.HelloResponse], error) {
	return connect.NewResponse(&apiv1.HelloResponse{
		Text: "Hello world",
	}), nil
}

var _ apiv1connect.ApiServiceHandler = (*service)(nil)
