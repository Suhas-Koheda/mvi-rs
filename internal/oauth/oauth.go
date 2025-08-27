package oauth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/suhas-koheda/mvi-rs/pkg/env"
)

var (
	key []byte
	t   *jwt.Token
	s   string
	err error
)

func GetJWT(email string, role string) (string, error) {
	key = []byte(loadenv.HashSecret)
	t = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"role":   role,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
		"issued": time.Now().Unix(),
	})
	s, err = t.SignedString(key)
	if err != nil {
		return "", err
	}
	return s, nil
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return key, nil
}

func ValidateJWT(jwtToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(jwtToken,
		keyFunc,
		jwt.WithValidMethods([]string{jwt.SigningMethodES256.Name}))
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}

func ExtractClaims(token string) (jwt.MapClaims, error) {
	parsedToken, err := ValidateJWT(token)
	if err != nil {
		return nil, err
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token claims")
}

func GetUserFromToken(tokenString string) (string, string, error) {
	claims, err := ExtractClaims(tokenString)
	if err != nil {
		return "", "", err
	}

	email, ok := claims["email"].(string)
	if !ok {
		return "", "", fmt.Errorf("user_id claim not found or invalid")
	}

	role, ok := claims["role"].(string)
	if !ok {
		return "", "", fmt.Errorf("role claim not found or invalid")
	}

	return email, role, nil
}
