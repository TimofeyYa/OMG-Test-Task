package models

type Task struct {
	Id        int
	CompanyId int
	Status    string
}

type Status struct {
	Name   string
	IsDone bool
}

type TokenPair struct {
	User   string
	Bearer string
}
