package jwts

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtToken struct {
	AccessToken  string        `json:"access_token"`
	RefreshToken string        `json:"refresh_token"`
	AccessExp    time.Duration `json:"access_exp"`
	RefreshExp   time.Duration `json:"refresh_exp"`
}

func CreateToken(val, accessSecret, refreshSecret string, accessExp, refreshExp time.Duration, ip string) *JwtToken {
	aExp := time.Now().Add(accessExp).Unix()
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"token": val,
		"exp":   aExp,
		"ip":    ip,
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

func ParseToken(tokenString, secret, ip string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// 安全断言 token 字段
		tokenVal, ok := claims["token"]
		if !ok || tokenVal == nil {
			return "", fmt.Errorf("token field missing")
		}
		val, ok := tokenVal.(string)
		if !ok {
			return "", fmt.Errorf("token field type error")
		}
		// 安全断言 exp 字段
		expVal, ok := claims["exp"]
		if !ok || expVal == nil {
			return "", fmt.Errorf("exp field missing")
		}
		expFloat, ok := expVal.(float64)
		if !ok {
			return "", fmt.Errorf("exp field type error")
		}
		if time.Now().Unix() > int64(expFloat) {
			return "", fmt.Errorf("token expired")
		}
		// 安全断言 ip 字段
		ipVal, ok := claims["ip"]
		if !ok || ipVal == nil {
			return "", fmt.Errorf("ip field missing")
		}
		ipStr, ok := ipVal.(string)
		if !ok {
			return "", fmt.Errorf("ip field type error")
		}
		if ipStr != ip {
			return "", fmt.Errorf("ip error")
		}
		return val, nil
	} else {
		fmt.Println(err)
		return "", err
	}
}
