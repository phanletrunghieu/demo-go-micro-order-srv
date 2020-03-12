package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/phanletrunghieu/demo-go-micro-order-srv/delivery/grpc"
	"github.com/phanletrunghieu/demo-go-micro-order-srv/delivery/http"
)

func main() {
	grpcSrv := grpc.New()
	httpSrv := http.New()

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		errs <- grpcSrv.Run()
	}()
	go func() {
		errs <- httpSrv.Run()
	}()

	log.Println("exit", <-errs)
}
