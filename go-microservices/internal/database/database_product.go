package database

import (
	"context"
	"errors"

	"github.com/dvdsantana/learninggo/go-microservices/internal/dberrors"
	"github.com/dvdsantana/learninggo/go-microservices/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (c Client) GetAllProducts(ctx context.Context, vendorID string) ([]models.Product, error) {
	var products []models.Product
	result := c.DB.WithContext(ctx).
		Where(models.Product{VendorID: vendorID}).
		Find(&products)

	return products, result.Error
}

func (c Client) AddProduct(ctx context.Context, product *models.Product) (*models.Product, error) {
	product.ProductID = uuid.New().String()
	result := c.DB.WithContext(ctx).
		Create(&product)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error
	}
	return product, nil
}

func (c Client) GetProductById(ctx context.Context, ID string) (*models.Product, error) {
	product := &models.Product{}
	result := c.DB.WithContext(ctx).
		Where(&models.Product{ProductID: ID}).
		First(&product)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &dberrors.NotFoundError{Entity: "product", ID: ID}
		}
		return nil, result.Error
	}
	return product, nil
}

func (c Client) UpdateProduct(ctx context.Context, product *models.Product) (*models.Product, error) {
	theProduct := &models.Product{}
	result := c.DB.WithContext(ctx).
		Model(theProduct).
		Clauses(clause.Returning{}).
		Where(&models.Product{ProductID: product.ProductID}).
		Updates(models.Product{
			Name:     product.Name,
			VendorID: product.VendorID,
			Price:    product.Price,
		})

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, &dberrors.NotFoundError{Entity: "product", ID: product.ProductID}
	}
	return theProduct, nil
}

func (c Client) DeleteProduct(ctx context.Context, ID string) error {
	product, error := c.GetProductById(ctx, ID)
	if error != nil {
		return &dberrors.NotFoundError{Entity: "product", ID: ID}
	}
	return c.DB.WithContext(ctx).Delete(product).Error
}
