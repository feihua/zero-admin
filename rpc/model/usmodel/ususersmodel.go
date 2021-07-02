package usmodel

import (
	"database/sql"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	usUsersFieldNames          = builderx.RawFieldNames(&UsUsers{})
	usUsersRows                = strings.Join(usUsersFieldNames, ",")
	usUsersRowsExpectAutoSet   = strings.Join(stringx.Remove(usUsersFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	usUsersRowsWithPlaceHolder = strings.Join(stringx.Remove(usUsersFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	usUsersRowsForInsert = strings.Join(stringx.Remove(usUsersFieldNames, "`id`"), ",")
	usUsersRowsForUpdate = strings.Join(stringx.Remove(usUsersFieldNames, "`id`", "`create_time`"), "=?,") + "=?"
	usUsersRowsForDeleteSoft = strings.Join([]string{"delete_time"}, "=?,") + "=?"

	cacheUsUsersIdPrefix          = "cache#usUsers#id#"
	cacheUsUsersPhoneNumberPrefix = "cache#usUsers#phoneNumber#"
)

type (
	UsUsersModel interface {
		Insert(data UsUsers) (sql.Result, error)
		InsertBySession(data UsUsers, session sqlx.Session) (sql.Result, error)
		FindOne(id int64) (*UsUsers, error)
		FindOneByPhoneNumber(phoneNumber string) (*UsUsers, error)
		FindAll(Current int64, PageSize int64) (*[]UsUsers, error)
		Count() (int64, error)
		Update(data UsUsers) error
		UpdateBySession(data UsUsers, session sqlx.Session) error
		Delete(id int64) error
		DeleteBySession(id int64, session sqlx.Session) error
		DeleteIdCache(id int64) error
		DeletePhoneNumberCache(phoneNumber string) error

		GetSqlCachedConn() sqlc.CachedConn
	}

	defaultUsUsersModel struct {
		sqlc.CachedConn
		table string
	}

	UsUsers struct {
		Id          int64          `db:"id"`
		PhoneNumber sql.NullString `db:"phone_number"`
		UniqueId    sql.NullString `db:"unique_id"`
		Number      sql.NullString `db:"number"`
		Email       sql.NullString `db:"email"`
		Name        sql.NullString `db:"name"`
		Password    sql.NullString `db:"password"`
		Sex         sql.NullString `db:"sex"`
		RoleId      sql.NullInt64  `db:"role_id"`
		State       sql.NullInt64  `db:"state"`
		CreateTime  sql.NullTime   `db:"create_time"`
		UpdateTime  sql.NullTime   `db:"update_time"`
		DeleteTime  sql.NullTime   `db:"delete_time"`
	}
)

func NewUsUsersModel(conn sqlx.SqlConn, c cache.CacheConf) UsUsersModel {
	return &defaultUsUsersModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`us_users`",
	}
}

func (m *defaultUsUsersModel) GetSqlCachedConn() sqlc.CachedConn {
	return m.CachedConn
}



func (m *defaultUsUsersModel) Insert(data UsUsers) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, usUsersRowsForInsert)
	data.CreateTime.Time = time.Now()
	data.CreateTime.Valid = true
	data.UpdateTime.Time = time.Now()
	data.UpdateTime.Valid = true

	ret, err := m.ExecNoCache(query, data.PhoneNumber, data.UniqueId, data.Number, data.Email, data.Name, data.Password, data.Sex, data.RoleId, data.State, data.CreateTime, data.UpdateTime, data.DeleteTime)
	if err == nil {
		id, _ := ret.LastInsertId()
		m.DeleteIdCache(id)
	}
	return ret, err
}

//for Transact
func (m *defaultUsUsersModel) InsertBySession(data UsUsers, session sqlx.Session) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, usUsersRowsForInsert)
	data.CreateTime.Time = time.Now()
	data.CreateTime.Valid = true
	data.UpdateTime.Time = time.Now()
	data.UpdateTime.Valid = true
	ret, err := session.Exec(query, data.PhoneNumber, data.UniqueId, data.Number, data.Email, data.Name, data.Password, data.Sex, data.RoleId, data.State, data.CreateTime, data.UpdateTime, data.DeleteTime)
	if err == nil {
		id, _ := ret.LastInsertId()
		m.DeleteIdCache(id)
	}
	return ret, err
}

func (m *defaultUsUsersModel) FindOne(id int64) (*UsUsers, error) {
	usUsersIdKey := fmt.Sprintf("%s%v", cacheUsUsersIdPrefix, id)
	var resp UsUsers
	err := m.QueryRow(&resp, usUsersIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? and `delete_time` is null limit 1", usUsersRows, m.table)
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

func (m *defaultUsUsersModel) FindOneByPhoneNumber(phoneNumber string) (*UsUsers, error) {
	usUsersPhoneNumberKey := fmt.Sprintf("%s%v", cacheUsUsersPhoneNumberPrefix, phoneNumber)
	var resp UsUsers

	err := m.QueryRowIndex(&resp, usUsersPhoneNumberKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `phone_number` = ? and `delete_time` is null limit 1", usUsersRows, m.table)
		if err := conn.QueryRow(&resp, query, phoneNumber); err != nil {
			logx.Info("QueryRow error:" + err.Error())
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)

	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsUsersModel) FindAll(Current int64, PageSize int64) (*[]UsUsers, error) {

	if Current < 1 {
		Current = 1
	}
	if PageSize < 1 {
		PageSize = 20
	}
	query := fmt.Sprintf("select %s from %s where `delete_time` is null limit ?,?", usUsersRows, m.table)
	var resp []UsUsers
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

func (m *defaultUsUsersModel) Count() (int64, error) {
	query := fmt.Sprintf("select count(*) as count from %s where `delete_time` is null", m.table)

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

func (m *defaultUsUsersModel) Update(data UsUsers) error {
	usUsersIdKey := fmt.Sprintf("%s%v", cacheUsUsersIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, usUsersRowsForUpdate)
		data.UpdateTime.Time = time.Now()
		data.UpdateTime.Valid = true
		return conn.Exec(query, data.PhoneNumber, data.UniqueId, data.Number, data.Email, data.Name, data.Password, data.Sex, data.RoleId, data.State, data.UpdateTime, data.DeleteTime, data.Id)
	}, usUsersIdKey)
	return err
}

//for Transact
func (m *defaultUsUsersModel) UpdateBySession(data UsUsers, session sqlx.Session) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, usUsersRowsForUpdate)
	data.UpdateTime.Time = time.Now()
	data.UpdateTime.Valid = true
	_, err := session.Exec(query, data.PhoneNumber, data.UniqueId, data.Number, data.Email, data.Name, data.Password, data.Sex, data.RoleId, data.State, data.UpdateTime, data.DeleteTime, data.Id)

	if err != nil {
		return err
	}
	if err := m.DeleteIdCache(data.Id); err != nil {
		return err
	}
	return nil
}

func (m *defaultUsUsersModel) Delete(id int64) error {

	usUsersIdKey := fmt.Sprintf("%s%v", cacheUsUsersIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, usUsersIdKey)
	return err
}

func (m *defaultUsUsersModel) DeleteSoft(id int64) error {
	usUsersIdKey := fmt.Sprintf("%s%v", cacheUsUsersIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, usUsersRowsForDeleteSoft)
		currTime := time.Now()
		return conn.Exec(query, currTime, id)
	}, usUsersIdKey)
	return err
}

//for Transact
func (m *defaultUsUsersModel) DeleteBySession(id int64, session sqlx.Session) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := session.Exec(query, id)

	if err != nil {
		return err
	}

	if err := m.DeleteIdCache(id); err != nil {
		return err
	}
	return err
}

//for Transact
func (m *defaultUsUsersModel) DeleteBySessionSoft(id int64, session sqlx.Session) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, usUsersRowsForDeleteSoft)
	currTime := time.Now()
	_, err := session.Exec(query,currTime, id)

	if err != nil {
		return err
	}

	if err := m.DeleteIdCache(id); err != nil {
		return err
	}
	return err
}

func (m *defaultUsUsersModel) DeleteIdCache(id int64) error {
	usUsersIdKey := fmt.Sprintf("%s%v", cacheUsUsersIdPrefix, id)
	return m.DelCache(usUsersIdKey)
}
func (m *defaultUsUsersModel) DeletePhoneNumberCache(phoneNumber string) error{
	usUsersPhoneNumberKey := fmt.Sprintf("%s%v", cacheUsUsersPhoneNumberPrefix, phoneNumber)
	return m.DelCache(usUsersPhoneNumberKey)
}

func (m *defaultUsUsersModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUsUsersIdPrefix, primary)
}

func (m *defaultUsUsersModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", usUsersRows, m.table)
	return conn.QueryRow(v, query, primary)
}
