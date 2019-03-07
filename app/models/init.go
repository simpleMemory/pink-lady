// Package models save database models
// define your database models in the package
package models

import (
	"github.com/axiaoxin/pink-lady/app/utils"
)

// BaseModel you should define you model with BaseModel
type BaseModel struct {
	ID        string          `gorm:"primary_key" json:"id"`
	CreatedAt utils.JSONTime  `json:"createdAt"`
	UpdatedAt utils.JSONTime  `json:"updatedAt"`
	DeletedAt *utils.JSONTime `sql:"index" json:"-"`
	DelFlag   int             `gorm:"column:del_flag"`
	UsedFlag  int             `gorm:"column:used_flag"`
}

// MigrationModels save models like &MODEL{} for auto migrate when server starting
// when you write your model you can append to it
var MigrationModels = []interface{}{}

// Migrate run AutoMigrate to create database tables which in MigrationModels
// running after InitGormDB
func Migrate() error {
	return utils.DB.AutoMigrate(MigrationModels...).Error
}
