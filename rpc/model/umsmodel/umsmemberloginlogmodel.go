package umsmodel

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/builderx"
)

var (
	umsMemberLoginLogFieldNames          = builderx.FieldNames(&UmsMemberLoginLog{})
	umsMemberLoginLogRows                = strings.Join(umsMemberLoginLogFieldNames, ",")
	umsMemberLoginLogRowsExpectAutoSet   = strings.Join(stringx.Remove(umsMemberLoginLogFieldNames, "id", "create_time", "update_time"), ",")
	umsMemberLoginLogRowsWithPlaceHolder = strings.Join(stringx.Remove(umsMemberLoginLogFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	UmsMemberLoginLogModel struct {
		conn  sqlx.SqlConn
		table string
	}

	UmsMemberLoginLog struct {
		Id         int64     `db:"id"`
		MemberId   int64     `db:"member_id"`
		CreateTime time.Time `db:"create_time"`
		Ip         string    `db:"ip"`
		City       string    `db:"city"`
		LoginType  int64     `db:"login_type"` // 登录类型：0->PC；1->android;2->ios;3->小程序
		Province   string    `db:"province"`
	}
)

func NewUmsMemberLoginLogModel(conn sqlx.SqlConn) *UmsMemberLoginLogModel {
	return &UmsMemberLoginLogModel{
		conn:  conn,
		table: "ums_member_login_log",
	}
}

func (m *UmsMemberLoginLogModel) Insert(data UmsMemberLoginLog) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, umsMemberLoginLogRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.MemberId, data.Ip, data.City, data.LoginType, data.Province)
	return ret, err
}

func (m *UmsMemberLoginLogModel) FindOne(id int64) (*UmsMemberLoginLog, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", umsMemberLoginLogRows, m.table)
	var resp UmsMemberLoginLog
	err := m.conn.QueryRow(&resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *UmsMemberLoginLogModel) FindAll(Current int64, PageSize int64) (*[]UmsMemberLoginLog, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", umsMemberLoginLogRows, m.table)
	var resp []UmsMemberLoginLog
	err := m.conn.QueryRows(&resp, query, (Current-1)*PageSize, PageSize)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *UmsMemberLoginLogModel) Count() (int64, error) {
	query := fmt.Sprintf("select count(*) as count from %s", m.table)

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

func (m *UmsMemberLoginLogModel) Update(data UmsMemberLoginLog) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, umsMemberLoginLogRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.MemberId, data.Ip, data.City, data.LoginType, data.Province, data.Id)
	return err
}

func (m *UmsMemberLoginLogModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
