package usmodel

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	usRolesFieldNames          = builderx.RawFieldNames(&UsRoles{})
	usRolesRows                = strings.Join(usRolesFieldNames, ",")
	usRolesRowsExpectAutoSet   = strings.Join(stringx.Remove(usRolesFieldNames, "`create_time`", "`update_time`"), ",")
	usRolesRowsWithPlaceHolder = strings.Join(stringx.Remove(usRolesFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheUsRolesIdPrefix = "cache#usRoles#id#"
)

type (
	UsRolesModel interface {
		Insert(data UsRoles) (sql.Result, error)
		FindOne(id int64) (*UsRoles, error)
		FindAll(Current int64, PageSize int64) (*[]UsRoles, error)
		Count() (int64, error)
		Update(data UsRoles) error
		Delete(id int64) error
	}

	defaultUsRolesModel struct {
		sqlc.CachedConn
		table string
	}

	UsRoles struct {
		Id       int64          `db:"id"`
		RoleName sql.NullString `db:"role_name"`
		Remark   sql.NullString `db:"remark"`
		CreateAt sql.NullTime   `db:"create_at"`
		UpdateAt sql.NullTime   `db:"update_at"`
		DeleteAt sql.NullTime   `db:"delete_at"`
	}
)

func NewUsRolesModel(conn sqlx.SqlConn, c cache.CacheConf) UsRolesModel {
	return &defaultUsRolesModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`us_roles`",
	}
}

func (m *defaultUsRolesModel) Insert(data UsRoles) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, usRolesRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Id, data.RoleName, data.Remark, data.CreateAt, data.UpdateAt, data.DeleteAt)

	return ret, err
}

func (m *defaultUsRolesModel) FindOne(id int64) (*UsRoles, error) {
	usRolesIdKey := fmt.Sprintf("%s%v", cacheUsRolesIdPrefix, id)
	var resp UsRoles
	err := m.QueryRow(&resp, usRolesIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", usRolesRows, m.table)
		return conn.QueryRow(v, query, id)
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

func (m *defaultUsRolesModel) FindAll(Current int64, PageSize int64) (*[]UsRoles, error) {

	if Current < 1 {
		Current = 1
	}
	if PageSize < 1 {
		PageSize = 20
	}
	query := fmt.Sprintf("select %s from %s limit ?,?", usRolesRows, m.table)
	var resp []UsRoles
	err := m.QueryRowsNoCache(&resp, query, (Current-1)*PageSize, PageSize)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsRolesModel) Count() (int64, error) {
	query := fmt.Sprintf("select count(*) as count from %s", m.table)

	var count int64
	err := m.QueryRowNoCache(&count, query)

	switch err {
	case nil:
		return count, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}

func (m *defaultUsRolesModel) Update(data UsRoles) error {
	usRolesIdKey := fmt.Sprintf("%s%v", cacheUsRolesIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, usRolesRowsWithPlaceHolder)
		return conn.Exec(query, data.RoleName, data.Remark, data.CreateAt, data.UpdateAt, data.DeleteAt, data.Id)
	}, usRolesIdKey)
	return err
}

func (m *defaultUsRolesModel) Delete(id int64) error {

	usRolesIdKey := fmt.Sprintf("%s%v", cacheUsRolesIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, usRolesIdKey)
	return err
}

func (m *defaultUsRolesModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUsRolesIdPrefix, primary)
}

func (m *defaultUsRolesModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", usRolesRows, m.table)
	return conn.QueryRow(v, query, primary)
}
