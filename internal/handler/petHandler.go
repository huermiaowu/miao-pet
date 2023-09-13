// no edit

package handler

import (
	"context"

	"pet/internal/service"
	"pet/internal/svc"
	"pet/pb"
)

type PetServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedPetServer
}

func NewPetServer(svcCtx *svc.ServiceContext) *PetServer {
	return &PetServer{
		svcCtx: svcCtx,
	}
}

// Ping /api/v1/pet/ping
func (s *PetServer) Ping(ctx context.Context, in *pb.PingReq) (*pb.PingResp, error) {
	res, err := service.NewPingService(ctx, s.svcCtx).Ping(in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// AddPet /api/v1/pet/addpet
func (s *PetServer) AddPet(ctx context.Context, in *pb.AddPetReq) (*pb.AddPetResp, error) {
	res, err := service.NewAddPetService(ctx, s.svcCtx).AddPet(in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetPet /api/v1/pet
func (s *PetServer) GetPet(ctx context.Context, in *pb.GetPetReq) (*pb.GetPetResp, error) {
	res, err := service.NewGetPetService(ctx, s.svcCtx).GetPet(in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UpdatePet /api/v1/pet/update
func (s *PetServer) UpdatePet(ctx context.Context, in *pb.UpdatePetReq) (*pb.UpdatePetResp, error) {
	res, err := service.NewUpdatePetService(ctx, s.svcCtx).UpdatePet(in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
