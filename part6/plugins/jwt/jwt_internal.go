package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	// tokenExpiredDate app token过期日期 30天
	jwtExpiredDate = 3600 * 24 * 30 * time.Second
	// jwtSecretKey   = "XsX20LwAEeUQpCcXcmbzt0yI0x0sObUQ"
)

type CustomClaims struct {
	UserID string `json:"uid"`
	jwt.StandardClaims
}

// 生成 jwt
func (j *Jwt) Encode(userID string) (tk string, err error) {
	m, err := j.createTokenClaims(userID)
	if err != nil {
		return "", fmt.Errorf("[Encode] 创建token Claim 失败，err: %s", err)
	}
	ret := jwt.NewWithClaims(jwt.SigningMethodHS256, m)

	tk, err = ret.SignedString([]byte(j.SecretKey))
	if err != nil {
		return "", fmt.Errorf("[Encode] 创建token失败，err: %s", err)
	}

	return tk, nil
}

// 解析 jwt
func (j *Jwt) Decode(tk string) (cc *CustomClaims, err error) {
	t, err := jwt.ParseWithClaims(tk, &CustomClaims{}, secretFunc(j.SecretKey))
	if err != nil {
		return nil, err
	}
	if t != nil {
		// 解密转换类型并返回
		if claims, ok := t.Claims.(*CustomClaims); ok && t.Valid {
			return claims, nil
		}
	}

	return nil, err

}

func (j *Jwt) createTokenClaims(userID string) (m *CustomClaims, err error) {
	now := time.Now()
	m = &CustomClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			Audience:  "",
			ExpiresAt: now.Add(jwtExpiredDate).Unix(),
			Id:        userID,
			IssuedAt:  now.Unix(),
			Issuer:    "entere.github.io",
			NotBefore: now.Unix(),
			Subject:   userID,
		},
	}
	return
}

// secretFunc validates the secret format.
func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		// Make sure the `alg` is what we except.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret), nil
	}
}
