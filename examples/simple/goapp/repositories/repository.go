package repositories

import (
	"context"

	"goapp/models"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{db: db}
}

func (r Repository) FindByID(ctx context.Context, id int) (*models.User, error) {
	var user models.User
	return &user, r.db.WithContext(ctx).First(&user, id).Error
}
