package db

import (
	"context"
	"gorm.io/gorm"
)

var petTableName = "Pet"

type Pet struct {
	gorm.Model
	Name   string `gorm:"column:name;type:varchar(254);not null"`
	Sex    int32  `gorm:"column:sex;type:Integer; not null;"`
	Age    int32  `gorm:"column:age;type:Integer;"`
	UserId uint64 `gorm:"column:user_id;type:bigint;not null"`
}

func (p *Pet) TableName() string {
	return petTableName
}

func CreatePet(ctx context.Context, pet *Pet) error {
	return DB.WithContext(ctx).Create(pet).Error
}

func UpdatePet(ctx context.Context, pet *Pet) error {
	return DB.WithContext(ctx).Updates(pet).Error
}

func QueryPetById(ctx context.Context, ID uint64) ([]*Pet, error) {
	res := make([]*Pet, 0)
	if err := DB.WithContext(ctx).Where("id = ?", ID).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func QueryPetByUserId(ctx context.Context, userId uint64) ([]*Pet, error) {
	res := make([]*Pet, 0)
	if err := DB.WithContext(ctx).Where("user_id = ?", userId).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func AllPets(ctx context.Context) ([]*Pet, error) {
	res := make([]*Pet, 0)
	if err := DB.WithContext(ctx).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
