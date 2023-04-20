// Code generated by goctl. DO NOT EDIT.

package umsmodel

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	umsMemberPrepaidCardLogFieldNames          = builder.RawFieldNames(&UmsMemberPrepaidCardLog{})
	umsMemberPrepaidCardLogRows                = strings.Join(umsMemberPrepaidCardLogFieldNames, ",")
	umsMemberPrepaidCardLogRowsExpectAutoSet   = strings.Join(stringx.Remove(umsMemberPrepaidCardLogFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	umsMemberPrepaidCardLogRowsWithPlaceHolder = strings.Join(stringx.Remove(umsMemberPrepaidCardLogFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheGozeroUmsMemberPrepaidCardLogIdPrefix = "cache:gozero:umsMemberPrepaidCardLog:id:"
)

type (
	umsMemberPrepaidCardLogModel interface {
		Insert(ctx context.Context, data *UmsMemberPrepaidCardLog) (sql.Result, error)
		InsertTx(ctx context.Context, session sqlx.Session, data *UmsMemberPrepaidCardLog) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UmsMemberPrepaidCardLog, error)
		Update(ctx context.Context, data *UmsMemberPrepaidCardLog) error
		UpdateTx(ctx context.Context, session sqlx.Session, data *UmsMemberPrepaidCardLog) error
		Delete(ctx context.Context, id int64) error
		DeleteTx(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultUmsMemberPrepaidCardLogModel struct {
		sqlc.CachedConn
		table string
	}

	UmsMemberPrepaidCardLog struct {
		Id            int64     `db:"id"`
		MemberId      int64     `db:"member_id"`
		PrepaidCardId int64     `db:"prepaid_card_id"`
		OrderId       int64     `db:"order_id"`    // 订单编号
		OrderSn       string    `db:"order_sn"`    // 订单号码
		FundsType     int64     `db:"funds_type"`  // 资金类型：0进，1出
		Amount        float64   `db:"amount"`      // 金额
		CreateTime    time.Time `db:"create_time"` // 创建时间
		UpdateTime    time.Time `db:"update_time"` // 更新时间
		Note          string    `db:"note"`
	}
)

func newUmsMemberPrepaidCardLogModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUmsMemberPrepaidCardLogModel {
	return &defaultUmsMemberPrepaidCardLogModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`ums_member_prepaid_card_log`",
	}
}

func (m *defaultUmsMemberPrepaidCardLogModel) Delete(ctx context.Context, id int64) error {
	gozeroUmsMemberPrepaidCardLogIdKey := fmt.Sprintf("%s%v", cacheGozeroUmsMemberPrepaidCardLogIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, gozeroUmsMemberPrepaidCardLogIdKey)
	return err
}

func (m *defaultUmsMemberPrepaidCardLogModel) DeleteTx(ctx context.Context, session sqlx.Session, id int64) error {
	gozeroUmsMemberPrepaidCardLogIdKey := fmt.Sprintf("%s%v", cacheGozeroUmsMemberPrepaidCardLogIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		if session != nil {
			return session.ExecCtx(ctx, query, id)
		}
		return conn.ExecCtx(ctx, query, id)
	}, gozeroUmsMemberPrepaidCardLogIdKey)
	return err
}

func (m *defaultUmsMemberPrepaidCardLogModel) FindOne(ctx context.Context, id int64) (*UmsMemberPrepaidCardLog, error) {
	gozeroUmsMemberPrepaidCardLogIdKey := fmt.Sprintf("%s%v", cacheGozeroUmsMemberPrepaidCardLogIdPrefix, id)
	var resp UmsMemberPrepaidCardLog
	err := m.QueryRowCtx(ctx, &resp, gozeroUmsMemberPrepaidCardLogIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", umsMemberPrepaidCardLogRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUmsMemberPrepaidCardLogModel) Insert(ctx context.Context, data *UmsMemberPrepaidCardLog) (sql.Result, error) {
	gozeroUmsMemberPrepaidCardLogIdKey := fmt.Sprintf("%s%v", cacheGozeroUmsMemberPrepaidCardLogIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, umsMemberPrepaidCardLogRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.MemberId, data.PrepaidCardId, data.OrderId, data.OrderSn, data.FundsType, data.Amount, data.Note)
	}, gozeroUmsMemberPrepaidCardLogIdKey)
	return ret, err
}

func (m *defaultUmsMemberPrepaidCardLogModel) InsertTx(ctx context.Context, session sqlx.Session, data *UmsMemberPrepaidCardLog) (sql.Result, error) {
	gozeroUmsMemberPrepaidCardLogIdKey := fmt.Sprintf("%s%v", cacheGozeroUmsMemberPrepaidCardLogIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, umsMemberPrepaidCardLogRowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query, data.MemberId, data.PrepaidCardId, data.OrderId, data.OrderSn, data.FundsType, data.Amount, data.Note)
		}
		return conn.ExecCtx(ctx, query, data.MemberId, data.PrepaidCardId, data.OrderId, data.OrderSn, data.FundsType, data.Amount, data.Note)
	}, gozeroUmsMemberPrepaidCardLogIdKey)
	return ret, err
}
func (m *defaultUmsMemberPrepaidCardLogModel) Update(ctx context.Context, data *UmsMemberPrepaidCardLog) error {
	gozeroUmsMemberPrepaidCardLogIdKey := fmt.Sprintf("%s%v", cacheGozeroUmsMemberPrepaidCardLogIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, umsMemberPrepaidCardLogRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.MemberId, data.PrepaidCardId, data.OrderId, data.OrderSn, data.FundsType, data.Amount, data.Note, data.Id)
	}, gozeroUmsMemberPrepaidCardLogIdKey)
	return err
}

func (m *defaultUmsMemberPrepaidCardLogModel) UpdateTx(ctx context.Context, session sqlx.Session, data *UmsMemberPrepaidCardLog) error {
	gozeroUmsMemberPrepaidCardLogIdKey := fmt.Sprintf("%s%v", cacheGozeroUmsMemberPrepaidCardLogIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, umsMemberPrepaidCardLogRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.MemberId, data.PrepaidCardId, data.OrderId, data.OrderSn, data.FundsType, data.Amount, data.Note, data.Id)
		}
		return conn.ExecCtx(ctx, query, data.MemberId, data.PrepaidCardId, data.OrderId, data.OrderSn, data.FundsType, data.Amount, data.Note, data.Id)
	}, gozeroUmsMemberPrepaidCardLogIdKey)
	return err
}

func (m *defaultUmsMemberPrepaidCardLogModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheGozeroUmsMemberPrepaidCardLogIdPrefix, primary)
}

func (m *defaultUmsMemberPrepaidCardLogModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", umsMemberPrepaidCardLogRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUmsMemberPrepaidCardLogModel) tableName() string {
	return m.table
}
