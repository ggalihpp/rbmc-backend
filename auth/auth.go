package auth

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/ggalihpp/rbmc-backend/membership"
	"github.com/ggalihpp/rbmc-backend/primary"
	"github.com/labstack/echo"
)

// SetupHandler -
func SetupHandler(e *echo.Group) {
	e.POST("/login", login)
}

func login(c echo.Context) error {
	var m membership.Member

	if err := c.Bind(&m); err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	if m.IsExist() {
		if m.VerifyPassword() {
			member, sc, err := membership.GetMemberByUsername(m.Username)
			if err != nil {
				return echo.NewHTTPError(sc, err.Error())
			}

			// Set custom claims
			claims := &primary.JwtCustomClaims{
				member.Username,
				member.Email,
				member.IsAdmin,
				member.IsCoordinator,
				member.Territory,
				jwt.StandardClaims{
					ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
				},
			}

			// Create token with claims
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

			// Generate encoded token and send it as response.
			t, err := token.SignedString([]byte("secret"))
			if err != nil {
				return err
			}
			return c.JSON(http.StatusOK, echo.Map{
				"token": t,
			})
		}

		return echo.NewHTTPError(417, "Password Incorect")
	}

	return echo.NewHTTPError(404, "User not found")

}
