package sysmodel

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
	sysLogFieldNames          = builderx.FieldNames(&SysLog{})
	sysLogRows                = strings.Join(sysLogFieldNames, ",")
	sysLogRowsExpectAutoSet   = strings.Join(stringx.Remove(sysLogFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	sysLogRowsWithPlaceHolder = strings.Join(stringx.Remove(sysLogFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	SysLogModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysLog struct {
		Id             int64     `db:"id"`               // 编号
		UserName       string    `db:"user_name"`        // 用户名
		Operation      string    `db:"operation"`        // 用户操作
		Method         string    `db:"method"`           // 请求方法
		Params         string    `db:"params"`           // 请求参数
		Time           int64     `db:"time"`             // 执行时长(毫秒)
		Ip             string    `db:"ip"`               // IP地址
		CreateBy       string    `db:"create_by"`        // 创建人
		CreateTime     time.Time `db:"create_time"`      // 创建时间
		LastUpdateBy   string    `db:"last_update_by"`   // 更新人
		LastUpdateTime time.Time `db:"last_update_time"` // 更新时间
	}
)

func NewSysLogModel(conn sqlx.SqlConn) *SysLogModel {
	return &SysLogModel{
		conn:  conn,
		table: "sys_log",
	}
}

func (m *SysLogModel) Insert(data SysLog) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, sysLogRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.UserName, data.Operation, data.Method, data.Params, data.Time, data.Ip, data.CreateBy, data.LastUpdateBy, data.LastUpdateTime)
	return ret, err
}

func (m *SysLogModel) FindOne(id int64) (*SysLog, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", sysLogRows, m.table)
	var resp SysLog
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

func (m *SysLogModel) FindAll(Current int64, PageSize int64) (*[]SysLog, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", sysLogRows, m.table)
	var resp []SysLog
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

func (m *SysLogModel) Count() (int64, error) {
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

func (m *SysLogModel) Update(data SysLog) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, sysLogRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.UserName, data.Operation, data.Method, data.Params, data.Time, data.Ip, data.CreateBy, data.LastUpdateBy, data.LastUpdateTime, data.Id)
	return err
}

func (m *SysLogModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
