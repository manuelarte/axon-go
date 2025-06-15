package api

import (
	"goapp/constants"

	axongo "github.com/manuelarte/axon-go"
)

var _ axongo.Payloadable = new(GetUserByIDQuery)

type GetUserByIDQuery struct {
	ID int `json:"id" binding:"required"`
}

func (g GetUserByIDQuery) GetType() string {
	return constants.GetUserByIDQueryType
}
