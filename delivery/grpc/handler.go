package grpc

import (
	"context"
	"errors"

	pb "github.com/phanletrunghieu/demo-go-micro-order-srv/proto/order"
)

// mock
var orders map[string]*pb.Order

func init() {
	orders = make(map[string]*pb.Order)
}

type Handler struct {
	// db *gorm.DB
}

func (h *Handler) CreateOrder(ctx context.Context, order *pb.Order, response *pb.Response) error {
	if order == nil {
		return errors.New("order cannot nil")
	}

	orders[order.Id] = order

	response.Success = true
	response.Order = order
	return nil
}

func (h *Handler) UpdateStatus(ctx context.Context, order *pb.Order, response *pb.Response) error {
	if order == nil || (order != nil && order.Id == "") {
		return errors.New("order_id cannot nil")
	}

	orders[order.Id].Status = order.Status

	response.Success = true
	response.Order = order
	return nil
}
