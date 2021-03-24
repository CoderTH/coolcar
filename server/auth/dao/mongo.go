package dao

import (
	"context"
	mgo "coolcar/shared/mongo"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const openIDField = "open_id"

//Auth Mongo操作实体
type Mongo struct {
	col *mongo.Collection
}

func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col: db.Collection("account"),
	}
}

func (m *Mongo) ResolveAccountID(c context.Context, openID string) (string, error) {
	res := m.col.FindOneAndUpdate(c, bson.M{
		openIDField: openID,
	}, mgo.Set(bson.M{
		openIDField: openID,
	}),options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After))

	err := res.Err()
	if err != nil {
		return "", fmt.Errorf("connot findOneAndUpate: %V", err)
	}
	var row mgo.IDField
	if err = res.Decode(&row); err != nil {
		return "", fmt.Errorf("cannot decode result:%V", err)
	}
	return row.ID.Hex(), nil

}
