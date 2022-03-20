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
	sysRoleDeptFieldNames          = builderx.FieldNames(&SysRoleDept{})
	sysRoleDeptRows                = strings.Join(sysRoleDeptFieldNames, ",")
	sysRoleDeptRowsExpectAutoSet   = strings.Join(stringx.Remove(sysRoleDeptFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	sysRoleDeptRowsWithPlaceHolder = strings.Join(stringx.Remove(sysRoleDeptFieldNames, "`id`", "`create_by`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	SysRoleDeptModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysRoleDept struct {
		Id             int64     `db:"id"`               // 编号
		RoleId         int64     `db:"role_id"`          // 角色ID
		DeptId         int64     `db:"dept_id"`          // 机构ID
		CreateBy       string    `db:"create_by"`        // 创建人
		CreateTime     time.Time `db:"create_time"`      // 创建时间
		LastUpdateBy   string    `db:"last_update_by"`   // 更新人
		LastUpdateTime time.Time `db:"last_update_time"` // 更新时间
	}
)

func NewSysRoleDeptModel(conn sqlx.SqlConn) *SysRoleDeptModel {
	return &SysRoleDeptModel{
		conn:  conn,
		table: "sys_role_dept",
	}
}

func (m *SysRoleDeptModel) Insert(data SysRoleDept) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, sysRoleDeptRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.RoleId, data.DeptId, data.CreateBy, data.LastUpdateBy, data.LastUpdateTime)
	return ret, err
}

func (m *SysRoleDeptModel) FindOne(id int64) (*SysRoleDept, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", sysRoleDeptRows, m.table)
	var resp SysRoleDept
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

func (m *SysRoleDeptModel) Update(data SysRoleDept) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, sysRoleDeptRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.RoleId, data.DeptId, data.LastUpdateBy, data.LastUpdateTime, data.Id)
	return err
}

func (m *SysRoleDeptModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
