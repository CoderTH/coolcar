package token

import (
	"github.com/dgrijalva/jwt-go"
	"testing"
)

const publicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAqcxPYZvp+CziUTi6+cCU
j7stduQbJ6hRHvbY4bZ77F8/ZmshPs6AAdMAK1au0B2pXNyYLnNm0HIYZM7YZv83
M98Qcy3yBTddPWbZfgmUVqAvnFj7PNuQmLVHyY06x2tQTT/5946GiPLwWjfQQHu3
MLZvCECn7E4V5TwKC9J/onuy92q4jX6E3lx3UGQ8hRnwDCLRKPw6N23lL+jcIqDg
qi51k2h6hGgBEN0FudCUerPz4kLW/zezLmYrmNbmUhGQ6poSTn1peliNCdQ1sC0i
RqgKlcBwcv/SU+eMMwNNgTy/pKT4xIJg8p+QXn9JdMpNMguz9uKU+CZr8KhB0iuq
rQIDAQAB
-----END PUBLIC KEY-----`




func TestJWTVerifier_Verify(t *testing.T) {
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
	if err != nil {
		t.Fatalf("cannot parse public key %v",err)
	}
	v :=JWTVerifier{
		PublicKey:pubKey,
	}
	tkn:="eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTYzOTc5MTUsImlhdCI6MTYxNjM5MDcxNSwiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiNjA1NWFmOGM4Mjc2ZTUzNzlkNGRhYjQxIn0.JOOWjQ6Jj5oYRXs0LsykrexLQCb65-K4L2yYvoxGjuyOr4nhNDnZpJclf5XfhIPA0EOJbOgzrxEwxp811AcT-VVkad2pVv16dsHeXoQDNQ1qL3oeqb8_OYR_6XrKN0XhUi5sDYESm0vNxJCMVkPCjKQ0MhUtWZwI53Ma0DYXgbLoAoblPxzOqlAm3V3GitVGJlxFEU6LVqKIB6CdY5RXqr73coCbUMP2dDHfQbE-8SE1MRs0y8KeNN-XOukGpVC0Su7F-pX4pkrmGCATF8W0YWVPAOf1kb9JPV6c7QVTY1LuSkOgz3jgL6KmVOcgkdiz4iyc6_CJn-QpVKsq4vEnhw"
	accountId, err := v.Verify(tkn)
	if err != nil {
		t.Errorf("verifcartion failed:%v",err)
	}
	want :="6055af8c8276e5379d4dab41"
	if accountId!=want {
		t.Errorf("wrong account id.want:%q,got:%q",want,accountId)
	}
}
