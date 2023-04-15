package entity

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// Before Judge anything here
// My Style Consist Writing can be interchanged because i use both
// 1. CamelCase
// 2. underscore for identifier temp, pointer, and etc

type Social struct {
	IDModels
	Name             string `gorm:"not null" json:"Name" form:"Name" valid:"required~Your name is required"`
	Social_Media_URL string `gorm:"not null" json:"LinkURL" form:"LinkURL" valid:"required~Your LinkURL is required"`
	User_ID          uint   `gorm:"not null; uniqueIndex"`
	TimeModels
}

func (t_Social *Social) BeforeCreate(t_DB *gorm.DB) (err error) {
	_, ErrCreated := govalidator.ValidateStruct(t_Social)
	if ErrCreated != nil {
		err = ErrCreated
		return
	}

	return nil
}

func (t_Social *Social) BeforeUpdate(t_DB *gorm.DB) (err error) {
	_, ErrUpdated := govalidator.ValidateStruct(t_Social)
	if ErrUpdated != nil {
		err = ErrUpdated
		return
	}

	return nil
}

func (t_Social *Social) BeforeDelete(t_DB *gorm.DB) (err error) {
	_, ErrDeleted := govalidator.ValidateStruct(t_Social)
	if ErrDeleted != nil {
		err = ErrDeleted
		return
	}
	return nil
}
