package http

import (
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JwtClaim struct {
	Id string `json:"id"`
	jwt.StandardClaims
}

func (c *JwtClaim) ToSignJwtString() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	// Generate encoded token and send it as response.
	return token.SignedString([]byte("secret"))
}

func ContextToClaim(c echo.Context) *JwtClaim {
	jwtClaims := &JwtClaim{}
	user := c.Get("user").(*jwt.Token)
	tmp, _ := json.Marshal(user.Claims)
	_ = json.Unmarshal(tmp, &jwtClaims)
	//claims := user.Claims.(*JwtClaim)
	//return claims

	return jwtClaims
}

func AuthMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("secret"),
	})
}