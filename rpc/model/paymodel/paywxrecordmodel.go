package paymodel

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ PayWxRecordModel = (*customPayWxRecordModel)(nil)

type (
	// PayWxRecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPayWxRecordModel.
	PayWxRecordModel interface {
		payWxRecordModel
		Count(ctx context.Context) (int64, error)
		FindAll(ctx context.Context, Current int64, PageSize int64) (*[]PayWxRecord, error)
		DeleteByIds(ctx context.Context, ids []int64) error
		FindOneByBusinessId(ctx context.Context, businessId string) (*PayWxRecord, error)
	}

	customPayWxRecordModel struct {
		*defaultPayWxRecordModel
	}
)

// NewPayWxRecordModel returns a model for the database table.
func NewPayWxRecordModel(conn sqlx.SqlConn) PayWxRecordModel {
	return &customPayWxRecordModel{
		defaultPayWxRecordModel: newPayWxRecordModel(conn),
	}
}
func (m *customPayWxRecordModel) FindAll(ctx context.Context, Current int64, PageSize int64) (*[]PayWxRecord, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", payWxRecordRows, m.table)
	var resp []PayWxRecord
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

func (m *customPayWxRecordModel) Count(ctx context.Context) (int64, error) {
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

func (m *customPayWxRecordModel) DeleteByIds(ctx context.Context, ids []int64) error {
	query := fmt.Sprintf("delete from %s where `id` in (?)", m.table)
	_, err := m.conn.ExecCtx(ctx, query, strings.Replace(strings.Trim(fmt.Sprint(ids), "[]"), " ", ",", -1))
	return err
}

func (m *defaultPayWxRecordModel) FindOneByBusinessId(ctx context.Context, businessId string) (*PayWxRecord, error) {
	query := fmt.Sprintf("select %s from %s where businessId = ? limit 1", payWxRecordRows, m.table)
	var resp PayWxRecord
	err := m.conn.QueryRow(&resp, query, businessId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
