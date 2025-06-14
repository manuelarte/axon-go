package models

import (
	"fmt"
	axongo "github.com/manuelarte/axon-go"
	"goapp/constants"
)

var _ axongo.Payloadable = new(User)

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func (u User) GetType() string {
	return fmt.Sprintf("%s.%s", constants.PackagePrefix, "models.User")
}
