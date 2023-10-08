package utils

import (
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
	access_token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userString,
		"exp": time.Now().Add(time.Minute * 30),
		"iat": time.Now(),
		"jti": uuid.UUIDv4(),
	})
	access_token_string, err := access_token.SignedString(config.APP_SECRET)
	if err != nil {
		return nil, err
	}
	refresh_token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userString,
		"exp": time.Now().Add(time.Hour * 24),
		"iat": time.Now(),
		"jti": uuid.UUIDv4(),
	})
	refresh_token_string, err := refresh_token.SignedString(config.APP_SECRET)
	if err != nil {
		return nil, err
	}
	return &Token{
		AccessToken:  access_token_string,
		RefreshToken: refresh_token_string,
	}, nil
}

func ValidateToken(tokenString string) (*JWTClaim, error) {
	var claim JWTClaim

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		hmacSampleSecret := []byte(config.APP_SECRET)
		return hmacSampleSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		claim = claims["sub"].(JWTClaim)
	}

	return &claim, nil
}
