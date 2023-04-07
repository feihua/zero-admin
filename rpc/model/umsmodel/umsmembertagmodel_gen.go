// Code generated by goctl. DO NOT EDIT.

package umsmodel

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	umsMemberTagFieldNames          = builder.RawFieldNames(&UmsMemberTag{})
	umsMemberTagRows                = strings.Join(umsMemberTagFieldNames, ",")
	umsMemberTagRowsExpectAutoSet   = strings.Join(stringx.Remove(umsMemberTagFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	umsMemberTagRowsWithPlaceHolder = strings.Join(stringx.Remove(umsMemberTagFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheGozeroUmsMemberTagIdPrefix = "cache:gozero:umsMemberTag:id:"
)

type (
	umsMemberTagModel interface {
		Insert(ctx context.Context, data *UmsMemberTag) (sql.Result, error)
		InsertTx(ctx context.Context, session sqlx.Session, data *UmsMemberTag) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UmsMemberTag, error)
		Update(ctx context.Context, data *UmsMemberTag) error
		UpdateTx(ctx context.Context, session sqlx.Session, data *UmsMemberTag) error
		Delete(ctx context.Context, id int64) error
		DeleteTx(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultUmsMemberTagModel struct {
		sqlc.CachedConn
		table string
	}

	UmsMemberTag struct {
		Id                int64   `db:"id"`
		Name              string  `db:"name"`
		FinishOrderCount  int64   `db:"finish_order_count"`  // 自动打标签完成订单数量
		FinishOrderAmount float64 `db:"finish_order_amount"` // 自动打标签完成订单金额
	}
)

func newUmsMemberTagModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUmsMemberTagModel {
	return &defaultUmsMemberTagModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`ums_member_tag`",
	}
}

func (m *defaultUmsMemberTagModel) Delete(ctx context.Context, id int64) error {
	gozeroUmsMemberTagIdKey := fmt.Sprintf("%s%v", cacheGozeroUmsMemberTagIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, gozeroUmsMemberTagIdKey)
	return err
}

func (m *defaultUmsMemberTagModel) DeleteTx(ctx context.Context, session sqlx.Session, id int64) error {
	gozeroUmsMemberTagIdKey := fmt.Sprintf("%s%v", cacheGozeroUmsMemberTagIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		if session != nil {
			return session.ExecCtx(ctx, query, id)
		}
		return conn.ExecCtx(ctx, query, id)
	}, gozeroUmsMemberTagIdKey)
	return err
}

func (m *defaultUmsMemberTagModel) FindOne(ctx context.Context, id int64) (*UmsMemberTag, error) {
	gozeroUmsMemberTagIdKey := fmt.Sprintf("%s%v", cacheGozeroUmsMemberTagIdPrefix, id)
	var resp UmsMemberTag
	err := m.QueryRowCtx(ctx, &resp, gozeroUmsMemberTagIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", umsMemberTagRows, m.table)
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

func (m *defaultUmsMemberTagModel) Insert(ctx context.Context, data *UmsMemberTag) (sql.Result, error) {
	gozeroUmsMemberTagIdKey := fmt.Sprintf("%s%v", cacheGozeroUmsMemberTagIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, umsMemberTagRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Name, data.FinishOrderCount, data.FinishOrderAmount)
	}, gozeroUmsMemberTagIdKey)
	return ret, err
}

func (m *defaultUmsMemberTagModel) InsertTx(ctx context.Context, session sqlx.Session, data *UmsMemberTag) (sql.Result, error) {
	gozeroUmsMemberTagIdKey := fmt.Sprintf("%s%v", cacheGozeroUmsMemberTagIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, umsMemberTagRowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Name, data.FinishOrderCount, data.FinishOrderAmount)
		}
		return conn.ExecCtx(ctx, query, data.Name, data.FinishOrderCount, data.FinishOrderAmount)
	}, gozeroUmsMemberTagIdKey)
	return ret, err
}
func (m *defaultUmsMemberTagModel) Update(ctx context.Context, data *UmsMemberTag) error {
	gozeroUmsMemberTagIdKey := fmt.Sprintf("%s%v", cacheGozeroUmsMemberTagIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, umsMemberTagRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Name, data.FinishOrderCount, data.FinishOrderAmount, data.Id)
	}, gozeroUmsMemberTagIdKey)
	return err
}

func (m *defaultUmsMemberTagModel) UpdateTx(ctx context.Context, session sqlx.Session, data *UmsMemberTag) error {
	gozeroUmsMemberTagIdKey := fmt.Sprintf("%s%v", cacheGozeroUmsMemberTagIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, umsMemberTagRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Name, data.FinishOrderCount, data.FinishOrderAmount, data.Id)
		}
		return conn.ExecCtx(ctx, query, data.Name, data.FinishOrderCount, data.FinishOrderAmount, data.Id)
	}, gozeroUmsMemberTagIdKey)
	return err
}

func (m *defaultUmsMemberTagModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheGozeroUmsMemberTagIdPrefix, primary)
}

func (m *defaultUmsMemberTagModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", umsMemberTagRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUmsMemberTagModel) tableName() string {
	return m.table
}
