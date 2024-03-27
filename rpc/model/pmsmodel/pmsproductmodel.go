package pmsmodel

import (
	"context"
	"fmt"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ PmsProductModel = (*customPmsProductModel)(nil)

type (
	// PmsProductModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsProductModel.
	PmsProductModel interface {
		pmsProductModel
		Count(ctx context.Context, in *pmsclient.ProductListReq) (int64, error)
		FindAll(ctx context.Context, in *pmsclient.ProductListReq) (*[]PmsProduct, error)
		DeleteByIds(ctx context.Context, ids []int64) error
		FindAllByIds(ctx context.Context, ids []int64) (*[]PmsProduct, error)
	}

	customPmsProductModel struct {
		*defaultPmsProductModel
	}
)

// NewPmsProductModel returns a model for the database table.
func NewPmsProductModel(conn sqlx.SqlConn) PmsProductModel {
	return &customPmsProductModel{
		defaultPmsProductModel: newPmsProductModel(conn),
	}
}

func (m *customPmsProductModel) FindAll(ctx context.Context, in *pmsclient.ProductListReq) (*[]PmsProduct, error) {

	where := "1=1"
	if len(in.Name) > 0 {
		where = where + fmt.Sprintf(" AND name like '%%%s%%'", in.Name)
	}
	if in.VerifyStatus != 2 {
		where = where + fmt.Sprintf(" AND verify_status = '%d'", in.VerifyStatus)
	}
	if in.ProductCategoryId != 0 {
		where = where + fmt.Sprintf(" AND product_category_id = '%d'", in.ProductCategoryId)
	}
	if in.BrandId != 0 {
		where = where + fmt.Sprintf(" AND brand_id = '%d'", in.BrandId)
	}
	if in.PublishStatus != 2 {
		where = where + fmt.Sprintf(" AND publish_status = '%d'", in.PublishStatus)
	}

	if in.DeleteStatus != 2 {
		where = where + fmt.Sprintf(" AND delete_status = '%d'", in.DeleteStatus)
	}

	query := fmt.Sprintf("select %s from %s where %s limit ?,?", pmsProductRows, m.table, where)
	var resp []PmsProduct
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

func (m *customPmsProductModel) Count(ctx context.Context, in *pmsclient.ProductListReq) (int64, error) {
	where := "1=1"
	if len(in.Name) > 0 {
		where = where + fmt.Sprintf(" AND name like '%%%s%%'", in.Name)
	}
	if in.VerifyStatus != 2 {
		where = where + fmt.Sprintf(" AND verify_status = '%d'", in.VerifyStatus)
	}

	if in.ProductCategoryId != 0 {
		where = where + fmt.Sprintf(" AND product_category_id = '%d'", in.ProductCategoryId)
	}
	if in.BrandId != 0 {
		where = where + fmt.Sprintf(" AND brand_id = '%d'", in.BrandId)
	}
	if in.PublishStatus != 2 {
		where = where + fmt.Sprintf(" AND publish_status = '%d'", in.PublishStatus)
	}

	if in.DeleteStatus != 2 {
		where = where + fmt.Sprintf(" AND delete_status = '%d'", in.DeleteStatus)
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

func (m *customPmsProductModel) DeleteByIds(ctx context.Context, ids []int64) error {
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

func (m *customPmsProductModel) FindAllByIds(ctx context.Context, ids []int64) (*[]PmsProduct, error) {
	// 拼接占位符 "?"
	placeholders := make([]string, len(ids))
	for i := range ids {
		placeholders[i] = "?"
	}

	// 构建删除语句
	//query := fmt.Sprintf("DELETE FROM %s WHERE id IN (%s)", m.table, strings.Join(placeholders, ","))

	// 构建参数列表
	args := make([]interface{}, len(ids))
	for i, id := range ids {
		args[i] = id
	}

	query := fmt.Sprintf("select %s from %s where `id` in (%s)", pmsProductRows, m.table, strings.Join(placeholders, ","))
	var resp []PmsProduct
	err := m.conn.QueryRows(&resp, query, args...)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
