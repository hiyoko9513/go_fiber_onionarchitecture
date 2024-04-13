package auth

import (
	"fmt"
	"time"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

const (
	UserCtxKey = "user"
)

type Claims struct {
	ID       string
	Password string
	Exp      int64
}

func NewClaims(ID string) *Claims {
	return &Claims{
		ID:  ID,
		Exp: getDefaultExpiration(),
	}
}

func (c Claims) toJWTClaims() *jwt.MapClaims {
	return &jwt.MapClaims{
		"id":  c.ID,
		"exp": c.Exp,
	}
}

func (c Claims) createToken() *jwt.Token {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, c.toJWTClaims())
}

func (c Claims) CreateTokenString() (string, error) {
	return c.createToken().SignedString([]byte(getSecretKey()))
}

// getSecretKey get secret key of jwt
// todo jwtのシークレットキーを.envから取得するようにする
func getSecretKey() string {
	return "secret"
}

// getDefaultExpiration get default expiration of jwt limit
// todo 時間は.envから取得する
func getDefaultExpiration() int64 {
	return time.Now().Add(time.Hour * 72).Unix()
}

// Auth middleware for protected routes
func Auth() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(getSecretKey())},
	})
}

// getJWT get jwt token from ctx
func getJWT(c *fiber.Ctx) *jwt.Token {
	return c.Locals(UserCtxKey).(*jwt.Token)
}

// GetClaimsFromCtx get claims from ctx
func GetClaimsFromCtx(c *fiber.Ctx) (*Claims, error) {
	claims := getJWT(c).Claims.(jwt.MapClaims)
	_, ok := claims["id"].(string)
	if !ok {
		return &Claims{}, fmt.Errorf("error getting claims id")
	}
	return &Claims{
		ID:  claims["id"].(string),
		Exp: int64(claims["exp"].(float64)),
	}, nil
}
