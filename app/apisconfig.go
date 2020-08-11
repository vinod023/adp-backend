package app

import (
	"adp-backend/jwtauth"
	"adp-backend/models"
	"context"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
)

type Config struct {
	DB     *mongo.Database
	Router *echo.Echo
}

func (config Config) ConfigAPIs() {
	config.Router.POST("/user", config.createUser)
	config.Router.POST("/login", config.login)
	config.Router.GET("/users", config.findUsers)
	config.Router.GET("/testuser", config.testUser)
}
func (config Config) testUser(c echo.Context) error {

	result := map[string]interface{}{
		"message": "Hello...!, it's working",
	}
	return c.JSON(http.StatusOK, result)
}

func (config Config) createUser(c echo.Context) error {

	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	insertResult, err := config.DB.Collection(models.TABLE_USER).InsertOne(context.TODO(), user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err)
	}

	result := map[string]interface{}{
		"message": insertResult.InsertedID,
	}
	return c.JSON(http.StatusOK, result)
}

func (config Config) login(c echo.Context) error {
	user := new(models.User)
	var err error
	if err = c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if len(user.UserID) == 0 || len(user.Password) == 0 {
		return c.JSON(http.StatusOK, "Invalid userId/password")
	}

	var dbUser models.User
	if err = config.DB.Collection(models.TABLE_USER).FindOne(context.TODO(), bson.M{"userid": user.UserID}).Decode(&dbUser); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	tokens, err := jwtauth.CreateToken(dbUser.UserID)
	fmt.Println("Data:", tokens, err)
	return c.JSON(http.StatusOK, "Login Success")
}

func (config Config) findUsers(c echo.Context) error {

	// cur, err := config.DB.Collection("employee").Find(context.TODO(), bson.D{{}}, options.Find())
	// if err != nil {
	// 	fmt.Print(err)
	// }

	// var employee []models.Employee
	// for cur.Next(context.TODO()) {
	// 	var elem models.Employee
	// 	err := cur.Decode(&elem)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}

	// 	employee = append(employee, elem)
	// }

	// if err := cur.Err(); err != nil {
	// 	fmt.Println(err)
	// }

	// cur.Close(context.TODO())

	return c.JSON(http.StatusOK, "Employees Data")
}
