package models

import (
	"github.com/jinzhu/gorm"
)

type Project struct {
	gorm.Model
	Name string
	Desc string 
	Tags string
	RepoLink string
	ComChannel string
}

type Mentor struct{
	gorm.Model
	Name string
	Email string
	Access_token string
	Github_handle string
}