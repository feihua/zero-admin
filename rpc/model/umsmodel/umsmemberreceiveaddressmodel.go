package umsmodel

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UmsMemberReceiveAddressModel = (*customUmsMemberReceiveAddressModel)(nil)

type (
	// UmsMemberReceiveAddressModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsMemberReceiveAddressModel.
	UmsMemberReceiveAddressModel interface {
		umsMemberReceiveAddressModel
		FindByIdAndMemberId(ctx context.Context, id int64, memberId int64) (*UmsMemberReceiveAddress, error)
	}

	customUmsMemberReceiveAddressModel struct {
		*defaultUmsMemberReceiveAddressModel
	}
)

// NewUmsMemberReceiveAddressModel returns a model for the database table.
func NewUmsMemberReceiveAddressModel(conn sqlx.SqlConn) UmsMemberReceiveAddressModel {
	return &customUmsMemberReceiveAddressModel{
		defaultUmsMemberReceiveAddressModel: newUmsMemberReceiveAddressModel(conn),
	}
}

func (m *customUmsMemberReceiveAddressModel) FindByIdAndMemberId(ctx context.Context, id int64, memberId int64) (*UmsMemberReceiveAddress, error) {
	query := fmt.Sprintf("select %s from %s where id = ? and member_id = ? limit 1", umsMemberReceiveAddressRows, m.table)
	var resp UmsMemberReceiveAddress
	err := m.conn.QueryRowCtx(ctx, &resp, query, id, memberId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
