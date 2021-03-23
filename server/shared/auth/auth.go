package auth

import (
	"context"
	"coolcar/shared/auth/token"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"os"
	"strings"
)

const (
	authorizationHeader = "authorization"
	bearerPrefix = "Bearer "
)
type tokenVerifier interface {
	Verify(token string)(string,error)
}
type interceptor struct {
	verifier tokenVerifier
}
//用户登陆拦截器生成
func Interceptor(publicKeyFile string)(grpc.UnaryServerInterceptor,error)  {
	open, err := os.Open(publicKeyFile)
	if err != nil {
		return nil, fmt.Errorf("cannot open public key file:%v",err)
	}
	read, err := ioutil.ReadAll(open)
	if err != nil {
		return nil, fmt.Errorf("cannot read public key file:%v",err)
	}
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(read)
	if err != nil {
		return nil, fmt.Errorf("cannot parse public key file:%v",err)
	}
	i:=&interceptor{
		verifier: &token.JWTVerifier{
			PublicKey: pubKey,
		},
	}
	return i.HandleReq,nil
}
//用户拦截器实现
func (i *interceptor) HandleReq(ctx context.Context,req interface{},info *grpc.UnaryServerInfo,handler grpc.UnaryHandler)(resp interface{},err error) {
	tkn,err:= tokenFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated,"")
	}
	accountId, err := i.verifier.Verify(tkn)
	if err != nil {
		return nil,status.Errorf(codes.Unauthenticated,"token not valid:%v",err)
	}
	return handler(ContextWithAccountID(ctx,AccountID(accountId)),req)
}
//从context中取出token
func tokenFromContext(ctx context.Context) (string,error) {
	m, ok := metadata.FromIncomingContext(ctx)
	if !ok{
		return "",status.Error(codes.Unauthenticated,"")
	}
	tkn := ""
	for _,v := range m[authorizationHeader]{
		if strings.HasPrefix(v,bearerPrefix){
			tkn = v[len(bearerPrefix):]
		}
	}
	if tkn==""{
		return "",status.Error(codes.Unauthenticated,"")
	}
	return tkn,nil
}
//context中的Key、
type accountIDKey struct {}

type AccountID string

func (a AccountID)String()string {
	return string(a)
}

//将accountID存入context中
func ContextWithAccountID(c context.Context,accountId AccountID)context.Context  {
	return context.WithValue(c,accountIDKey{},accountId)
}
//从context中取出AccountId
func AccountIDFromContext(c context.Context) (AccountID,error) {
	v := c.Value(accountIDKey{})
	aid,ok := v.(AccountID)
	if !ok{
		return "",status.Error(codes.Unauthenticated,"")
	}
	return aid,nil
}

