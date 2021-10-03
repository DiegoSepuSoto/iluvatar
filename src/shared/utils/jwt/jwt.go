package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/gommon/log"
	"os"
	"strings"
	"time"
)

type JWT interface {
	GenerateJWTForStudent(email string) (string, error)
	GenerateRefreshToken(email string) (string, error)
}

type TokenData map[string]string

type tokenGenerator struct {}

type CustomClaims struct {
	Data TokenData `json:"data"`
	jwt.StandardClaims
}

var tokenGeneratorValue *tokenGenerator

func Token() *tokenGenerator {
	if tokenGeneratorValue == nil {
		tokenGeneratorValue = &tokenGenerator{}
	}
	return tokenGeneratorValue
}

func (*tokenGenerator) getJWTSeed() []byte {
	return []byte(os.Getenv("JWT_TOKEN_SEED"))
}

func (t *tokenGenerator) GenerateJWTForStudent(email string) (string, error) {
	return t.GenerateJWTForStudentAndDuration(email, TokenData{"email": email}, time.Hour)
}

func (t *tokenGenerator) GenerateRefreshToken(email string) (string, error) {
	const threeMonthsInHours = 2190 * time.Hour
	return t.GenerateJWTForStudentAndDuration(email, TokenData{"email": email}, threeMonthsInHours)
}

func (t *tokenGenerator) GenerateJWTForStudentAndDuration(audience string, data TokenData, duration time.Duration) (string, error) {
	now := time.Now()
	StandardClaims := jwt.StandardClaims{
		Issuer:    "ILUVATAR_API",
		Audience:  audience,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: now.Add(duration).Unix(),
	}
	customClaims := CustomClaims{
		Data:           data,
		StandardClaims: StandardClaims,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, customClaims)
	seed := t.getJWTSeed()
	return token.SignedString(seed)
}

func (t *tokenGenerator) GetDataFromToken(token, key string) (string, error) {
	if token == "" {
		return "", errors.New("empty token")
	}
	claims, err := t.GetClaimsFromToken(token)
	if err != nil {
		log.Error("error getting claims from token")
		return "", errors.New("error getting claims from token")
	}
	claimsMap := claimsToMap(claims)
	if claimsMap != nil {
		if data, ok := claimsMap[key]; ok {
			return data, nil
		}
	}
	return "", errors.New("not found information in token")
}

func (t *tokenGenerator) Validate(tokenStr string) error {
	if tokenStr == "" {
		return errors.New("empty token")
	}
	token, err := t.getParsedToken(tokenStr)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	if !token.Valid {
		log.Error("invalid link Token")
		return errors.New("invalid link token")
	}
	return nil
}

func (t *tokenGenerator) getParsedToken(tokenStr string) (*jwt.Token, error) {
	certificateBytes := t.getJWTSeed()
	if strings.Contains(tokenStr, "Bearer ") {
		tokenStr = strings.Split(tokenStr, "Bearer ")[1]
	}
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return certificateBytes, nil
	})
	return token, err
}

func claimsToMap(claims *CustomClaims) map[string]string {
	data := make(map[string]string)
	for k, v := range claims.Data {
		data[k] = v
	}
	data["aud"] = claims.Audience
	data["iss"] = claims.Issuer
	return data
}

func (t *tokenGenerator) GetClaimsFromToken(tokenStr string) (*CustomClaims, error) {
	certificateBytes := t.getJWTSeed()
	if strings.Contains(tokenStr, "Bearer ") {
		tokenStr = strings.Split(tokenStr, "Bearer ")[1]
	}
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return certificateBytes, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	log.Error("invalid token")
	return nil, errors.New("invalid token")
}