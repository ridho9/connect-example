package flightprice

import (
	"context"
	"project/gen/common"
	flightpricev1 "project/gen/flightprice/v1"
	"project/gen/flightprice/v1/flightpricev1connect"

	"connectrpc.com/connect"
)

type service struct{}

func NewService() *service {
	return &service{}
}

func (s *service) Hello(context.Context, *connect.Request[common.NoRequest]) (*connect.Response[flightpricev1.HelloResponse], error) {
	return connect.NewResponse(&flightpricev1.HelloResponse{
		Text: "Hello world",
	}), nil
}

var _ flightpricev1connect.FlightServiceHandler = (*service)(nil)
