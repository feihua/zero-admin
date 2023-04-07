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
	umsMemberMemberTagRelationFieldNames          = builder.RawFieldNames(&UmsMemberMemberTagRelation{})
	umsMemberMemberTagRelationRows                = strings.Join(umsMemberMemberTagRelationFieldNames, ",")
	umsMemberMemberTagRelationRowsExpectAutoSet   = strings.Join(stringx.Remove(umsMemberMemberTagRelationFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	umsMemberMemberTagRelationRowsWithPlaceHolder = strings.Join(stringx.Remove(umsMemberMemberTagRelationFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheGozeroUmsMemberMemberTagRelationIdPrefix = "cache:gozero:umsMemberMemberTagRelation:id:"
)

type (
	umsMemberMemberTagRelationModel interface {
		Insert(ctx context.Context, data *UmsMemberMemberTagRelation) (sql.Result, error)
		InsertTx(ctx context.Context, session sqlx.Session, data *UmsMemberMemberTagRelation) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UmsMemberMemberTagRelation, error)
		Update(ctx context.Context, data *UmsMemberMemberTagRelation) error
		UpdateTx(ctx context.Context, session sqlx.Session, data *UmsMemberMemberTagRelation) error
		Delete(ctx context.Context, id int64) error
		DeleteTx(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultUmsMemberMemberTagRelationModel struct {
		sqlc.CachedConn
		table string
	}

	UmsMemberMemberTagRelation struct {
		Id       int64 `db:"id"`
		MemberId int64 `db:"member_id"`
		TagId    int64 `db:"tag_id"`
	}
)

func newUmsMemberMemberTagRelationModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUmsMemberMemberTagRelationModel {
	return &defaultUmsMemberMemberTagRelationModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`ums_member_member_tag_relation`",
	}
}

func (m *defaultUmsMemberMemberTagRelationModel) Delete(ctx context.Context, id int64) error {
	gozeroUmsMemberMemberTagRelationIdKey := fmt.Sprintf("%s%v", cacheGozeroUmsMemberMemberTagRelationIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, gozeroUmsMemberMemberTagRelationIdKey)
	return err
}

func (m *defaultUmsMemberMemberTagRelationModel) DeleteTx(ctx context.Context, session sqlx.Session, id int64) error {
	gozeroUmsMemberMemberTagRelationIdKey := fmt.Sprintf("%s%v", cacheGozeroUmsMemberMemberTagRelationIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		if session != nil {
			return session.ExecCtx(ctx, query, id)
		}
		return conn.ExecCtx(ctx, query, id)
	}, gozeroUmsMemberMemberTagRelationIdKey)
	return err
}

func (m *defaultUmsMemberMemberTagRelationModel) FindOne(ctx context.Context, id int64) (*UmsMemberMemberTagRelation, error) {
	gozeroUmsMemberMemberTagRelationIdKey := fmt.Sprintf("%s%v", cacheGozeroUmsMemberMemberTagRelationIdPrefix, id)
	var resp UmsMemberMemberTagRelation
	err := m.QueryRowCtx(ctx, &resp, gozeroUmsMemberMemberTagRelationIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", umsMemberMemberTagRelationRows, m.table)
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

func (m *defaultUmsMemberMemberTagRelationModel) Insert(ctx context.Context, data *UmsMemberMemberTagRelation) (sql.Result, error) {
	gozeroUmsMemberMemberTagRelationIdKey := fmt.Sprintf("%s%v", cacheGozeroUmsMemberMemberTagRelationIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, umsMemberMemberTagRelationRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.MemberId, data.TagId)
	}, gozeroUmsMemberMemberTagRelationIdKey)
	return ret, err
}

func (m *defaultUmsMemberMemberTagRelationModel) InsertTx(ctx context.Context, session sqlx.Session, data *UmsMemberMemberTagRelation) (sql.Result, error) {
	gozeroUmsMemberMemberTagRelationIdKey := fmt.Sprintf("%s%v", cacheGozeroUmsMemberMemberTagRelationIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, umsMemberMemberTagRelationRowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query, data.MemberId, data.TagId)
		}
		return conn.ExecCtx(ctx, query, data.MemberId, data.TagId)
	}, gozeroUmsMemberMemberTagRelationIdKey)
	return ret, err
}
func (m *defaultUmsMemberMemberTagRelationModel) Update(ctx context.Context, data *UmsMemberMemberTagRelation) error {
	gozeroUmsMemberMemberTagRelationIdKey := fmt.Sprintf("%s%v", cacheGozeroUmsMemberMemberTagRelationIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, umsMemberMemberTagRelationRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.MemberId, data.TagId, data.Id)
	}, gozeroUmsMemberMemberTagRelationIdKey)
	return err
}

func (m *defaultUmsMemberMemberTagRelationModel) UpdateTx(ctx context.Context, session sqlx.Session, data *UmsMemberMemberTagRelation) error {
	gozeroUmsMemberMemberTagRelationIdKey := fmt.Sprintf("%s%v", cacheGozeroUmsMemberMemberTagRelationIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, umsMemberMemberTagRelationRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.MemberId, data.TagId, data.Id)
		}
		return conn.ExecCtx(ctx, query, data.MemberId, data.TagId, data.Id)
	}, gozeroUmsMemberMemberTagRelationIdKey)
	return err
}

func (m *defaultUmsMemberMemberTagRelationModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheGozeroUmsMemberMemberTagRelationIdPrefix, primary)
}

func (m *defaultUmsMemberMemberTagRelationModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", umsMemberMemberTagRelationRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUmsMemberMemberTagRelationModel) tableName() string {
	return m.table
}
