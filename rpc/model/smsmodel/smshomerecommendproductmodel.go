package smsmodel

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
	"zero-admin/rpc/sms/sms"
)

var _ SmsHomeRecommendProductModel = (*customSmsHomeRecommendProductModel)(nil)

type (
	// SmsHomeRecommendProductModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSmsHomeRecommendProductModel.
	SmsHomeRecommendProductModel interface {
		smsHomeRecommendProductModel
		Count(ctx context.Context, in *sms.HomeRecommendProductListReq) (int64, error)
		FindAll(ctx context.Context, in *sms.HomeRecommendProductListReq) (*[]SmsHomeRecommendProduct, error)
		DeleteByIds(ctx context.Context, ids []int64) error
	}

	customSmsHomeRecommendProductModel struct {
		*defaultSmsHomeRecommendProductModel
	}
)

// NewSmsHomeRecommendProductModel returns a model for the database table.
func NewSmsHomeRecommendProductModel(conn sqlx.SqlConn) SmsHomeRecommendProductModel {
	return &customSmsHomeRecommendProductModel{
		defaultSmsHomeRecommendProductModel: newSmsHomeRecommendProductModel(conn),
	}
}

func (m *customSmsHomeRecommendProductModel) FindAll(ctx context.Context, in *sms.HomeRecommendProductListReq) (*[]SmsHomeRecommendProduct, error) {

	where := "1=1"
	if len(in.ProductName) > 0 {
		where = where + fmt.Sprintf(" AND title like '%%%s%%'", in.ProductName)
	}
	if in.RecommendStatus != 2 {
		where = where + fmt.Sprintf(" AND status = %d", in.RecommendStatus)
	}
	query := fmt.Sprintf("select %s from %s where %s limit ?,?", smsHomeRecommendProductRows, m.table, where)
	var resp []SmsHomeRecommendProduct
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

func (m *customSmsHomeRecommendProductModel) Count(ctx context.Context, in *sms.HomeRecommendProductListReq) (int64, error) {
	where := "1=1"
	if len(in.ProductName) > 0 {
		where = where + fmt.Sprintf(" AND title like '%%%s%%'", in.ProductName)
	}
	if in.RecommendStatus != 2 {
		where = where + fmt.Sprintf(" AND status = %d", in.RecommendStatus)
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

func (m *customSmsHomeRecommendProductModel) DeleteByIds(ctx context.Context, ids []int64) error {
	query := fmt.Sprintf("delete from %s where `id` in (?)", m.table)
	_, err := m.conn.ExecCtx(ctx, query, strings.Replace(strings.Trim(fmt.Sprint(ids), "[]"), " ", ",", -1))
	return err
}
