package app

import (
	"adp-backend/models"
	"context"

	"github.com/labstack/gommon/log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBConfig() *mongo.Database {
	clientOperations := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	client, err := mongo.Connect(context.TODO(), clientOperations)
	if err != nil {
		log.Info(err)
		//panic(err)
	}
	return client.Database(models.DB_NAME)
}

func ConfigRouter() *echo.Echo {
	router := echo.New()
	router.HideBanner = false
	router.HidePort = false
	router.Debug = true

	// Middleware
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	//CORS
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	return router
}
