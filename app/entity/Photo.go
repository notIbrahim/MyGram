package entity

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	IDModels
	User_ID  uint   `gorm:"not null;"`
	Title    string `gorm:"not null;" json:"PhotoTitle,omitempty" form:"PhotoEmail" valid:"required~Your PhotoTitle is required"`
	Caption  string `gorm:"null;" json:"PhotoCaption,omitempty"`
	PhotoURL string `gorm:"not null;" json:"PhotoURL,omitempty" form:"PhotoURL" valid:"required~Your PhotoURL is required"`
	TimeModels
}

func (t_Photo *Photo) BeforeCreate(t_DB *gorm.DB) (err error) {
	_, ErrCreated := govalidator.ValidateStruct(t_Photo)
	if ErrCreated != nil {
		err = ErrCreated
		return
	}

	return nil
}

func (t_Photo *Photo) BeforeUpdate(t_DB *gorm.DB) (err error) {
	_, ErrUpdated := govalidator.ValidateStruct(t_Photo)
	if ErrUpdated != nil {
		err = ErrUpdated
		return
	}

	return nil
}

func (t_Photo *Photo) BeforeDelete(t_DB *gorm.DB) (err error) {
	_, ErrDeleted := govalidator.ValidateStruct(t_Photo)
	if ErrDeleted != nil {
		err = ErrDeleted
		return
	}
	return nil
}
