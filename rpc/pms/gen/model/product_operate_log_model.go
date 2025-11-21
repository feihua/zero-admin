package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var _ ProductOperateLogModel = (*customProductOperateLogModel)(nil)

type (
	// ProductOperateLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductOperateLogModel.
	ProductOperateLogModel interface {
		productOperateLogModel
		FindPage(ctx context.Context, productId, pageNo, pageSize int64) ([]*ProductOperateLog, error)
	}

	customProductOperateLogModel struct {
		*defaultProductOperateLogModel
	}
)

// NewProductOperateLogModel returns a model for the mongo.
func NewProductOperateLogModel(url, db, collection string) ProductOperateLogModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customProductOperateLogModel{
		defaultProductOperateLogModel: newDefaultProductOperateLogModel(conn),
	}
}
func (m *customProductOperateLogModel) FindPage(ctx context.Context, productId, pageNo, pageSize int64) ([]*ProductOperateLog, error) {

	var data []*ProductOperateLog
	opts := options.Find().SetSkip((pageNo - 1) * pageSize).SetLimit(pageSize)
	err := m.conn.Find(ctx, &data, bson.M{"productId": productId}, opts)
	switch err {
	case nil:
		return data, nil
	case mon.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
