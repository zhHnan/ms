package jwts

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JwtToken struct {
	AccessToken  string        `json:"access_token"`
	RefreshToken string        `json:"refresh_token"`
	AccessExp    time.Duration `json:"access_exp"`
	RefreshExp   time.Duration `json:"refresh_exp"`
}

func CreateToken(val, accessSecret, refreshSecret string, accessExp, refreshExp time.Duration) *JwtToken {
	aExp := time.Now().Add(accessExp).Unix()
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"token": val,
		"exp":   aExp,
	})
	aToken, _ := accessToken.SignedString([]byte(accessSecret))
	rExp := time.Now().Add(refreshExp).Unix()
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"token": val,
		"exp":   rExp,
	})
	rToken, _ := refreshToken.SignedString([]byte(refreshSecret))
	return &JwtToken{
		AccessToken:  aToken,
		RefreshToken: rToken,
		AccessExp:    accessExp,
		RefreshExp:   refreshExp,
	}
}

func ParseToken(tokenString, secret string) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Printf("error: %v", claims)
	} else {
		fmt.Println(err)
	}
}
