package smsmodel

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ SmsFlashPromotionLogModel = (*customSmsFlashPromotionLogModel)(nil)

type (
	// SmsFlashPromotionLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSmsFlashPromotionLogModel.
	SmsFlashPromotionLogModel interface {
		smsFlashPromotionLogModel
		Count(ctx context.Context) (int64, error)
		FindAll(ctx context.Context, Current int64, PageSize int64) (*[]SmsFlashPromotionLog, error)
		DeleteByIds(ctx context.Context, ids []int64) error
	}

	customSmsFlashPromotionLogModel struct {
		*defaultSmsFlashPromotionLogModel
	}
)

// NewSmsFlashPromotionLogModel returns a model for the database table.
func NewSmsFlashPromotionLogModel(conn sqlx.SqlConn) SmsFlashPromotionLogModel {
	return &customSmsFlashPromotionLogModel{
		defaultSmsFlashPromotionLogModel: newSmsFlashPromotionLogModel(conn),
	}
}

func (m *customSmsFlashPromotionLogModel) FindAll(ctx context.Context, Current int64, PageSize int64) (*[]SmsFlashPromotionLog, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", smsFlashPromotionLogRows, m.table)
	var resp []SmsFlashPromotionLog
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

func (m *customSmsFlashPromotionLogModel) Count(ctx context.Context) (int64, error) {
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

func (m *customSmsFlashPromotionLogModel) DeleteByIds(ctx context.Context, ids []int64) error {
	query := fmt.Sprintf("delete from %s where `id` in (?)", m.table)
	_, err := m.conn.ExecCtx(ctx, query, strings.Replace(strings.Trim(fmt.Sprint(ids), "[]"), " ", ",", -1))
	return err
}
