package api

import (
	"context"
	"fmt"
	axongo "github.com/manuelarte/axon-go"
	"goapp/constants"

	"goapp/models"
	"goapp/repositories"
)

var _ axongo.Payloadable = new(UserRead)

type UserRead struct {
	ID      int    `json:"id" structs:"id"`
	Name    string `json:"name" structs:"name"`
	Surname string `json:"surname" structs:"surname"`
}

func NewUserReadFromUser(u *models.User) *UserRead {
	if u == nil {
		return nil
	}
	return &UserRead{
		ID:      u.ID,
		Name:    u.Name,
		Surname: u.Surname,
	}
}

func (g *UserRead) GetType() string {
	return constants.UserReadType
}

type UserReadProjection struct {
	Repository repositories.Repository
}

func (p *UserReadProjection) GetUserByID(ctx context.Context, q GetUserByIDQuery) (*UserRead, error) {
	user, err := p.Repository.FindByID(ctx, q.ID)
	if err != nil {
		return nil, fmt.Errorf("error finding user by id: %w", err)
	}
	return NewUserReadFromUser(user), nil
}
