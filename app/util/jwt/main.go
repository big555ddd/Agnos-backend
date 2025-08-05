package jwt

import (
	"app/internal/logger"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

const (
	VAL_USER  = "AUTH_USER"
	VAL_TOKEN = "AUTH_TOKEN"
	VAL_AGENT = "AUTH_AGENT"
)

type ClaimData struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Hospital string `json:"hospital"`
}

type Claims struct {
	Data ClaimData `json:"data"`
	Uuid string    `json:"uuid"`
	jwt.RegisteredClaims
}

func CreateToken(claims ClaimData) (string, *Claims, error) {

	now := time.Now()
	id := uuid.New().String()
	duration := viper.GetInt64("JWT_DURATION")
	claimsData := Claims{
		Data: claims,
		Uuid: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(duration) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claimsData)
	secret := []byte(viper.GetString("JWT_SECRET"))

	tokenString, err := token.SignedString(secret)
	if err != nil {
		logger.Errf("%s", err.Error())
		return "", nil, err
	}
	return tokenString, &claimsData, nil
}

func Verify(rawToken string) (*Claims, bool, error) {
	token, err := jwt.ParseWithClaims(rawToken, &Claims{}, getSecret)
	if err != nil {
		return nil, false, err
	}

	claims, ok := token.Claims.(*Claims)
	return claims, token.Valid && ok, nil
}

func GetClaims(c *gin.Context) (*ClaimData, error) {
	val, exists := c.Get(VAL_USER)
	if !exists {
		return nil, errors.New("claims doesn't exists")
	}
	data := val.(*ClaimData)
	return data, nil
}

func getSecret(token *jwt.Token) (interface{}, error) {

	return []byte(viper.GetString("JWT_SECRET")), nil
}

func GenerateExpires() time.Time {
	now := time.Now()
	duration := viper.GetInt64("JWT_DURATION")
	return now.Add(time.Duration(duration) * time.Hour)
}
