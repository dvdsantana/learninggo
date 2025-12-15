package database

import (
	"context"

	"github.com/dvdsantana/learninggo/go-microservices/internal/models"
)

func (c Client) GetAllServices(ctx context.Context) ([]models.Service, error) {
	var services []models.Service
	result := c.DB.WithContext(ctx).
		Find(&services)

	return services, result.Error
}