package main

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.POST("/register", handleRegistration)
	e.Start(":8080")
}
func handleRegistration(c echo.Context) error {
	type RegisterData struct {
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
		Password  string `json:"password"`
		Email     string `json:"email"`
		ID        string `json:"id"`
	}
	var data RegisterData
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request data",
		})
	}
	db, err := sql.Open("mysql", "root:Handa@2019@tcp(127.0.0.1:3306)/hander.go")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "connection to the database has failed",
		})
	}
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO users (firstname, lastname, password, email, id) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Database preparation failed",
		})
	}
	defer stmt.Close()
	_, err = stmt.Exec(data.Firstname, data.Lastname, data.Password, data.Email, data.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Data insertion into the database failed",
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Registration successful",
	})
}
