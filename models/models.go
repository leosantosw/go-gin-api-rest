package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Id   string `json:"id"`
	Name string `json:"name"`
	CPF  string `json:"cpf"`
	RG   string `json:"rg"`
}
