package umsmodel

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"zero-admin/rpc/ums/umsclient"
)

var _ UmsMemberReceiveAddressModel = (*customUmsMemberReceiveAddressModel)(nil)

type (
	// UmsMemberReceiveAddressModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsMemberReceiveAddressModel.
	UmsMemberReceiveAddressModel interface {
		umsMemberReceiveAddressModel
		Count(ctx context.Context, in *umsclient.MemberReceiveAddressListReq) (int64, error)
		FindAll(ctx context.Context, in *umsclient.MemberReceiveAddressListReq) (*[]UmsMemberReceiveAddress, error)
		DeleteByIdAndMemberId(ctx context.Context, id int64, MemberId int64) error
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

func (m *customUmsMemberReceiveAddressModel) FindAll(ctx context.Context, in *umsclient.MemberReceiveAddressListReq) (*[]UmsMemberReceiveAddress, error) {

	where := "1=1"

	if in.MemberId != 0 {
		where = where + fmt.Sprintf(" AND member_id = '%d'", in.MemberId)
	}

	query := fmt.Sprintf("select %s from %s where %s limit ?,?", umsMemberReceiveAddressRows, m.table, where)

	var resp []UmsMemberReceiveAddress
	err := m.conn.QueryRows(&resp, query, (in.Current-1)*in.PageSize, in.PageSize)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customUmsMemberReceiveAddressModel) Count(ctx context.Context, in *umsclient.MemberReceiveAddressListReq) (int64, error) {
	where := "1=1"

	if in.MemberId != 0 {
		where = where + fmt.Sprintf(" AND member_id = '%d'", in.MemberId)
	}
	query := fmt.Sprintf("select count(*) as count from %s where %s", m.table, where)

	var count int64
	err := m.conn.QueryRow(&count, query)

	switch err {
	case nil:
		return count, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}

func (m *customUmsMemberReceiveAddressModel) DeleteByIdAndMemberId(ctx context.Context, id int64, memberId int64) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE member_id = ? and id = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, memberId, id)
	return err
}

func (m *customUmsMemberReceiveAddressModel) FindByIdAndMemberId(ctx context.Context, id int64, memberId int64) (*UmsMemberReceiveAddress, error) {
	query := fmt.Sprintf("select %s from %s where id = ? and member_id = ? limit 1", umsMemberReceiveAddressRows, m.table)
	var resp UmsMemberReceiveAddress
	err := m.conn.QueryRow(&resp, query, id, memberId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
