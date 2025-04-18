package jwts

import (
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
