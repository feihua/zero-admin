package umsmodel

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ UmsMemberLevelModel = (*customUmsMemberLevelModel)(nil)

type (
	// UmsMemberLevelModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsMemberLevelModel.
	UmsMemberLevelModel interface {
		umsMemberLevelModel
		Count(ctx context.Context, Name string) (int64, error)
		FindAll(ctx context.Context, Current int64, PageSize int64, Name string) (*[]UmsMemberLevel, error)
		DeleteByIds(ctx context.Context, ids []int64) error
	}

	customUmsMemberLevelModel struct {
		*defaultUmsMemberLevelModel
	}
)

// NewUmsMemberLevelModel returns a model for the database table.
func NewUmsMemberLevelModel(conn sqlx.SqlConn) UmsMemberLevelModel {
	return &customUmsMemberLevelModel{
		defaultUmsMemberLevelModel: newUmsMemberLevelModel(conn),
	}
}

func (m *customUmsMemberLevelModel) FindAll(ctx context.Context, Current int64, PageSize int64, Name string) (*[]UmsMemberLevel, error) {

	where := "1=1"
	if len(Name) > 0 {
		where = where + fmt.Sprintf(" AND name like '%%%s%%'", Name)
	}

	query := fmt.Sprintf("select %s from %s where %s limit ?,?", umsMemberLevelRows, m.table, where)

	var resp []UmsMemberLevel
	err := m.conn.QueryRows(&resp, query, (Current-1)*PageSize, PageSize)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customUmsMemberLevelModel) Count(ctx context.Context, Name string) (int64, error) {
	where := "1=1"
	if len(Name) > 0 {
		where = where + fmt.Sprintf(" AND name like '%%%s%%'", Name)
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

func (m *customUmsMemberLevelModel) DeleteByIds(ctx context.Context, ids []int64) error {
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
