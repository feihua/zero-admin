package mymongo

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ MemberProductCollectionModel = (*customMemberProductCollectionModel)(nil)

type (
	// MemberProductCollectionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMemberProductCollectionModel.
	MemberProductCollectionModel interface {
		memberProductCollectionModel
		Deletes(ctx context.Context, id string, memberId int64) (int64, error)
		FindPage(ctx context.Context, memberId, pageNo, pageSize int64) ([]*MemberProductCollection, error)
	}

	customMemberProductCollectionModel struct {
		*defaultMemberProductCollectionModel
	}
)

// NewMemberProductCollectionModel returns a model for the mongo.
func NewMemberProductCollectionModel(url, db, collection string) MemberProductCollectionModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customMemberProductCollectionModel{
		defaultMemberProductCollectionModel: newDefaultMemberProductCollectionModel(conn),
	}
}
func (m *customMemberProductCollectionModel) Deletes(ctx context.Context, id string, memberId int64) (int64, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, ErrInvalidObjectId
	}

	res, err := m.conn.DeleteOne(ctx, bson.M{"_id": oid, "memberId": memberId})
	return res, err
}
func (m *customMemberProductCollectionModel) FindPage(ctx context.Context, memberId, pageNo, pageSize int64) ([]*MemberProductCollection, error) {

	var data []*MemberProductCollection
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
