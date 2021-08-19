package jwt

import (
	"errors"
	"github.com/ccb1900/tinyweb/config"
	"github.com/golang-jwt/jwt"
	"time"
)

type MyClaim struct {
	ID string
	UserName string
	jwt.StandardClaims
}

func (c MyClaim) Valid() error  {
	return nil
}
func ParseJwtToken(token string) (*MyClaim,error)  {
	t, err := jwt.ParseWithClaims(token, &MyClaim{}, func(token *jwt.Token) (interface{}, error) {
		return config.Get("jwt.secret"),nil
	})
	if err != nil {
		return nil,err
	}

	if claims, ok := t.Claims.(*MyClaim); ok && t.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
func CreateJwtToken(id,username string) (string,error)  {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(config.GetInt("jwt.expire"))*time.Second)
	issuer := config.Get("app.name")
	claims := MyClaim{
		ID:       id,
		UserName: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    issuer,
			Id: "1000",
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(config.Get("jwt.secret"))
	return token, err
}
