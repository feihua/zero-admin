package mymongo

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ MemberBrandAttentionModel = (*customMemberBrandAttentionModel)(nil)

type (
	// MemberBrandAttentionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMemberBrandAttentionModel.
	MemberBrandAttentionModel interface {
		memberBrandAttentionModel
		Deletes(ctx context.Context, brandId, memberId int64) (int64, error)
		FindPage(ctx context.Context, memberId, pageNo, pageSize int64) ([]*MemberBrandAttention, error)
	}

	customMemberBrandAttentionModel struct {
		*defaultMemberBrandAttentionModel
	}
)

// NewMemberBrandAttentionModel returns a model for the mongo.
func NewMemberBrandAttentionModel(url, db, collection string) MemberBrandAttentionModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customMemberBrandAttentionModel{
		defaultMemberBrandAttentionModel: newDefaultMemberBrandAttentionModel(conn),
	}
}
func (m *customMemberBrandAttentionModel) Deletes(ctx context.Context, brandId, memberId int64) (int64, error) {

	res, err := m.conn.DeleteOne(ctx, bson.M{"brandId": brandId, "memberId": memberId})
	return res, err
}

func (m *customMemberBrandAttentionModel) FindPage(ctx context.Context, memberId, pageNo, pageSize int64) ([]*MemberBrandAttention, error) {

	var data []*MemberBrandAttention
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
