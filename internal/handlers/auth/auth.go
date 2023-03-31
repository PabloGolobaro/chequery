package auth

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
	"time"
)

const (
	getToken        = "/auth"
	exampleUsername = "JohnSnow"
	examplePassword = "password"
)

type AuthDataRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthDataResponce struct {
	Token string `json:"token"`
}

type JwtCustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type authHandler struct {
	log       *zap.SugaredLogger
	jwtSecret string
}

func NewAuthHandler(log *zap.SugaredLogger, jwtSecret string) *authHandler {
	return &authHandler{log: log, jwtSecret: jwtSecret}
}

func (u *authHandler) Register(router *echo.Group) {
	router.Add(http.MethodPost, getToken, u.GenerateToken)
	router.Add(http.MethodGet, getToken, u.PassFormData)
}
func (u *authHandler) PassFormData(c echo.Context) error {
	return c.Render(http.StatusOK, "auth_form", nil)
}

func (u *authHandler) GenerateToken(c echo.Context) error {
	authData := AuthDataRequest{
		c.FormValue("username"),
		c.FormValue("password"),
	}

	if authData.Username == "" || authData.Password == "" {
		err := c.Bind(&authData)
		if err != nil {
			return echo.ErrBadRequest
		}
	}
	u.log.Info(authData.Username)
	u.log.Info(authData.Password)
	// Throws unauthorized error
	if authData.Username != exampleUsername || authData.Password != examplePassword {
		return echo.ErrUnauthorized
	}

	// Set custom claims
	claims := &JwtCustomClaims{
		authData.Username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(u.jwtSecret))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, AuthDataResponce{Token: t})
}
