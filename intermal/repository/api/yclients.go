package api

import "context"

type YClients struct {
	uri string
}

func NewYClients(uri string) *YClients {
	return &YClients{uri: uri}
}

func (y *YClients) GetAPIStaff(c context.Context, companyId int) (*any, error) {
	return nil, nil
}
