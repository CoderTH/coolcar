package token

import (
	"crypto/rsa"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type JWTVerifier struct {
	PublicKey *rsa.PublicKey

}
//JWTToken解析实现
func (v *JWTVerifier)Verify(token string)(string,error) {
	t, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		//签名验证
		return v.PublicKey, nil
	})
	if err != nil {
		return "", fmt.Errorf("cannot parse token: %v",err)
	}
	if !t.Valid{
		return "",fmt.Errorf("token not valid")
	}
	claims ,ok := t.Claims.(*jwt.StandardClaims)
	if !ok{
		return "",fmt.Errorf("token claim is not StandardClaims")

	}
	//数据校验：过期时间等
	err = claims.Valid()
	if err != nil {
		return "",fmt.Errorf("claim not valid")
	}
	return claims.Subject,nil
}
