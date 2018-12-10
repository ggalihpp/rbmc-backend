package membership

import (
	"fmt"
	"net/http"

	"github.com/davecgh/go-spew/spew"

	"github.com/creasty/defaults"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/ggalihpp/rbmc-backend/database"
	"github.com/ggalihpp/rbmc-backend/primary"

	"github.com/labstack/echo"
)

// This layer contain handler for each route, keep it clean from any logic

// SetupHandler -
func SetupHandler(e *echo.Group) {
	e.POST("", registerUserEP)
	e.PUT("", updateMemberEP)
	e.GET("", getMembersEP)

	e.POST("/test", test)
}

func registerUserEP(c echo.Context) error {

	m := new(Member)

	if err := c.Bind(m); err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	m.Password, _ = hashString(m.Password)
	m.Username = getUsername(m.Email)

	if m.IsExist() {
		return echo.NewHTTPError(http.StatusExpectationFailed, "That username is exist")
	}

	if err := registerUser(m); err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	return c.JSON(http.StatusCreated, h{
		"result": "OK",
	})
}

func updateMemberEP(c echo.Context) error {
	// CHECK TOKEN //
	claims := c.Get("user").(*jwt.Token)
	user := claims.Claims.(*primary.JwtCustomClaims)
	/////////////////

	m := new(Member)

	if err := c.Bind(m); err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	m.Password, _ = hashString(m.Password)
	m.Username = getUsername(m.Email)

	if !m.IsExist() {
		return echo.NewHTTPError(http.StatusExpectationFailed, "That username is not exist")
	}

	switch {
	case user.IsAdmin:
		break
	case user.IsCoordinator && m.Territory == user.Territory:
		break
	case user.Username == m.Username:
		break
	default:
		return echo.ErrUnauthorized
	}

	if err := updateMember(m); err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	return c.JSON(http.StatusOK, h{
		"result": "OK",
	})
}

func getMembersEP(c echo.Context) error {
	// CHECK TOKEN //
	claims := c.Get("user").(*jwt.Token)
	user := claims.Claims.(*primary.JwtCustomClaims)
	/////////////////

	spew.Dump(user)

	username := c.QueryParam("username")
	territory := c.QueryParam("territory")

	var result interface{}
	var err error
	var sc int
	var qHelper database.QueryHelper

	if err = c.Bind(&qHelper); err != nil {
		fmt.Println("BERART GA MASUK SINI")
		return echo.NewHTTPError(500, err.Error())
	}

	if err := defaults.Set(&qHelper); err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	switch {
	case username != "":
		result, sc, err = GetMemberByUsername(username)
	case territory != "":
		if !user.IsAdmin {
			if user.IsCoordinator && user.Territory == territory {
				result, sc, err = getMembersByTerritory(territory, qHelper)
			} else {
				return echo.ErrUnauthorized
			}

		} else {
			result, sc, err = getMembersByTerritory(territory, qHelper)
		}
	default:
		if !user.IsAdmin {
			return echo.ErrUnauthorized
		}

		result, sc, err = getAllMembers(qHelper)
	}

	if err != nil {
		return echo.NewHTTPError(sc, err.Error())
	}

	return c.JSON(200, result)
}

func test(c echo.Context) error {
	m := new(Member)

	if err := c.Bind(m); err != nil {
		return err
	}

	m.Password, _ = hashString(m.Password)
	m.Username = getUsername(m.Email)

	token, _ := testController()

	return c.JSON(200, token)
}
