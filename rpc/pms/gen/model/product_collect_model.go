package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var _ ProductCollectModel = (*customProductCollectModel)(nil)

type (
	// ProductCollectModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductCollectModel.
	ProductCollectModel interface {
		productCollectModel
		Deletes(ctx context.Context, id string, memberId int64) (int64, error)
		FindPage(ctx context.Context, memberId, pageNo, pageSize int64) ([]*ProductCollect, error)
	}

	customProductCollectModel struct {
		*defaultProductCollectModel
	}
)

// NewProductCollectModel returns a model for the mongo.
func NewProductCollectModel(url, db, collection string) ProductCollectModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customProductCollectModel{
		defaultProductCollectModel: newDefaultProductCollectModel(conn),
	}
}
func (m *customProductCollectModel) Deletes(ctx context.Context, id string, memberId int64) (int64, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, ErrInvalidObjectId
	}

	res, err := m.conn.DeleteOne(ctx, bson.M{"_id": oid, "memberId": memberId})
	return res, err
}
func (m *customProductCollectModel) FindPage(ctx context.Context, memberId, pageNo, pageSize int64) ([]*ProductCollect, error) {

	var data []*ProductCollect
	opts := options.Find().SetSkip((pageNo - 1) * pageSize).SetLimit(pageSize)
	err := m.conn.Find(ctx, &data, bson.M{"memberId": memberId}, opts)
	switch err {
	case nil:
		return data, nil
	case mon.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
