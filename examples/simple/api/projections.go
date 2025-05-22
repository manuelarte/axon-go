package api

import (
	"context"
	"fmt"

	"goapp/models"
	"goapp/repositories"
)

type UserRead struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
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

func (g *UserRead) GetPackageName() string {
	return "org.github.axonserver.connector.go.example.kotlinapp.api.UserRead"
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
