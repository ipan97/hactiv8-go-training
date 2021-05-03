package main

import (
	"github.com/ipan97/hactiv8-assigment2/actions"
	"github.com/ipan97/hactiv8-assigment2/config"
	"github.com/ipan97/hactiv8-assigment2/migrations"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db, err := config.PGConnect(&config.Database{
		Driver:            "postgres",
		Host:              "localhost",
		Port:              5432,
		User:              "postgres",
		Password:          "postgres",
		DatabaseName:      "orders_by",
		MaxIdleConnection: 10,
		MaxOpenConnection: 10,
	})
	if err != nil {
		log.Errorf("Error DB connect cause : %v", err)
	}

	// Up db migrations
	migrations.Up(db)

	r := echo.New()

	// Middleware
	r.Use(middleware.Logger())
	r.Use(middleware.Recover())

	// Routes
	orderHandler := actions.NewOrderHandler(db)
	r.POST("/orders", orderHandler.Create)
	r.GET("/orders", orderHandler.GetAll)
	r.PUT("/orders", orderHandler.Update)
	r.DELETE("/orders/:id", orderHandler.Delete)

	// Start server
	r.Logger.Fatal(r.Start(":8080"))
}
