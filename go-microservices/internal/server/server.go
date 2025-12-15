package server

import (
	"log"
	"net/http"

	"github.com/dvdsantana/learninggo/go-microservices/internal/database"
	"github.com/dvdsantana/learninggo/go-microservices/internal/models"
	"github.com/labstack/echo/v4"
)

type Server interface {
	Start() error
	Readiness(ctx echo.Context) error
	Liveness(ctx echo.Context) error

	GetAllCustomers(ctx echo.Context) error
	AddCustomer(ctx echo.Context) error
	GetCustomerById(ctx echo.Context) error
	UpdateCustomer(ctx echo.Context) error

	GetAllProducts(ctx echo.Context) error
	AddProduct(ctx echo.Context) error
	GetProductById(ctx echo.Context) error
	UpdateProduct(ctx echo.Context) error

	GetAllServices(ctx echo.Context) error
	AddService(ctx echo.Context) error
	GetServiceById(ctx echo.Context) error
	UpdateService(ctx echo.Context) error

	GetAllVendors(ctx echo.Context) error
	AddVendor(ctx echo.Context) error
	GetVendorById(ctx echo.Context) error
	UpdateVendor(ctx echo.Context) error
}

type EchoServer struct {
	echo *echo.Echo
	DB   database.DatabaseClient
}

func NewEchoServer(db database.DatabaseClient) Server {
	server := &EchoServer{
		echo: echo.New(),
		DB:   db,
	}
	server.registerRoutes()
	return server
}

func (s *EchoServer) Start() error {
	if err := s.echo.Start(":8080"); err != nil && err != http.ErrServerClosed {
		log.Fatal("failed to start server:", err)
		return err
	}

	return nil
}

func (s *EchoServer) registerRoutes() {
	s.echo.GET("/readiness", s.Readiness)
	s.echo.GET("/liveness", s.Liveness)

	customerGroup := s.echo.Group("/customers")
	customerGroup.GET("", s.GetAllCustomers)
	customerGroup.POST("", s.AddCustomer)
	customerGroup.GET("/:id", s.GetCustomerById)
	customerGroup.PUT("/:id", s.UpdateCustomer)

	productGroup := s.echo.Group("/products")
	productGroup.GET("", s.GetAllProducts)
	productGroup.POST("", s.AddProduct)
	productGroup.GET("/:id", s.GetProductById)
	productGroup.PUT("/:id", s.UpdateProduct)

	serviceGroup := s.echo.Group("/services")
	serviceGroup.GET("", s.GetAllServices)
	serviceGroup.POST("", s.AddService)
	serviceGroup.GET("/:id", s.GetServiceById)
	serviceGroup.PUT("/:id", s.UpdateService)

	vendorGroup := s.echo.Group("/vendors")
	vendorGroup.GET("", s.GetAllVendors)
	vendorGroup.POST("", s.AddVendor)
	vendorGroup.GET("/:id", s.GetVendorById)
	vendorGroup.PATCH("/:id", s.UpdateVendor)

}

func (s *EchoServer) Readiness(ctx echo.Context) error {
	if s.DB.Ready() {
		return ctx.JSON(http.StatusOK, models.Health{Status: "OK"})
	}
	return ctx.JSON(http.StatusServiceUnavailable, models.Health{Status: "NOT OK"})
}

func (s *EchoServer) Liveness(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.Health{Status: "OK"})
}
