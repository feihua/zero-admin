// Code generated by goctl. DO NOT EDIT.

package smsmodel

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	smsFlashPromotionFieldNames          = builder.RawFieldNames(&SmsFlashPromotion{})
	smsFlashPromotionRows                = strings.Join(smsFlashPromotionFieldNames, ",")
	smsFlashPromotionRowsExpectAutoSet   = strings.Join(stringx.Remove(smsFlashPromotionFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	smsFlashPromotionRowsWithPlaceHolder = strings.Join(stringx.Remove(smsFlashPromotionFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	smsFlashPromotionModel interface {
		Insert(ctx context.Context, data *SmsFlashPromotion) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SmsFlashPromotion, error)
		Update(ctx context.Context, data *SmsFlashPromotion) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSmsFlashPromotionModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SmsFlashPromotion struct {
		Id         int64     `db:"id"`          // 编号
		Title      string    `db:"title"`       // 标题
		StartDate  time.Time `db:"start_date"`  // 开始日期
		EndDate    time.Time `db:"end_date"`    // 结束日期
		Status     int64     `db:"status"`      // 上下线状态
		CreateTime time.Time `db:"create_time"` // 创建时间
	}
)

func newSmsFlashPromotionModel(conn sqlx.SqlConn) *defaultSmsFlashPromotionModel {
	return &defaultSmsFlashPromotionModel{
		conn:  conn,
		table: "`sms_flash_promotion`",
	}
}

func (m *defaultSmsFlashPromotionModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultSmsFlashPromotionModel) FindOne(ctx context.Context, id int64) (*SmsFlashPromotion, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", smsFlashPromotionRows, m.table)
	var resp SmsFlashPromotion
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSmsFlashPromotionModel) Insert(ctx context.Context, data *SmsFlashPromotion) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, smsFlashPromotionRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Title, data.StartDate, data.EndDate, data.Status)
	return ret, err
}

func (m *defaultSmsFlashPromotionModel) Update(ctx context.Context, data *SmsFlashPromotion) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, smsFlashPromotionRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Title, data.StartDate, data.EndDate, data.Status, data.Id)
	return err
}

func (m *defaultSmsFlashPromotionModel) tableName() string {
	return m.table
}
