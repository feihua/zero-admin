package umsmodel

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
	"zero-admin/rpc/ums/umsclient"
)

var _ UmsMemberModel = (*customUmsMemberModel)(nil)

type (
	// UmsMemberModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsMemberModel.
	UmsMemberModel interface {
		umsMemberModel
		Count(ctx context.Context, in *umsclient.MemberListReq) (int64, error)
		FindAll(ctx context.Context, in *umsclient.MemberListReq) (*[]UmsMember, error)
		FindMemberByNameOrPhone(ctx context.Context, name, phone string) (*UmsMember, error)
		DeleteByIds(ctx context.Context, ids []int64) error
	}

	customUmsMemberModel struct {
		*defaultUmsMemberModel
	}
)

// NewUmsMemberModel returns a model for the database table.
func NewUmsMemberModel(conn sqlx.SqlConn) UmsMemberModel {
	return &customUmsMemberModel{
		defaultUmsMemberModel: newUmsMemberModel(conn),
	}
}

func (m *customUmsMemberModel) FindAll(ctx context.Context, in *umsclient.MemberListReq) (*[]UmsMember, error) {
	where := "1=1"
	if len(in.Username) > 0 {
		where = where + fmt.Sprintf(" AND username like '%%%s%%'", in.Username)
	}
	if len(in.Phone) > 0 {
		where = where + fmt.Sprintf(" AND phone like '%%%s%%'", in.Phone)
	}

	if in.Status != 2 {
		where = where + fmt.Sprintf(" AND status = '%d'", in.Status)
	}

	query := fmt.Sprintf("select %s from %s where %s limit ?,?", umsMemberRows, m.table, where)

	var resp []UmsMember
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

func (m *customUmsMemberModel) Count(ctx context.Context, in *umsclient.MemberListReq) (int64, error) {
	where := "1=1"
	if len(in.Username) > 0 {
		where = where + fmt.Sprintf(" AND username like '%%%s%%'", in.Username)
	}
	if len(in.Phone) > 0 {
		where = where + fmt.Sprintf(" AND phone like '%%%s%%'", in.Phone)
	}

	if in.Status != 2 {
		where = where + fmt.Sprintf(" AND status = '%d'", in.Status)
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

func (m *customUmsMemberModel) FindMemberByNameOrPhone(ctx context.Context, name, phone string) (*UmsMember, error) {
	var resp UmsMember
	query := fmt.Sprintf("select %s from %s where `username` = ? or `phone` = ? limit 1", umsMemberRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, name, phone)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customUmsMemberModel) DeleteByIds(ctx context.Context, ids []int64) error {
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
