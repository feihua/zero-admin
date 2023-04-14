// Code generated by goctl. DO NOT EDIT.

package pmsmodel

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
	pmsProductLadderFieldNames          = builder.RawFieldNames(&PmsProductLadder{})
	pmsProductLadderRows                = strings.Join(pmsProductLadderFieldNames, ",")
	pmsProductLadderRowsExpectAutoSet   = strings.Join(stringx.Remove(pmsProductLadderFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	pmsProductLadderRowsWithPlaceHolder = strings.Join(stringx.Remove(pmsProductLadderFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheGozeroPmsProductLadderIdPrefix = "cache:gozero:pmsProductLadder:id:"
)

type (
	pmsProductLadderModel interface {
		Insert(ctx context.Context, data *PmsProductLadder) (sql.Result, error)
		InsertTx(ctx context.Context, session sqlx.Session, data *PmsProductLadder) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*PmsProductLadder, error)
		Update(ctx context.Context, data *PmsProductLadder) error
		UpdateTx(ctx context.Context, session sqlx.Session, data *PmsProductLadder) error
		Delete(ctx context.Context, id int64) error
		DeleteTx(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultPmsProductLadderModel struct {
		sqlc.CachedConn
		table string
	}

	PmsProductLadder struct {
		Id        int64   `db:"id"`
		ProductId int64   `db:"product_id"`
		Count     int64   `db:"count"`    // 满足的商品数量
		Discount  float64 `db:"discount"` // 折扣
		Price     float64 `db:"price"`    // 折后价格
	}
)

func newPmsProductLadderModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultPmsProductLadderModel {
	return &defaultPmsProductLadderModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`pms_product_ladder`",
	}
}

func (m *defaultPmsProductLadderModel) Delete(ctx context.Context, id int64) error {
	gozeroPmsProductLadderIdKey := fmt.Sprintf("%s%v", cacheGozeroPmsProductLadderIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, gozeroPmsProductLadderIdKey)
	return err
}

func (m *defaultPmsProductLadderModel) DeleteTx(ctx context.Context, session sqlx.Session, id int64) error {
	gozeroPmsProductLadderIdKey := fmt.Sprintf("%s%v", cacheGozeroPmsProductLadderIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		if session != nil {
			return session.ExecCtx(ctx, query, id)
		}
		return conn.ExecCtx(ctx, query, id)
	}, gozeroPmsProductLadderIdKey)
	return err
}

func (m *defaultPmsProductLadderModel) FindOne(ctx context.Context, id int64) (*PmsProductLadder, error) {
	gozeroPmsProductLadderIdKey := fmt.Sprintf("%s%v", cacheGozeroPmsProductLadderIdPrefix, id)
	var resp PmsProductLadder
	err := m.QueryRowCtx(ctx, &resp, gozeroPmsProductLadderIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", pmsProductLadderRows, m.table)
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

func (m *defaultPmsProductLadderModel) Insert(ctx context.Context, data *PmsProductLadder) (sql.Result, error) {
	gozeroPmsProductLadderIdKey := fmt.Sprintf("%s%v", cacheGozeroPmsProductLadderIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, pmsProductLadderRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ProductId, data.Count, data.Discount, data.Price)
	}, gozeroPmsProductLadderIdKey)
	return ret, err
}

func (m *defaultPmsProductLadderModel) InsertTx(ctx context.Context, session sqlx.Session, data *PmsProductLadder) (sql.Result, error) {
	gozeroPmsProductLadderIdKey := fmt.Sprintf("%s%v", cacheGozeroPmsProductLadderIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, pmsProductLadderRowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query, data.ProductId, data.Count, data.Discount, data.Price)
		}
		return conn.ExecCtx(ctx, query, data.ProductId, data.Count, data.Discount, data.Price)
	}, gozeroPmsProductLadderIdKey)
	return ret, err
}
func (m *defaultPmsProductLadderModel) Update(ctx context.Context, data *PmsProductLadder) error {
	gozeroPmsProductLadderIdKey := fmt.Sprintf("%s%v", cacheGozeroPmsProductLadderIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, pmsProductLadderRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.ProductId, data.Count, data.Discount, data.Price, data.Id)
	}, gozeroPmsProductLadderIdKey)
	return err
}

func (m *defaultPmsProductLadderModel) UpdateTx(ctx context.Context, session sqlx.Session, data *PmsProductLadder) error {
	gozeroPmsProductLadderIdKey := fmt.Sprintf("%s%v", cacheGozeroPmsProductLadderIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, pmsProductLadderRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.ProductId, data.Count, data.Discount, data.Price, data.Id)
		}
		return conn.ExecCtx(ctx, query, data.ProductId, data.Count, data.Discount, data.Price, data.Id)
	}, gozeroPmsProductLadderIdKey)
	return err
}

func (m *defaultPmsProductLadderModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheGozeroPmsProductLadderIdPrefix, primary)
}

func (m *defaultPmsProductLadderModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", pmsProductLadderRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultPmsProductLadderModel) tableName() string {
	return m.table
}
