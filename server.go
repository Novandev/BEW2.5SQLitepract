package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"

	"github.com/labstack/echo"
)

type Person struct {
	gorm.Model
	Name string
	Age uint
	Occupation string
}


func main() {
	db, err := gorm.Open("sqlite3", "person.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&Person{})

	e := echo.New()
	e.GET("/", func(c echo.Context) error {

		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":4000"))

	// Routes
	e.POST("/", createUser)
	e.GET("/person/:id", getUser)
	e.PUT("/person/:id", updateUser)
	e.DELETE("/person/:id", deleteUser)
}

