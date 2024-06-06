package models

import "gorm.io/gorm"

type EnvInfo struct {
	gorm.Model
	SvcName string
	Infi    string
}
