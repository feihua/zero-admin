package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var _ ProductCommentModel = (*customProductCommentModel)(nil)

type (
	// ProductCommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductCommentModel.
	ProductCommentModel interface {
		productCommentModel
		FindPage(ctx context.Context, productId, pageNo, pageSize int64) ([]*ProductComment, error)
	}

	customProductCommentModel struct {
		*defaultProductCommentModel
	}
)

// NewProductCommentModel returns a model for the mongo.
func NewProductCommentModel(url, db, collection string) ProductCommentModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customProductCommentModel{
		defaultProductCommentModel: newDefaultProductCommentModel(conn),
	}
}
func (m *customProductCommentModel) FindPage(ctx context.Context, productId, pageNo, pageSize int64) ([]*ProductComment, error) {

	var data []*ProductComment
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
