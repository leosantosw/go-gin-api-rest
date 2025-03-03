package models

type Student struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	CPF  string `json:"cpf"`
	RG   string `json:"rg"`
}
