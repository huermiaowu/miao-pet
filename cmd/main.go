package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/huermiaowu/miao-pet/internal/app"
	"github.com/huermiaowu/miao-pet/internal/config"
	"github.com/huermiaowu/miao-pet/internal/handler"
	"github.com/huermiaowu/miao-pet/internal/svc"
	pet "github.com/huermiaowu/miao-pet/pb"
	"github.com/huerni/gmitex/pkg/http/handlers"
	"github.com/huerni/gmitex/pkg/http/response"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	c, err := config.InitConfig("etc/cfg.toml")
	if err != nil {
		panic(err)
	}

	g := app.NewGmServer(c, pet.RegisterPetHandlerFromEndpoint, func(server *grpc.Server) {
		ctx := svc.NewServiceContext(c)
		srv := handler.NewPetServer(ctx)
		pet.RegisterPetServer(server, srv)
	})

	g.HttpServer.AddMuxOp(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.HTTPBodyMarshaler{
			Marshaler: &response.CustomMarshaler{
				M: &runtime.JSONPb{
					MarshalOptions:   protojson.MarshalOptions{},
					UnmarshalOptions: protojson.UnmarshalOptions{},
				}}}),
		runtime.WithErrorHandler(handlers.ErrorHandler),
	)

	g.Start(context.Background())
	g.WaitForShutdown(context.Background())
}
