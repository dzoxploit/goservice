package config

import (
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github,com/labstack/echo"
)

const (
	UserId = "User_Test_Beta"
	SecretKey = "secret"
	SigningMethod = "HS512"
)

type LoginForm struct {}