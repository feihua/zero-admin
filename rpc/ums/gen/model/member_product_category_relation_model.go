package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ MemberProductCategoryRelationModel = (*customMemberProductCategoryRelationModel)(nil)

type (
	// MemberProductCategoryRelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMemberProductCategoryRelationModel.
	MemberProductCategoryRelationModel interface {
		memberProductCategoryRelationModel
		Deletes(ctx context.Context, id string, memberId int64) (int64, error)
		FindPage(ctx context.Context, memberId int64) ([]*MemberProductCategoryRelation, error)
	}

	customMemberProductCategoryRelationModel struct {
		*defaultMemberProductCategoryRelationModel
	}
)

// NewMemberProductCategoryRelationModel returns a model for the mongo.
func NewMemberProductCategoryRelationModel(url, db, collection string) MemberProductCategoryRelationModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customMemberProductCategoryRelationModel{
		defaultMemberProductCategoryRelationModel: newDefaultMemberProductCategoryRelationModel(conn),
	}
}
func (m *customMemberProductCategoryRelationModel) Deletes(ctx context.Context, id string, memberId int64) (int64, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, ErrInvalidObjectId
	}

	res, err := m.conn.DeleteOne(ctx, bson.M{"_id": oid, "memberId": memberId})
	return res, err
}
func (m *customMemberProductCategoryRelationModel) FindPage(ctx context.Context, memberId int64) ([]*MemberProductCategoryRelation, error) {

	var data []*MemberProductCategoryRelation
	err := m.conn.Find(ctx, &data, bson.M{"memberId": memberId})
	switch err {
	case nil:
		return data, nil
	case mon.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
