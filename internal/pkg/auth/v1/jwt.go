package auth

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"hiyoko-fiber/internal/shared"
	"hiyoko-fiber/utils"

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
func getSecretKey() string {
	return utils.Env("JWT_SECRET_KEY").GetString()
}

// getDefaultExpiration get default expiration of jwt limit
func getDefaultExpiration() int64 {
	duration := utils.Env("JWT_EXP").GetDuration()
	return time.Now().Add(duration).Unix()
}

// Auth middleware for protected routes
func Auth() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(getSecretKey())},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return shared.ResponseUnauthorized(c)
		},
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

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func GenerateRandomBase64String(s int) (string, error) {
	b, err := generateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}
