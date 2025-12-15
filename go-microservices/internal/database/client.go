package database

import (
	"context"
	"fmt"
	"time"

	"github.com/dvdsantana/learninggo/go-microservices/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type DatabaseClient interface {
	Ready() bool
	GetAllCustomers(ctx context.Context, emailAddress string) ([]models.Customer, error)
	GetAllProducts(ctx context.Context, vendorID string) ([]models.Product, error)
	GetAllServices(ctx context.Context) ([]models.Service, error)
	GetAllVendors(ctx context.Context) ([]models.Vendor, error)
}

type Client struct {
	DB *gorm.DB
}

func NewDatabaseClient() (DatabaseClient, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		"localhost",
		"postgres",
		"postgres",
		"postgres",
		5432)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "wisdom.",
		},
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
		QueryFields: true,
	})

	if err != nil {
		return nil, err
	}

	return &Client{DB: db}, nil
}

func (c *Client) Ready() bool {
	var ready string
	c.DB.Raw("SELECT 1 as ready").Scan(&ready)
	return ready == "1"
}
