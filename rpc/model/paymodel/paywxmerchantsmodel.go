package paymodel

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ PayWxMerchantsModel = (*customPayWxMerchantsModel)(nil)

type (
	// PayWxMerchantsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPayWxMerchantsModel.
	PayWxMerchantsModel interface {
		payWxMerchantsModel
		Count(ctx context.Context) (int64, error)
		FindAll(ctx context.Context, Current int64, PageSize int64) (*[]PayWxMerchants, error)
		DeleteByIds(ctx context.Context, ids []int64) error
		FindOneByMerId(ctx context.Context, merId string, payType string) (*PayWxMerchants, error)
	}

	customPayWxMerchantsModel struct {
		*defaultPayWxMerchantsModel
	}
)

// NewPayWxMerchantsModel returns a model for the database table.
func NewPayWxMerchantsModel(conn sqlx.SqlConn) PayWxMerchantsModel {
	return &customPayWxMerchantsModel{
		defaultPayWxMerchantsModel: newPayWxMerchantsModel(conn),
	}
}
func (m *customPayWxMerchantsModel) FindAll(ctx context.Context, Current int64, PageSize int64) (*[]PayWxMerchants, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", payWxMerchantsRows, m.table)
	var resp []PayWxMerchants
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

func (m *customPayWxMerchantsModel) Count(ctx context.Context) (int64, error) {
	query := fmt.Sprintf("select count(*) as count from %s", m.table)

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

func (m *customPayWxMerchantsModel) DeleteByIds(ctx context.Context, ids []int64) error {
	query := fmt.Sprintf("delete from %s where `id` in (?)", m.table)
	_, err := m.conn.ExecCtx(ctx, query, strings.Replace(strings.Trim(fmt.Sprint(ids), "[]"), " ", ",", -1))
	return err
}

func (m *defaultPayWxMerchantsModel) FindOneByMerId(ctx context.Context, merId string, payType string) (*PayWxMerchants, error) {
	query := fmt.Sprintf("select %s from %s where mer_id = ? and pay_type = ? limit 1", payWxMerchantsRows, m.table)
	var resp PayWxMerchants
	err := m.conn.QueryRow(&resp, query, merId, payType)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
