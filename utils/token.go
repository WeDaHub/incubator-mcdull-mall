package utils

import (
	"errors"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const SecretKey = "123456"

type Payload struct {
	jwt.StandardClaims
	Uid int `json:"uid"`
}

// CreateToken 创建Token
func CreateToken(uid int, exp int) (string, error) {
	claims := &Payload{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Add(time.Second * time.Duration(exp)).Unix()),
			Issuer:    "dazuo",
			NotBefore: time.Now().Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Uid: uid,
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SecretKey))
	if err != nil {
		log.Printf("call jwt NewWithClaims failed, err:%v", err)
		return "", errors.New("token无效")
	}
	return token, nil
}

// ValidateToken 验证Token
func ValidateToken(tokenStr string) bool {
	token, err := jwt.ParseWithClaims(tokenStr, &Payload{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		log.Printf("call jwt ParseWithClaims failed, err:%v", err)
		return false
	}
	return token.Valid
}

// ParseToken 解析Token
func ParseToken(tokenStr string) (*Payload, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Payload{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		log.Printf("call jwt ParseWithClaims failed, err:%v", err)
		return nil, errors.New("token解析失败")
	}
	if claims, ok := token.Claims.(*Payload); ok && token.Valid {
		log.Printf("Uid = %v, ExpiresAt = %v\n", claims.Uid, claims.StandardClaims.ExpiresAt)
		return claims, nil
	} else {
		return nil, errors.New("token无效")
	}
}
