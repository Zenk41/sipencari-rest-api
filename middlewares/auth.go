package middlewares

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JWTCustomClaims struct {
	ID   string `json:"id"`
	Role string `json:"role"`
	jwt.StandardClaims
}

type ConfigJWT struct {
	SecretJWT      string
	ExpireDuration int
}

func (cj *ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JWTCustomClaims{},
		SigningKey: []byte(cj.SecretJWT),
	}
}

// GenerateToken perform generating token and refresh token with exp
func (cj *ConfigJWT) GenerateToken(userID string, role string) (string, error) {

	claims := JWTCustomClaims{
		userID,
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(cj.ExpireDuration))).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString([]byte(cj.SecretJWT))

	return t, nil
}

// DecodeTokenClaims perform decode or claiming token
func DecodeTokenClaims(c echo.Context) *JWTCustomClaims {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*JWTCustomClaims)

	return claims
}

// checkRole perform checking roles user
func checkUserRole(roles []string, hasRole string) bool {
	for _, role := range roles {
		if role == hasRole {
			return true
		}
	}
	return false
}

// AuthorizedUserAs perform authorized user with role
func AuthorizedUserAs(roles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := DecodeTokenClaims(c)
			if !checkUserRole(roles, claims.Role) {
				return echo.ErrForbidden
			}
			return next(c)
		}
	}
}
