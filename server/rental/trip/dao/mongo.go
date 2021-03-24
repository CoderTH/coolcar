package dao

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	mgo "coolcar/shared/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

//Auth Mongo操作实体
type Mongo struct {
	col *mongo.Collection
}

func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col: db.Collection("trip"),
	}
}
//mongo表结构
type TripRecord struct {
	mgo.IDField
	mgo.UpdatedAtField
	Trip *rentalpb.Trip`bson:"trip"`
}

func (m *Mongo)CreateTrip(c context.Context,trip *rentalpb.Trip)(*TripRecord,error){
	r := &TripRecord{
		Trip: trip,
	}
	r.ID = mgo.NewObjID()
	r.UpdatedAt =mgo.UpdatedAt()
	_, err := m.col.InsertOne(c, r)
	if err != nil {
		return nil,err
	}
	return r,nil

}


