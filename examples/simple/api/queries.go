package api

import (
	"fmt"
	axongo "github.com/manuelarte/axon-go"
	"goapp/constants"
)

var _ axongo.Payloadable = new(GetUserByIDQuery)

type GetUserByIDQuery struct {
	ID int `json:"id" binding:"required"`
}

func (g GetUserByIDQuery) GetType() string {
	return fmt.Sprintf("%s.%s", constants.PackagePrefix, "api.GetUserByIDQuery")
}
