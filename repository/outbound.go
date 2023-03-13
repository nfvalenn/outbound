package repository

import (
	"warehouse/model"

	"gorm.io/gorm"
)

type OutboundRepository struct {
	db *gorm.DB
}

func NewOutboundRepository(db *gorm.DB) OutboundRepository  {
	return OutboundRepository{db}
}

func (r *OutboundRepository) Save(outbound *model.Outbound) error {
	return r.db.Create(&outbound).Error
}

func (r *OutboundRepository) FindAll() ([]model.Outbound, error)  {
	var outbounds []model.Outbound
	err := r.db.Find(&outbounds).Error
		return outbounds, err
}

func (r *OutboundRepository) FindById(id uint) (model.Outbound, error)  {
	var outbound model.Outbound
	err := r.db.First(&outbound, id).Error
		return outbound, err
}

func (r *OutboundRepository) Update(outbound *model.Outbound) error {
	return r.db.Save(outbound).Error
}

func (r *OutboundRepository) Delete(outbound *model.Outbound) error {
	return r.db.Delete(outbound).Error
}

