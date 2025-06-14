package api

import (
	axongo "github.com/manuelarte/axon-go"
	"goapp/constants"
)

var _ axongo.Payloadable = new(GetUserByIDQuery)

type GetUserByIDQuery struct {
	ID int `json:"id" binding:"required"`
}

func (g GetUserByIDQuery) GetType() string {
	return constants.GetUserByIDQueryType
}
