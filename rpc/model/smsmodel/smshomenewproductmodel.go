package smsmodel

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
	"zero-admin/rpc/sms/smsclient"
)

var _ SmsHomeNewProductModel = (*customSmsHomeNewProductModel)(nil)

type (
	// SmsHomeNewProductModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSmsHomeNewProductModel.
	SmsHomeNewProductModel interface {
		smsHomeNewProductModel
		Count(ctx context.Context, in *smsclient.HomeNewProductListReq) (int64, error)
		FindAll(ctx context.Context, in *smsclient.HomeNewProductListReq) (*[]SmsHomeNewProduct, error)
		FindOneByBrandId(ctx context.Context, brandId int64) (*SmsHomeNewProduct, error)
		DeleteByIds(ctx context.Context, ids []int64) error
	}

	customSmsHomeNewProductModel struct {
		*defaultSmsHomeNewProductModel
	}
)

// NewSmsHomeNewProductModel returns a model for the database table.
func NewSmsHomeNewProductModel(conn sqlx.SqlConn) SmsHomeNewProductModel {
	return &customSmsHomeNewProductModel{
		defaultSmsHomeNewProductModel: newSmsHomeNewProductModel(conn),
	}
}

func (m *customSmsHomeNewProductModel) FindAll(ctx context.Context, in *smsclient.HomeNewProductListReq) (*[]SmsHomeNewProduct, error) {

	where := "1=1"
	if len(in.ProductName) > 0 {
		where = where + fmt.Sprintf(" AND product_name like '%%%s%%'", in.ProductName)
	}
	if in.RecommendStatus != 2 {
		where = where + fmt.Sprintf(" AND recommend_status = %d", in.RecommendStatus)
	}
	query := fmt.Sprintf("select %s from %s where %s limit ?,?", smsHomeNewProductRows, m.table, where)
	var resp []SmsHomeNewProduct
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

func (m *customSmsHomeNewProductModel) FindOneByBrandId(ctx context.Context, brandId int64) (*SmsHomeNewProduct, error) {

	where := fmt.Sprintf("product_id = %d", brandId)
	query := fmt.Sprintf("select %s from %s where %s ", smsHomeNewProductRows, m.table, where)
	var resp SmsHomeNewProduct
	err := m.conn.QueryRow(&resp, query)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
func (m *customSmsHomeNewProductModel) Count(ctx context.Context, in *smsclient.HomeNewProductListReq) (int64, error) {
	where := "1=1"
	if len(in.ProductName) > 0 {
		where = where + fmt.Sprintf(" AND product_name like '%%%s%%'", in.ProductName)
	}
	if in.RecommendStatus != 2 {
		where = where + fmt.Sprintf(" AND recommend_status = %d", in.RecommendStatus)
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

func (m *customSmsHomeNewProductModel) DeleteByIds(ctx context.Context, ids []int64) error {
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
