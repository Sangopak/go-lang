package model

type Todo struct {
	Id int `json:"id"`
	Detail string `json:"message"`
	Finished bool `json:"finished"`
}