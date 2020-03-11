package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	micro "github.com/micro/go-micro/v2"
	"github.com/phanletrunghieu/demo-go-micro-order-srv/delivery/grpc"
	pb "github.com/phanletrunghieu/demo-go-micro-order-srv/proto/order"
)

func main() {
	srv := micro.NewService(
		micro.Name("service.order"),
	)

	srv.Init()

	pb.RegisterOrderServiceHandler(srv.Server(), &grpc.Handler{})

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		errs <- srv.Run()
	}()

	log.Println("exit", <-errs)
}
