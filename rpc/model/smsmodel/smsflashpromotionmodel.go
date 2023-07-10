package smsmodel

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
	"zero-admin/rpc/sms/smsclient"
)

var _ SmsFlashPromotionModel = (*customSmsFlashPromotionModel)(nil)

type (
	// SmsFlashPromotionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSmsFlashPromotionModel.
	SmsFlashPromotionModel interface {
		smsFlashPromotionModel
		Count(ctx context.Context, in *smsclient.FlashPromotionListReq) (int64, error)
		FindAll(ctx context.Context, in *smsclient.FlashPromotionListReq) (*[]SmsFlashPromotion, error)
		FindAllByCurrentDate(ctx context.Context, currentDate string) (*[]SmsFlashPromotion, error)
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

func (m *customSmsFlashPromotionModel) FindAll(ctx context.Context, in *smsclient.FlashPromotionListReq) (*[]SmsFlashPromotion, error) {

	where := "1=1"
	if len(in.Title) > 0 {
		where = where + fmt.Sprintf(" AND title like '%%%s%%'", in.Title)
	}
	if in.Status != 2 {
		where = where + fmt.Sprintf(" AND status = %d", in.Status)
	}
	if len(in.StartDate) > 0 {
		where = where + fmt.Sprintf(" AND start_date >= '%s'", in.StartDate)
	}
	if len(in.EndDate) > 0 {
		where = where + fmt.Sprintf(" AND end_date <= '%s'", in.EndDate)
	}
	query := fmt.Sprintf("select %s from %s where %s limit ?,?", smsFlashPromotionRows, m.table, where)
	var resp []SmsFlashPromotion
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

func (m *customSmsFlashPromotionModel) Count(ctx context.Context, in *smsclient.FlashPromotionListReq) (int64, error) {
	where := "1=1"
	if len(in.Title) > 0 {
		where = where + fmt.Sprintf(" AND title like '%%%s%%'", in.Title)
	}
	if in.Status != 2 {
		where = where + fmt.Sprintf(" AND status = %d", in.Status)
	}
	if len(in.StartDate) > 0 {
		where = where + fmt.Sprintf(" AND start_date >= '%s'", in.StartDate)
	}
	if len(in.EndDate) > 0 {
		where = where + fmt.Sprintf(" AND end_date <= '%s'", in.EndDate)
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

func (m *customSmsFlashPromotionModel) FindAllByCurrentDate(ctx context.Context, currentDate string) (*[]SmsFlashPromotion, error) {

	where := fmt.Sprintf("status = '1' AND start_date <= '%s' AND end_date >= '%s'", currentDate, currentDate)

	query := fmt.Sprintf("select %s from %s where %s", smsFlashPromotionRows, m.table, where)
	var resp []SmsFlashPromotion
	err := m.conn.QueryRows(&resp, query)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
