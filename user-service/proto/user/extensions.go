package user

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

// BeforeCreate - hook to gorm and create UUID string for User id instead of using autoincrement
func (model *User) BeforeCreate(scope *gorm.Scope) error {
	uuidv4 := uuid.NewV4()
	return scope.SetColumn("Id", uuidv4.String())
}
