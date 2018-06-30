package example

import (
	"net/http"

	"github.com/labstack/echo"
)

// This layer contain handler for each route, keep it clean from any logic

// SetupHandler -
func SetupHandler(e *echo.Group) {
	e.GET("", getExample)
}

func getExample(c echo.Context) error {
	input := c.QueryParam("input")

	result, err := logicExample(input)
	if err != nil {
		return c.JSON(http.StatusExpectationFailed, err.Error())
	}

	updateSomething()

	return c.JSON(http.StatusOK, h{
		"result": result,
	})
}
