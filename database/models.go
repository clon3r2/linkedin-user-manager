package database

import (
	"gorm.io/gorm"
	"log"
)

type LinkedinUser struct {
	gorm.Model
	Name       string
	ListTypeID int
	ListType   ListType
}

func (obj *LinkedinUser) all() (allUsers []LinkedinUser) {
	DBConn.Find(&allUsers)
	return
}

type ListType struct {
	gorm.Model
	TypeName string
}

func (obj *ListType) all() (allTypes []ListType) {
	DBConn.Find(allTypes)
	return
}

func Migrate() {
	err := DBConn.AutoMigrate(&LinkedinUser{}, &ListType{})
	if err != nil {
		log.Fatal(err)
	}
}
