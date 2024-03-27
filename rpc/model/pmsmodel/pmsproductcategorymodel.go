package pmsmodel

import (
	"context"
	"fmt"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ PmsProductCategoryModel = (*customPmsProductCategoryModel)(nil)

type (
	// PmsProductCategoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsProductCategoryModel.
	PmsProductCategoryModel interface {
		pmsProductCategoryModel
		FindAll(ctx context.Context, in *pmsclient.ProductCategoryListReq) (*[]PmsProductCategory, int64, error)
		DeleteByIds(ctx context.Context, ids []int64) error
		FindByParentId(ctx context.Context, parentId int64) (*[]PmsProductCategory, error)
	}

	customPmsProductCategoryModel struct {
		*defaultPmsProductCategoryModel
	}
)

// NewPmsProductCategoryModel returns a model for the database table.
func NewPmsProductCategoryModel(conn sqlx.SqlConn) PmsProductCategoryModel {
	return &customPmsProductCategoryModel{
		defaultPmsProductCategoryModel: newPmsProductCategoryModel(conn),
	}
}

func (m *customPmsProductCategoryModel) FindAll(ctx context.Context, in *pmsclient.ProductCategoryListReq) (*[]PmsProductCategory, int64, error) {
	where := "1=1"

	if in.ParentId != 2000 {
		where = where + fmt.Sprintf(" AND parent_id = '%d'", in.ParentId)
	}
	if in.ShowStatus != 2 {
		where = where + fmt.Sprintf(" AND show_status = '%d'", in.ShowStatus)
	}

	query := fmt.Sprintf("select %s from %s where %s limit ?,?", pmsProductCategoryRows, m.table, where)
	var resp []PmsProductCategory
	err := m.conn.QueryRowsCtx(ctx, &resp, query, 0, 1000)

	if err != nil {
		return nil, 0, err
	}
	query = fmt.Sprintf("select count(*) as count from %s where %s", m.table, where)

	var count int64
	err = m.conn.QueryRow(&count, query)

	if err != nil {
		return nil, 0, err
	}

	return &resp, count, nil
}

func (m *customPmsProductCategoryModel) DeleteByIds(ctx context.Context, ids []int64) error {
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

func (m *customPmsProductCategoryModel) FindByParentId(ctx context.Context, parentId int64) (*[]PmsProductCategory, error) {

	query := fmt.Sprintf("select %s from %s  where parent_id = ?", pmsProductCategoryRows, m.table)
	var resp []PmsProductCategory
	err := m.conn.QueryRows(&resp, query, parentId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
