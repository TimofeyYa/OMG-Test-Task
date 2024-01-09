package models

type Task struct {
	Id        int
	CompanyId int
	Status    string
}

type TokenPair struct {
	User   string
	Bearer string
}
