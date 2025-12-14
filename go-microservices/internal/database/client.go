package database

import (
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm/schema"
)

type DatabaseClient interface {
	Ready() bool	
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
			SingularTable: true,
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
