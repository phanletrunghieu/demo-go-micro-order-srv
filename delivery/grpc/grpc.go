package grpc

import (
	micro "github.com/micro/go-micro/v2"
	pb "github.com/phanletrunghieu/demo-go-micro-order-srv/proto/order"
)

func New() micro.Service {
	srv := micro.NewService(
		micro.Name("srv.grpc.order"),
	)

	srv.Init()

	pb.RegisterOrderServiceHandler(srv.Server(), &Handler{})

	return srv
}
