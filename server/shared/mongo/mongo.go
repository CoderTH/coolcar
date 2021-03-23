package mgo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
//mongo操作库二次封装
const IDField ="_id"
type ObjID struct {
	ID primitive.ObjectID `bson:"_id"`
}
func Set(v interface{})bson.M  {
	return bson.M{
		"$set":v,
	}
}
