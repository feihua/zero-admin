package umsmodel

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
	"zero-admin/rpc/ums/ums"
)

var _ UmsMemberReadHistoryModel = (*customUmsMemberReadHistoryModel)(nil)

type (
	// UmsMemberReadHistoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsMemberReadHistoryModel.
	UmsMemberReadHistoryModel interface {
		umsMemberReadHistoryModel
		Count(ctx context.Context, in *ums.MemberReadHistoryListReq) (int64, error)
		FindAll(ctx context.Context, in *ums.MemberReadHistoryListReq) (*[]UmsMemberReadHistory, error)
		DeleteByIds(ctx context.Context, ids []int64) error
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

func (m *customUmsMemberReadHistoryModel) FindAll(ctx context.Context, in *ums.MemberReadHistoryListReq) (*[]UmsMemberReadHistory, error) {
	where := "1=1"
	if in.MemberId != 0 {
		where = where + fmt.Sprintf(" AND member_id = '%d'", in.MemberId)
	}

	if in.ProductId != 0 {
		where = where + fmt.Sprintf(" AND product_id = '%d'", in.ProductId)
	}

	query := fmt.Sprintf("select %s from %s where %s limit ?,?", umsMemberRows, m.table, where)

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

func (m *customUmsMemberReadHistoryModel) Count(ctx context.Context, in *ums.MemberReadHistoryListReq) (int64, error) {
	where := "1=1"
	if in.MemberId != 0 {
		where = where + fmt.Sprintf(" AND member_id = '%d'", in.MemberId)
	}

	if in.ProductId != 0 {
		where = where + fmt.Sprintf(" AND product_id = '%d'", in.ProductId)
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

func (m *customUmsMemberReadHistoryModel) DeleteByIds(ctx context.Context, ids []int64) error {
	// 拼接占位符 "?"
	placeholders := make([]string, len(ids))
	for i := range ids {
		placeholders[i] = "?"
	}

	// 构建删除语句
	query := fmt.Sprintf("DELETE FROM %s WHERE id IN (%s)", m.table, strings.Join(placeholders, ","))

	// 构建参数列表
	args := make([]interface{}, len(ids))
	for i, id := range ids {
		args[i] = id
	}

	// 执行删除语句
	_, err := m.conn.ExecCtx(ctx, query, args...)
	return err
}
