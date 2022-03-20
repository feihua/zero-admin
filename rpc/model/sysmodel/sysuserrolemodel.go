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
	sysUserRoleFieldNames          = builderx.FieldNames(&SysUserRole{})
	sysUserRoleRows                = strings.Join(sysUserRoleFieldNames, ",")
	sysUserRoleRowsExpectAutoSet   = strings.Join(stringx.Remove(sysUserRoleFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	sysUserRoleRowsWithPlaceHolder = strings.Join(stringx.Remove(sysUserRoleFieldNames, "`id`", "`create_by`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	SysUserRoleModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysUserRole struct {
		Id             int64     `db:"id"`               // 编号
		UserId         int64     `db:"user_id"`          // 用户ID
		RoleId         int64     `db:"role_id"`          // 角色ID
		CreateBy       string    `db:"create_by"`        // 创建人
		CreateTime     time.Time `db:"create_time"`      // 创建时间
		LastUpdateBy   string    `db:"last_update_by"`   // 更新人
		LastUpdateTime time.Time `db:"last_update_time"` // 更新时间
	}
)

func NewSysUserRoleModel(conn sqlx.SqlConn) *SysUserRoleModel {
	return &SysUserRoleModel{
		conn:  conn,
		table: "sys_user_role",
	}
}

func (m *SysUserRoleModel) Insert(data SysUserRole) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, sysUserRoleRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.UserId, data.RoleId, data.CreateBy, data.LastUpdateBy, data.LastUpdateTime)
	return ret, err
}

func (m *SysUserRoleModel) FindOne(id int64) (*SysUserRole, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", sysUserRoleRows, m.table)
	var resp SysUserRole
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

func (m *SysUserRoleModel) Update(data SysUserRole) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, sysUserRoleRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.UserId, data.RoleId, data.LastUpdateBy, data.LastUpdateTime, data.Id)
	return err
}

func (m *SysUserRoleModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where user_id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
