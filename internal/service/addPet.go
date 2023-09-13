package service

import (
	"context"
	"github.com/huermiaowu/miao-pet/internal/db"
	"github.com/huerni/gmitex/pkg/errno"
	"gorm.io/gorm"

	"github.com/huermiaowu/miao-pet/internal/svc"
	"github.com/huermiaowu/miao-pet/pb"
)

type AddPetService struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddPetService(ctx context.Context, svcCtx *svc.ServiceContext) *AddPetService {
	return &AddPetService{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (s *AddPetService) AddPet(in *pb.AddPetReq) (*pb.AddPetResp, error) {
	// todo: add your logic here and delete this line
	pet := db.Pet{
		Model:  gorm.Model{},
		Name:   in.Name,
		Sex:    in.Sex,
		Age:    in.Age,
		UserId: in.UserId,
	}
	err := db.CreatePet(s.ctx, &pet)
	if err != nil {
		return nil, errno.ConvertErr(err)
	}
	petInfo := pb.PetInfo{
		Id:     uint64(pet.ID),
		Name:   pet.Name,
		Sex:    pet.Sex,
		Age:    pet.Age,
		UserId: pet.UserId,
	}
	return &pb.AddPetResp{PetInfo: &petInfo}, nil
}
