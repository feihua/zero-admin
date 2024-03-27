package umsmodel

import (
	"context"
	"fmt"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ UmsMemberLoginLogModel = (*customUmsMemberLoginLogModel)(nil)

type (
	// UmsMemberLoginLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsMemberLoginLogModel.
	UmsMemberLoginLogModel interface {
		umsMemberLoginLogModel
		Count(ctx context.Context, in *umsclient.MemberLoginLogListReq) (int64, error)
		FindAll(ctx context.Context, in *umsclient.MemberLoginLogListReq) (*[]UmsMemberLoginLog, error)
		DeleteByIds(ctx context.Context, ids []int64) error
	}

	customUmsMemberLoginLogModel struct {
		*defaultUmsMemberLoginLogModel
	}
)

// NewUmsMemberLoginLogModel returns a model for the database table.
func NewUmsMemberLoginLogModel(conn sqlx.SqlConn) UmsMemberLoginLogModel {
	return &customUmsMemberLoginLogModel{
		defaultUmsMemberLoginLogModel: newUmsMemberLoginLogModel(conn),
	}
}

func (m *customUmsMemberLoginLogModel) FindAll(ctx context.Context, in *umsclient.MemberLoginLogListReq) (*[]UmsMemberLoginLog, error) {
	where := "1=1"

	if in.MemberId != 0 {
		where = where + fmt.Sprintf(" AND member_id = '%d'", in.MemberId)
	}

	query := fmt.Sprintf("select %s from %s where %s limit ?,?", umsMemberLoginLogRows, m.table, where)
	var resp []UmsMemberLoginLog
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

func (m *customUmsMemberLoginLogModel) Count(ctx context.Context, in *umsclient.MemberLoginLogListReq) (int64, error) {
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

func (m *customUmsMemberLoginLogModel) DeleteByIds(ctx context.Context, ids []int64) error {
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
