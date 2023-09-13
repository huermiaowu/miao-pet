package service

import (
	"context"
	"fmt"
	"github.com/huerni/gmitex/pkg/errno"
	"gorm.io/gorm"
	"pet/internal/db"

	"pet/internal/svc"
	"pet/pb"
)

type UpdatePetService struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePetService(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePetService {
	return &UpdatePetService{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (s *UpdatePetService) UpdatePet(in *pb.UpdatePetReq) (*pb.UpdatePetResp, error) {
	// todo: add your logic here and delete this line
	pets, err := db.QueryPetById(s.ctx, in.Id)
	if err != nil {
		return nil, errno.ConvertErr(err)
	}
	if len(pets) == 0 {
		return nil, errno.Err(int(pb.ErrCode_PetNodFound), fmt.Sprintf("宠物id:%d不存在", in.Id))
	}
	new_pet := db.Pet{
		Model: gorm.Model{
			ID: uint(in.Id),
		},
		Name: in.Name,
		Sex:  in.Sex,
		Age:  in.Age,
	}
	err = db.UpdatePet(s.ctx, &new_pet)
	if err != nil {
		return nil, errno.ConvertErr(err)
	}
	petInfo := pb.PetInfo{
		Id:     uint64(in.Id),
		Name:   in.Name,
		Sex:    in.Sex,
		Age:    in.Age,
		UserId: pets[0].UserId,
	}
	return &pb.UpdatePetResp{PetInfo: &petInfo}, nil
}
