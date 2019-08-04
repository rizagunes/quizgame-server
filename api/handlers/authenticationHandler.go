package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type JwtClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "ben" && password == "bildim" {
		// create jwt token
		token, err := createJwtToken(username)
		if err != nil {
			log.Println("Error: Creating JWT token", err)
			return c.String(http.StatusInternalServerError, "Whoops, look like something went wrong.")
		}

		return c.JSON(http.StatusOK, map[string]string{
			"message": "You were logged in!",
			"token":   token,
		})
	}

	return c.String(http.StatusUnauthorized, "You entered and incorrect username or password.")
}

func createJwtToken(username string) (string, error) {
	claims := JwtClaims{
		username,
		jwt.StandardClaims{
			Id:        "main_user_id",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	token, err := rawToken.SignedString([]byte("secretKey"))
	if err != nil {
		return "", err
	}

	return token, nil
}
