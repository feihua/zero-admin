package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ ProductCommentReplayModel = (*customProductCommentReplayModel)(nil)

type (
	// ProductCommentReplayModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductCommentReplayModel.
	ProductCommentReplayModel interface {
		productCommentReplayModel
		FindPage(ctx context.Context, commentId, pageNo, pageSize int64) ([]*ProductCommentReplay, error)
	}

	customProductCommentReplayModel struct {
		*defaultProductCommentReplayModel
	}
)

// NewProductCommentReplayModel returns a model for the mongo.
func NewProductCommentReplayModel(url, db, collection string) ProductCommentReplayModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customProductCommentReplayModel{
		defaultProductCommentReplayModel: newDefaultProductCommentReplayModel(conn),
	}
}
func (m *customProductCommentReplayModel) FindPage(ctx context.Context, commentId, pageNo, pageSize int64) ([]*ProductCommentReplay, error) {

	var data []*ProductCommentReplay
	opts := options.Find().SetSkip((pageNo - 1) * pageSize).SetLimit(pageSize)
	err := m.conn.Find(ctx, &data, bson.M{"commentId": commentId}, opts)
	switch err {
	case nil:
		return data, nil
	case mon.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
