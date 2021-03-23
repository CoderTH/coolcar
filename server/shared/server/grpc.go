package server

import (
	"coolcar/shared/auth"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)
//grpc服务配置
type GRPCConfig struct {
	Name string
	Addr string
	RegisterFunc func(*grpc.Server)
	AuthPublicKeyFile string
	Logger *zap.Logger
}

//	启动GRPC服务器
func RunGRPCServer(c *GRPCConfig) error {
	Namefield := zap.String("name", c.Name)
	listen, err := net.Listen("tcp", c.Addr)
	if err != nil {
		c.Logger.Fatal("cannot listen ",Namefield,zap.Error(err))
	}
	var opts []grpc.ServerOption
	if c.AuthPublicKeyFile!=""{
		in,err:= auth.Interceptor(c.AuthPublicKeyFile)
		if err !=nil{
			c.Logger.Fatal("cannot create auth interceptor",Namefield,zap.Error(err))
		}
		opts = append(opts,grpc.UnaryInterceptor(in))
	}

	s := grpc.NewServer(opts...)
	c.RegisterFunc(s)
	c.Logger.Info("server started",Namefield,zap.String("addr",c.Addr))
	return  s.Serve(listen)

}
