package auth

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	"coolcar/auth/dao"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

//微信登陆服务
type Service struct {
	OpenIDResolver OpenIDResolver
	Logger         *zap.Logger
	Mongo          *dao.Mongo
	TokenExpire	time.Duration
	TokenGenerator
}

//获取openID接口
type OpenIDResolver interface {
	Resolve(code string) (string, error)
}
type TokenGenerator interface {
	GenerateToken(accountID string,expire time.Duration)(string,error)
}

func (s *Service) Login(c context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	openId, err := s.OpenIDResolver.Resolve(req.Code)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "cannot resolve openid:%V", err)
	}
	accountID,err:=s.Mongo.ResolveAccountID(c,openId)
	if err != nil {
		s.Logger.Error("cannot resolve account id",zap.Error(err))
		return nil,status.Errorf(codes.Internal,"")
	}
	tkn ,err := s.TokenGenerator.GenerateToken(accountID,s.TokenExpire)
	if err != nil {
		s.Logger.Error("cannot generate token",zap.Error(err))
		return nil,status.Error(codes.Internal,"")
	}
	return &authpb.LoginResponse{
		AccessToken: tkn,
		ExpiresIn:   int32(s.TokenExpire.Seconds()),
	}, nil
}
