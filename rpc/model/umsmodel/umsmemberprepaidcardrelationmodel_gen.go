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
	umsMemberPrepaidCardRelationFieldNames          = builder.RawFieldNames(&UmsMemberPrepaidCardRelation{})
	umsMemberPrepaidCardRelationRows                = strings.Join(umsMemberPrepaidCardRelationFieldNames, ",")
	umsMemberPrepaidCardRelationRowsExpectAutoSet   = strings.Join(stringx.Remove(umsMemberPrepaidCardRelationFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	umsMemberPrepaidCardRelationRowsWithPlaceHolder = strings.Join(stringx.Remove(umsMemberPrepaidCardRelationFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheGozeroUmsMemberPrepaidCardRelationIdPrefix = "cache:gozero:umsMemberPrepaidCardRelation:id:"
)

type (
	umsMemberPrepaidCardRelationModel interface {
		Insert(ctx context.Context, data *UmsMemberPrepaidCardRelation) (sql.Result, error)
		InsertTx(ctx context.Context, session sqlx.Session, data *UmsMemberPrepaidCardRelation) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UmsMemberPrepaidCardRelation, error)
		Update(ctx context.Context, data *UmsMemberPrepaidCardRelation) error
		UpdateTx(ctx context.Context, session sqlx.Session, data *UmsMemberPrepaidCardRelation) error
		Delete(ctx context.Context, id int64) error
		DeleteTx(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultUmsMemberPrepaidCardRelationModel struct {
		sqlc.CachedConn
		table string
	}

	UmsMemberPrepaidCardRelation struct {
		Id            int64     `db:"id"`
		MemberId      int64     `db:"member_id"`
		PrepaidCardId int64     `db:"prepaid_card_id"`
		PrepaidCardSn string    `db:"prepaid_card_sn"` // 预付卡号
		Balance       float64   `db:"balance"`         // 余额
		Status        int64     `db:"status"`          // 状态：0->正常；1->禁用
		CreateTime    time.Time `db:"create_time"`     // 创建时间
		ExpiredTime   time.Time `db:"expired_time"`    // 失效时间
		UpdateTime    time.Time `db:"update_time"`     // 更新时间
		Note          string    `db:"note"`
	}
)

func newUmsMemberPrepaidCardRelationModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUmsMemberPrepaidCardRelationModel {
	return &defaultUmsMemberPrepaidCardRelationModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`ums_member_prepaid_card_relation`",
	}
}

func (m *defaultUmsMemberPrepaidCardRelationModel) Delete(ctx context.Context, id int64) error {
	gozeroUmsMemberPrepaidCardRelationIdKey := fmt.Sprintf("%s%v", cacheGozeroUmsMemberPrepaidCardRelationIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, gozeroUmsMemberPrepaidCardRelationIdKey)
	return err
}

func (m *defaultUmsMemberPrepaidCardRelationModel) DeleteTx(ctx context.Context, session sqlx.Session, id int64) error {
	gozeroUmsMemberPrepaidCardRelationIdKey := fmt.Sprintf("%s%v", cacheGozeroUmsMemberPrepaidCardRelationIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		if session != nil {
			return session.ExecCtx(ctx, query, id)
		}
		return conn.ExecCtx(ctx, query, id)
	}, gozeroUmsMemberPrepaidCardRelationIdKey)
	return err
}

func (m *defaultUmsMemberPrepaidCardRelationModel) FindOne(ctx context.Context, id int64) (*UmsMemberPrepaidCardRelation, error) {
	gozeroUmsMemberPrepaidCardRelationIdKey := fmt.Sprintf("%s%v", cacheGozeroUmsMemberPrepaidCardRelationIdPrefix, id)
	var resp UmsMemberPrepaidCardRelation
	err := m.QueryRowCtx(ctx, &resp, gozeroUmsMemberPrepaidCardRelationIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", umsMemberPrepaidCardRelationRows, m.table)
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

func (m *defaultUmsMemberPrepaidCardRelationModel) Insert(ctx context.Context, data *UmsMemberPrepaidCardRelation) (sql.Result, error) {
	gozeroUmsMemberPrepaidCardRelationIdKey := fmt.Sprintf("%s%v", cacheGozeroUmsMemberPrepaidCardRelationIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, umsMemberPrepaidCardRelationRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.MemberId, data.PrepaidCardId, data.PrepaidCardSn, data.Balance, data.Status, data.ExpiredTime, data.Note)
	}, gozeroUmsMemberPrepaidCardRelationIdKey)
	return ret, err
}

func (m *defaultUmsMemberPrepaidCardRelationModel) InsertTx(ctx context.Context, session sqlx.Session, data *UmsMemberPrepaidCardRelation) (sql.Result, error) {
	gozeroUmsMemberPrepaidCardRelationIdKey := fmt.Sprintf("%s%v", cacheGozeroUmsMemberPrepaidCardRelationIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, umsMemberPrepaidCardRelationRowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query, data.MemberId, data.PrepaidCardId, data.PrepaidCardSn, data.Balance, data.Status, data.ExpiredTime, data.Note)
		}
		return conn.ExecCtx(ctx, query, data.MemberId, data.PrepaidCardId, data.PrepaidCardSn, data.Balance, data.Status, data.ExpiredTime, data.Note)
	}, gozeroUmsMemberPrepaidCardRelationIdKey)
	return ret, err
}
func (m *defaultUmsMemberPrepaidCardRelationModel) Update(ctx context.Context, data *UmsMemberPrepaidCardRelation) error {
	gozeroUmsMemberPrepaidCardRelationIdKey := fmt.Sprintf("%s%v", cacheGozeroUmsMemberPrepaidCardRelationIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, umsMemberPrepaidCardRelationRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.MemberId, data.PrepaidCardId, data.PrepaidCardSn, data.Balance, data.Status, data.ExpiredTime, data.Note, data.Id)
	}, gozeroUmsMemberPrepaidCardRelationIdKey)
	return err
}

func (m *defaultUmsMemberPrepaidCardRelationModel) UpdateTx(ctx context.Context, session sqlx.Session, data *UmsMemberPrepaidCardRelation) error {
	gozeroUmsMemberPrepaidCardRelationIdKey := fmt.Sprintf("%s%v", cacheGozeroUmsMemberPrepaidCardRelationIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, umsMemberPrepaidCardRelationRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.MemberId, data.PrepaidCardId, data.PrepaidCardSn, data.Balance, data.Status, data.ExpiredTime, data.Note, data.Id)
		}
		return conn.ExecCtx(ctx, query, data.MemberId, data.PrepaidCardId, data.PrepaidCardSn, data.Balance, data.Status, data.ExpiredTime, data.Note, data.Id)
	}, gozeroUmsMemberPrepaidCardRelationIdKey)
	return err
}

func (m *defaultUmsMemberPrepaidCardRelationModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheGozeroUmsMemberPrepaidCardRelationIdPrefix, primary)
}

func (m *defaultUmsMemberPrepaidCardRelationModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", umsMemberPrepaidCardRelationRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUmsMemberPrepaidCardRelationModel) tableName() string {
	return m.table
}
