package sysmodel

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
	"zero-admin/rpc/sys/sys"
)

var _ SysLoginLogModel = (*customSysLoginLogModel)(nil)

type (
	// SysLoginLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysLoginLogModel.
	SysLoginLogModel interface {
		sysLoginLogModel
		Count(ctx context.Context, in *sys.LoginLogListReq) (int64, error)
		FindAll(ctx context.Context, in *sys.LoginLogListReq) (*[]SysLoginLog, error)
		DeleteByIds(ctx context.Context, ids []int64) error
	}

	customSysLoginLogModel struct {
		*defaultSysLoginLogModel
	}
)

// NewSysLoginLogModel returns a model for the database table.
func NewSysLoginLogModel(conn sqlx.SqlConn) SysLoginLogModel {
	return &customSysLoginLogModel{
		defaultSysLoginLogModel: newSysLoginLogModel(conn),
	}
}

func (m *customSysLoginLogModel) FindAll(ctx context.Context, in *sys.LoginLogListReq) (*[]SysLoginLog, error) {
	where := "1=1"
	if len(in.UserName) > 0 {
		where = where + fmt.Sprintf(" AND user_name like '%%%s%%'", in.UserName)
	}
	if len(in.Ip) > 0 {
		where = where + fmt.Sprintf(" AND ip like '%%%s%%'", in.Ip)
	}
	where = where + fmt.Sprint(" ORDER BY create_time DESC")
	query := fmt.Sprintf("select %s from %s where %s limit ?,?", sysLoginLogRows, m.table, where)
	var resp []SysLoginLog
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

func (m *customSysLoginLogModel) Count(ctx context.Context, in *sys.LoginLogListReq) (int64, error) {
	where := "1=1"
	if len(in.UserName) > 0 {
		where = where + fmt.Sprintf(" AND user_name like '%%%s%%'", in.UserName)
	}
	if len(in.Ip) > 0 {
		where = where + fmt.Sprintf(" AND ip like '%%%s%%'", in.Ip)
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

func (m *customSysLoginLogModel) DeleteByIds(ctx context.Context, ids []int64) error {
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
