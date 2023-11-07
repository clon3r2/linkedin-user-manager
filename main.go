package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type LinkedinUser struct {
	gorm.Model

	name       string `gorm:"size:255;not null"`
	ListTypeID int
	ListType   ListType
}

type ListType struct {
	gorm.Model

	TypeName string `gorm:"size:255;not null"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("error opening db connection: %v", err))
	}

	err = db.AutoMigrate(&ListType{}, &LinkedinUser{})
	if err != nil {
		panic(fmt.Sprintf("error migrating models: %v", err))
	}
	db.Create(&ListType{
		TypeName: "mamad",
	})
	s := db.Where("type_name = ?", "mamad").Find(&ListType{})
	fmt.Printf("\n\nfirst == %+v", s)
	//db.
	//db.Create(&LinkedinUser{
	//	name: "gholi",
	//
	//	ListType: ListType{},
	//})
}
