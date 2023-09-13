package router

import (
	"context"
	"github.com/huerni/gmitex/pkg/gw"
)

type PetRouter struct {
	Paths     []string
	AuthPaths []string
}

func NewPetRouter() *PetRouter {
	paths := make([]string, 0)
	authPaths := make([]string, 0)

	paths = append(paths, "/api/v1/pet")

	return &PetRouter{Paths: paths, AuthPaths: authPaths}
}

func (ur *PetRouter) AddRouters(ctx context.Context, client gw.GatewayClient, info any) error {
	return client.AddRoute(ctx, info)
}
