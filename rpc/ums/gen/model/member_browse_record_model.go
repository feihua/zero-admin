package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var _ MemberBrowseRecordModel = (*customMemberBrowseRecordModel)(nil)

type (
	// MemberBrowseRecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMemberBrowseRecordModel.
	MemberBrowseRecordModel interface {
		memberBrowseRecordModel
		Deletes(ctx context.Context, id string, memberId int64) (int64, error)
		FindPage(ctx context.Context, memberId, pageNo, pageSize int64) ([]*MemberBrowseRecord, error)
	}

	customMemberBrowseRecordModel struct {
		*defaultMemberBrowseRecordModel
	}
)

// NewMemberBrowseRecordModel returns a model for the mongo.
func NewMemberBrowseRecordModel(url, db, collection string) MemberBrowseRecordModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customMemberBrowseRecordModel{
		defaultMemberBrowseRecordModel: newDefaultMemberBrowseRecordModel(conn),
	}
}
func (m *customMemberBrowseRecordModel) Deletes(ctx context.Context, id string, memberId int64) (int64, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, ErrInvalidObjectId
	}

	res, err := m.conn.DeleteOne(ctx, bson.M{"_id": oid, "memberId": memberId})
	return res, err
}
func (m *customMemberBrowseRecordModel) FindPage(ctx context.Context, memberId, pageNo, pageSize int64) ([]*MemberBrowseRecord, error) {

	var data []*MemberBrowseRecord
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
