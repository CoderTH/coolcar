package trip

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	Logger *zap.Logger
}

//创建trip行程服务
func (s *Service)CreateTrip(context.Context, *rentalpb.CreateTripRequest) (*rentalpb.TripEntity, error){
	return nil,status.Error(codes.Unimplemented,"")
}
//获取行程服务
func (s *Service)GetTrip(context.Context, *rentalpb.GetTripRequest) (*rentalpb.Trip, error){
	return nil,status.Error(codes.Unimplemented,"")
}
//批量获取行程
func (s *Service)GetTrips(context.Context, *rentalpb.GetTripsRequest) (*rentalpb.GetTripsResponse, error){
	return nil,status.Error(codes.Unimplemented,"")
}
//更新行程
func (s *Service)UpdateTrip(context.Context, *rentalpb.UpdateTripRequest) (*rentalpb.Trip, error){
	return nil,status.Error(codes.Unimplemented,"")
}
