package http

import (
	"github.com/micro/go-micro/v2/web"
)

func New() web.Service {

	srv := web.NewService(
		web.Name("srv.api.order"),
	)
	srv.Init()

	h := NewHandler()
	srv.Handle("/", h)
	return srv
}
