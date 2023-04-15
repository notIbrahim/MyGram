package entity

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	IDModels
	User_ID  uint   `gorm:"not null;"`
	Photo_ID uint   `gorm:"not null;" `
	Message  string `gorm:"not null;" json:"Comments" form:"Comments" valid:"required~Comments can't be blank"`
	TimeModels
}

func (t_Comment *Comment) BeforeCreate(t_DB *gorm.DB) (err error) {
	_, ErrCreated := govalidator.ValidateStruct(t_Comment)
	if ErrCreated != nil {
		err = ErrCreated
		return
	}

	return nil
}

func (t_Comment *Comment) BeforeUpdate(t_DB *gorm.DB) (err error) {
	_, ErrUpdated := govalidator.ValidateStruct(t_Comment)
	if ErrUpdated != nil {
		err = ErrUpdated
		return
	}

	return nil
}
