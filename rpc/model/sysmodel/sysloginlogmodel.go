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
	sysLoginLogFieldNames          = builderx.FieldNames(&SysLoginLog{})
	sysLoginLogRows                = strings.Join(sysLoginLogFieldNames, ",")
	sysLoginLogRowsExpectAutoSet   = strings.Join(stringx.Remove(sysLoginLogFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	sysLoginLogRowsWithPlaceHolder = strings.Join(stringx.Remove(sysLoginLogFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	SysLoginLogModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysLoginLog struct {
		Id             int64     `db:"id"`               // 编号
		UserName       string    `db:"user_name"`        // 用户名
		Status         string    `db:"status"`           // 登录状态（online:在线，登录初始状态，方便统计在线人数；login:退出登录后将online置为login；logout:退出登录）
		Ip             string    `db:"ip"`               // IP地址
		CreateBy       string    `db:"create_by"`        // 创建人
		CreateTime     time.Time `db:"create_time"`      // 创建时间
		LastUpdateBy   string    `db:"last_update_by"`   // 更新人
		LastUpdateTime time.Time `db:"last_update_time"` // 更新时间
	}
)

func NewSysLoginLogModel(conn sqlx.SqlConn) *SysLoginLogModel {
	return &SysLoginLogModel{
		conn:  conn,
		table: "sys_login_log",
	}
}

func (m *SysLoginLogModel) Insert(data SysLoginLog) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, sysLoginLogRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.UserName, data.Status, data.Ip, data.CreateBy, data.LastUpdateBy, data.LastUpdateTime)
	return ret, err
}

func (m *SysLoginLogModel) FindOne(id int64) (*SysLoginLog, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", sysLoginLogRows, m.table)
	var resp SysLoginLog
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

func (m *SysLoginLogModel) FindAll(Current int64, PageSize int64) (*[]SysLoginLog, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", sysLoginLogRows, m.table)
	var resp []SysLoginLog
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

func (m *SysLoginLogModel) Count() (int64, error) {
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
func (m *SysLoginLogModel) Update(data SysLoginLog) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, sysLoginLogRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.UserName, data.Status, data.Ip, data.CreateBy, data.LastUpdateBy, data.LastUpdateTime, data.Id)
	return err
}

func (m *SysLoginLogModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
