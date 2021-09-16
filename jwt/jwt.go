package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

var (
	invalid      = errors.New("令牌无效")
	expired      = errors.New("令牌过期")
	nonactivated = errors.New("令牌尚未激活")
)

type (
	jwtToken struct {
		SigningKey []byte
	}

	// serialize 序列化方法
	serialize func(token *jwt.Token) error
)

func New(key string) *jwtToken {
	return &jwtToken{SigningKey: []byte(key)}
}

//Get 获取Token
func (j *jwtToken) Get(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

//Parse 解析Token
//  tokenString 令牌值
//  claims 实现了 jwt.Claims的对象
//  serialize 实际将token中包含的对象序列化为想要的对象的方法
func (j *jwtToken) Parse(tokenString string, claims jwt.Claims, serialize serialize) error {
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return invalid
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return expired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nonactivated
			}
		}
		return invalid
	}

	if err := serialize(token); err != nil && token.Valid {
		return invalid
	}
	return invalid
}
