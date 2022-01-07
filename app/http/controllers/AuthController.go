package controllers

import (
	"Gone/config"
	"Gone/database"
	"fmt"
	"go_service/app/models"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

type UserForm struct {
	Id    string `json:"id"`
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

func MainJwt(c echo.Context) error {
	//reference data
	appContext := c.(*config.AppContext)
	fmt.Println(appContext.User.Name)
	return c.String(http.StatusOK, "SUCCESS: you are on the top secret jwt page aaaaa! : "+appContext.User.Name)
}

func Welcome(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func GetUser(c echo.Context) error {
	//	id := c.Param("id")
	//	team := c.QueryParam("team")
	//	member := c.QueryParam("member")
	// return c.String(http.StatusOK, "id: "+id+", team: "+team+", member: "+member)
	dbEngine := database.ConnectApi()
	strId := "12342423423143"
	user := &models.User{}
	dbEngine.SQL("select * from tbl_user where id = ? ", strId).Get(user)
	userForm := UserForm{
		Id:    user.Id_user,
		Name:  user.Name,
		Email: user.Email,
		
	}
	return c.JSON(http.StatusOK, userForm)
}

func SaveUserByForm(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name: "+name+", email: "+email)
}

func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "ID: "+id+", successfully updated.")
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusNoContent, id)
}

func SaveUserByJson(c echo.Context) error {
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}

func FileUpload(c echo.Context) error {
	avatar, err := c.FormFile("avatar")
	if err != nil {
		return err
	}

	// Source
	src, err := avatar.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create("upload-files/" + avatar.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, "Your file `"+avatar.Filename+"` successfully uploaded")
}