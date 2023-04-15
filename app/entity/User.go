package entity

import (
	"MyGram/pkg/helper"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	IDModels
	UniqueID string `gorm:"not null;uniqueKey" json:"UID" form:"UID"`
	Username string `gorm:"not null;uniqueIndex" json:"username,omitempty" form:"username" valid:"required~Your username is required"`
	Email    string `gorm:"not null;uniqueIndex" json:"email,omitempty" form:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password string `gorm:"not null" json:"password,omitempty" form:"password" valid:"required~Your password is required,minstringlength(6)~Password must be 6 characters or more"`
	Age      uint   `gorm:"not null" json:"age" form:"age,omitempty" valid:"required~Your age is required,range(9|60)~Your age should be above 8 years old"`
	TimeModels
}

func (t_User *User) BeforeCreate(t_DB *gorm.DB) (err error) {
	_, ErrCreated := govalidator.ValidateStruct(t_User)
	if ErrCreated != nil {
		err = ErrCreated
		return
	}
	t_User.Password = helper.LookupPassword(t_User.Password)
	return nil
}
