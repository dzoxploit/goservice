package config

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	responCode := 500

	respon := BasicResponse{}

	respon.Success = false
	respon.Message = ""

	sendErrorResponse := func() {
		if !c.Response().Committed {
			if c.Request().Method == http.MethodHead {
				err = c.NoContent(responCode)
			} else {
				err = c.JSON(responCode, respon)
			}

			if err != nil {
				c.Logger().Error(err)
			}
		}
	}

	validationErrors, ok := err.(validator.ValidationErrors)

	if ok {
		responCode = 400

		respon.Message = "Validation errors"
		respon.Errors = map[string]string{}
		for _, ve := range validationErrors {
			respon.Errors[ve.Field()] = "Contains unexpected value"
		}

		sendErrorResponse()
		return
	}

	if he, ok := err.(*echo.HTTPError); ok {
		responCode = he.Code
		respon.Message = fmt.Sprintf("%v", he.Message)

		if he.Internal != nil {
			err = fmt.Errorf("%v, %v", err, he.Internal)

			respon.Message += " - ErrorMessageInternal: " + he.Internal.Error()
		}
	} else {
		respon.Message = http.StatusText(responCode)
	}
	//else if MainCfg.Development {
	//	resp.Message = err.Error()
	//}

	if responCode == 500 {
		// 4 KB stack.
		stack := make([]byte, 4<<10)
		length := runtime.Stack(stack, false)
		fmt.Printf("[RECOVER From Exception]: %v %s\n", err, stack[:length])
	}

	c.Logger().Debug(err)
	sendErrorResponse()
}