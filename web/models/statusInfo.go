package models

import "gorm.io/gorm"

type StatusInfo struct {
	gorm.Model
	Component string `gorm:"unique"`
	Version   string
	Status    string
}
