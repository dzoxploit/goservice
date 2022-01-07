package config

import (
	"Gone/app/models"

	"github.com/labstack/echo"
)

type AppContext struct {
	echo.Context
	User *models.User
}