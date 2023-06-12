package smsmodel

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ SmsFlashPromotionModel = (*customSmsFlashPromotionModel)(nil)

type (
	// SmsFlashPromotionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSmsFlashPromotionModel.
	SmsFlashPromotionModel interface {
		smsFlashPromotionModel
		Count(ctx context.Context, title string, status int64) (int64, error)
		FindAll(ctx context.Context, title string, status int64, Current int64, PageSize int64) (*[]SmsFlashPromotion, error)
		DeleteByIds(ctx context.Context, ids []int64) error
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

func (m *customSmsFlashPromotionModel) FindAll(ctx context.Context, title string, status int64, Current int64, PageSize int64) (*[]SmsFlashPromotion, error) {

	where := "1=1"
	if len(title) > 0 {
		where = where + fmt.Sprintf(" AND title like '%%%s%%'", title)
	}
	if status != 2 {
		where = where + fmt.Sprintf(" AND status = %d", status)
	}
	query := fmt.Sprintf("select %s from %s where %s limit ?,?", smsFlashPromotionRows, m.table, where)
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

func (m *customSmsFlashPromotionModel) Count(ctx context.Context, title string, status int64) (int64, error) {
	where := "1=1"
	if len(title) > 0 {
		where = where + fmt.Sprintf(" AND title like '%%%s%%'", title)
	}
	if status != 2 {
		where = where + fmt.Sprintf(" AND status = %d", status)
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

func (m *defaultSmsFlashPromotionModel) DeleteByIds(ctx context.Context, ids []int64) error {
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
