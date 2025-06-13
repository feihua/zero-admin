package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
)

var _ ProductVertifyRecordModel = (*customProductVertifyRecordModel)(nil)

type (
	// ProductVertifyRecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductVertifyRecordModel.
	ProductVertifyRecordModel interface {
		productVertifyRecordModel
		FindAll(ctx context.Context, productI int64) ([]*ProductVertifyRecord, error)
	}

	customProductVertifyRecordModel struct {
		*defaultProductVertifyRecordModel
	}
)

// NewProductVertifyRecordModel returns a model for the mongo.
func NewProductVertifyRecordModel(url, db, collection string) ProductVertifyRecordModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customProductVertifyRecordModel{
		defaultProductVertifyRecordModel: newDefaultProductVertifyRecordModel(conn),
	}
}
func (m *customProductVertifyRecordModel) FindAll(ctx context.Context, productId int64) ([]*ProductVertifyRecord, error) {

	var data []*ProductVertifyRecord
	err := m.conn.Find(ctx, &data, bson.M{"productId": productId})
	switch err {
	case nil:
		return data, nil
	case mon.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
