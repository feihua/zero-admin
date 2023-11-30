package umsmodel

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"zero-admin/rpc/ums/umsclient"
)

var _ UmsMemberReadHistoryModel = (*customUmsMemberReadHistoryModel)(nil)

type (
	// UmsMemberReadHistoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsMemberReadHistoryModel.
	UmsMemberReadHistoryModel interface {
		umsMemberReadHistoryModel
		Count(ctx context.Context, in *umsclient.MemberReadHistoryListReq) (int64, error)
		FindAll(ctx context.Context, in *umsclient.MemberReadHistoryListReq) (*[]UmsMemberReadHistory, error)
		DeleteByIdAndMemberId(ctx context.Context, id int64, MemberId int64) error
	}

	customUmsMemberReadHistoryModel struct {
		*defaultUmsMemberReadHistoryModel
	}
)

// NewUmsMemberReadHistoryModel returns a model for the database table.
func NewUmsMemberReadHistoryModel(conn sqlx.SqlConn) UmsMemberReadHistoryModel {
	return &customUmsMemberReadHistoryModel{
		defaultUmsMemberReadHistoryModel: newUmsMemberReadHistoryModel(conn),
	}
}

func (m *customUmsMemberReadHistoryModel) FindAll(ctx context.Context, in *umsclient.MemberReadHistoryListReq) (*[]UmsMemberReadHistory, error) {
	where := "1=1"
	if in.MemberId != 0 {
		where = where + fmt.Sprintf(" AND member_id = '%d'", in.MemberId)
	}

	query := fmt.Sprintf("select %s from %s where %s limit ?,?", umsMemberReadHistoryRows, m.table, where)

	var resp []UmsMemberReadHistory
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

func (m *customUmsMemberReadHistoryModel) Count(ctx context.Context, in *umsclient.MemberReadHistoryListReq) (int64, error) {
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

func (m *customUmsMemberReadHistoryModel) DeleteByIdAndMemberId(ctx context.Context, id int64, MemberId int64) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE member_id = ? ", m.table)

	_, err := m.conn.ExecCtx(ctx, query, MemberId)
	return err
}
