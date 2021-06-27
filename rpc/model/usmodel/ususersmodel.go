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
	usUsersFieldNames          = builderx.RawFieldNames(&UsUsers{})
	usUsersRows                = strings.Join(usUsersFieldNames, ",")
	usUsersRowsExpectAutoSet   = strings.Join(stringx.Remove(usUsersFieldNames, "`id`","`create_time`","`update_time`"), ",")
	usUsersRowsWithPlaceHolder = strings.Join(stringx.Remove(usUsersFieldNames, "`id`","`create_time`","`update_time`"), "=?,") + "=?"

	cacheUsUsersIdPrefix = "cache#usUsers#id#"
	cacheUsUsersPhoneNumberPrefix = "cache#usUsers#phoneNumber#"
)

type (
	UsUsersModel interface {
		Insert(data UsUsers) (sql.Result, error)
		FindOne(id int64) (*UsUsers, error)
		FindOneByPhoneNumber(phoneNumber string) (*UsUsers, error)
		FindAll(Current int64, PageSize int64) (*[]UsUsers, error)
		Count() (int64, error)
		Update(data UsUsers) error
		Delete(id int64) error
	}

	defaultUsUsersModel struct {
		sqlc.CachedConn
		table string
	}

	UsUsers struct {
		Id          int64          `db:"id"`
		PhoneNumber sql.NullString `db:"phone_number"`
		UniqueId    sql.NullString `db:"unique_id"`
		Number      sql.NullString  `db:"number"`
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

func (m *defaultUsUsersModel) Insert(data UsUsers) (sql.Result, error) {
	//query := fmt.Sprintf("insert into %s (%s) values ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, usUsersRowsExpectAutoSet)
	//ret, err := m.ExecNoCache(query,  data.PhoneNumber, data.UniqueId, data.Number, data.Email, data.Name, data.Password, data.Sex, data.RoleId, data.State, data.DeleteTime)
	//return ret, err

	//会清除usUsersPhoneNumberKey的缓存
	usUsersPhoneNumberKey := fmt.Sprintf("%s%v", cacheUsUsersPhoneNumberPrefix, data.PhoneNumber.String)
	ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, usUsersRowsExpectAutoSet)
		return conn.Exec(query,  data.PhoneNumber, data.UniqueId, data.Number, data.Email, data.Name, data.Password, data.Sex, data.RoleId, data.State, data.DeleteTime)
	}, usUsersPhoneNumberKey)
	//fmt.Println("usUsersPhoneNumberKey11111:"+usUsersPhoneNumberKey)
	return ret, err
}

func (m *defaultUsUsersModel) FindOne(id int64) (*UsUsers, error) {
	usUsersIdKey := fmt.Sprintf("%s%v", cacheUsUsersIdPrefix, id)
	var resp UsUsers
	err := m.QueryRow(&resp, usUsersIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", usUsersRows, m.table)
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

func (m *defaultUsUsersModel) FindOneByPhoneNumber(phoneNumber string) (*UsUsers, error){
	usUsersIdKey := fmt.Sprintf("%s%v", cacheUsUsersPhoneNumberPrefix, phoneNumber)
	var resp UsUsers
	err := m.QueryRow(&resp, usUsersIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `phone_number` = ? limit 1", usUsersRows, m.table)
		return conn.QueryRow(v, query, phoneNumber)
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

func (m *defaultUsUsersModel) FindAll(Current int64, PageSize int64) (*[]UsUsers, error) {

	if Current < 1 {
		Current = 1
	}
	if PageSize < 1 {
		PageSize = 20
	}
	query := fmt.Sprintf("select %s from %s limit ?,?", usUsersRows, m.table)
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


func (m *defaultUsUsersModel) Update(data UsUsers) error {
	usUsersIdKey := fmt.Sprintf("%s%v", cacheUsUsersIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, usUsersRowsWithPlaceHolder)
		return conn.Exec(query, data.PhoneNumber, data.UniqueId, data.Number, data.Email, data.Name, data.Password, data.Sex, data.RoleId, data.State, data.DeleteTime, data.Id)
	}, usUsersIdKey)
	return err
}

func (m *defaultUsUsersModel) Delete(id int64) error {

	usUsersIdKey := fmt.Sprintf("%s%v", cacheUsUsersIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, usUsersIdKey)
	return err
}

func (m *defaultUsUsersModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUsUsersIdPrefix, primary)
}

func (m *defaultUsUsersModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", usUsersRows, m.table)
	return conn.QueryRow(v, query, primary)
}
