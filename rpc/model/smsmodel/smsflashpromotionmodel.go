package smsmodel

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SmsFlashPromotionModel = (*customSmsFlashPromotionModel)(nil)

type (
	// SmsFlashPromotionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSmsFlashPromotionModel.
	SmsFlashPromotionModel interface {
		smsFlashPromotionModel
		Count(ctx context.Context) (int64, error)
		FindAll(ctx context.Context, Current int64, PageSize int64) (*[]SmsFlashPromotion, error)
	}

	customSmsFlashPromotionModel struct {
		*defaultSmsFlashPromotionModel
	}
)

// NewSmsFlashPromotionModel returns a model for the database table.
func NewSmsFlashPromotionModel(conn sqlx.SqlConn) SmsFlashPromotionModel {
	return &customSmsFlashPromotionModel{
		defaultSmsFlashPromotionModel: newSmsFlashPromotionModel(conn),
	}
}

func (m *customSmsFlashPromotionModel) FindAll(ctx context.Context, Current int64, PageSize int64) (*[]SmsFlashPromotion, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", smsFlashPromotionRows, m.table)
	var resp []SmsFlashPromotion
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

func (m *customSmsFlashPromotionModel) Count(ctx context.Context) (int64, error) {
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
