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
	AddCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error)
	GetCustomerById(ctx context.Context, ID string) (*models.Customer, error)
	UpdateCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error)

	GetAllProducts(ctx context.Context, vendorID string) ([]models.Product, error)
	AddProduct(ctx context.Context, product *models.Product) (*models.Product, error)
	GetProductById(ctx context.Context, ID string) (*models.Product, error)
	UpdateProduct(ctx context.Context, product *models.Product) (*models.Product, error)

	GetAllServices(ctx context.Context) ([]models.Service, error)
	AddService(ctx context.Context, service *models.Service) (*models.Service, error)
	GetServiceById(ctx context.Context, ID string) (*models.Service, error)
	UpdateServicePrice(ctx context.Context, service *models.Service) (*models.Service, error)

	GetAllVendors(ctx context.Context) ([]models.Vendor, error)
	AddVendor(ctx context.Context, vendor *models.Vendor) (*models.Vendor, error)
	GetVendorById(ctx context.Context, ID string) (*models.Vendor, error)
	UpdateVendor(ctx context.Context, vendor *models.Vendor) (*models.Vendor, error)
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
