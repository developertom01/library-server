package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"github.com/developertom01/library-server/config"
	"github.com/dgryski/trifles/uuid"
	"github.com/golang-jwt/jwt/v5"
)

type JWTClaim struct {
	ID        uint
	UUID      string
	FirstName string
	LastName  string
	Email     string
}

type Token struct {
	AccessToken  string
	RefreshToken string
}

func SignToken(claim JWTClaim) (*Token, error) {
	userString, err := json.Marshal(claim)
	if err != nil {
		return nil, err
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userString,
		"exp": time.Now().Add(time.Minute * 30).Unix(),
		"iat": time.Now().Unix(),
		"jti": uuid.UUIDv4(),
	})
	accessTokenString, err := accessToken.SignedString([]byte(config.APP_SECRET))
	if err != nil {
		return nil, err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userString,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"iat": time.Now().Unix(),
		"jti": uuid.UUIDv4(),
	})
	refreshToken_string, err := refreshToken.SignedString([]byte(config.APP_SECRET))
	if err != nil {
		return nil, err
	}

	return &Token{
		AccessToken:  accessTokenString,
		RefreshToken: refreshToken_string,
	}, nil
}

func ValidateToken(tokenString string) (*JWTClaim, error) {
	var claim *JWTClaim
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		hmacSampleSecret := []byte(config.APP_SECRET)
		return hmacSampleSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		claim, err = extractJWTClaim(claims["sub"].(string))
		if err != nil {
			return nil, err
		}
	}

	return claim, nil
}

func extractJWTClaim(subClaim string) (*JWTClaim, error) {
	var claim JWTClaim

	decodedBytes, err := base64.StdEncoding.DecodeString(subClaim)
	err = json.Unmarshal(decodedBytes, &claim)
	if err != nil {
		return nil, err
	}

	return &claim, nil
}
