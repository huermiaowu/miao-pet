package service

import (
	"context"
	"github.com/huermiaowu/miao-pet/internal/db"
	"github.com/huermiaowu/miao-pet/internal/svc"
	"github.com/huermiaowu/miao-pet/pb"
	"github.com/huerni/gmitex/pkg/errno"
)

type GetPetService struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPetService(ctx context.Context, svcCtx *svc.ServiceContext) *GetPetService {
	return &GetPetService{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (s *GetPetService) GetPet(in *pb.GetPetReq) (*pb.GetPetResp, error) {
	// todo: add your logic here and delete this line
	petInfos := make([]*pb.PetInfo, 0)
	var pets []*db.Pet
	var err error
	if in.Id != nil {
		pets, err = db.QueryPetById(s.ctx, *in.Id)
		if err != nil {
			return nil, errno.ConvertErr(err)
		}
	} else if in.UserId != nil {
		pets, err = db.QueryPetByUserId(s.ctx, *in.UserId)
		if err != nil {
			return nil, errno.ConvertErr(err)
		}
	} else {
		pets, err = db.AllPets(s.ctx)
		if err != nil {
			return nil, errno.ConvertErr(err)
		}
	}

	for _, pet := range pets {
		petInfos = append(petInfos, &pb.PetInfo{
			Id:     uint64(pet.ID),
			Name:   pet.Name,
			Sex:    pet.Sex,
			Age:    pet.Age,
			UserId: pet.UserId,
		})
	}
	return &pb.GetPetResp{PetInfos: petInfos}, nil
}
