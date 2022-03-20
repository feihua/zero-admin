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
	sysRoleFieldNames          = builderx.FieldNames(&SysRole{})
	sysRoleRows                = strings.Join(sysRoleFieldNames, ",")
	sysRoleRowsExpectAutoSet   = strings.Join(stringx.Remove(sysRoleFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	sysRoleRowsWithPlaceHolder = strings.Join(stringx.Remove(sysRoleFieldNames, "`id`", "`create_by`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	SysRoleModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysRole struct {
		Id             int64     `db:"id"`               // 编号
		Name           string    `db:"name"`             // 角色名称
		Remark         string    `db:"remark"`           // 备注
		CreateBy       string    `db:"create_by"`        // 创建人
		CreateTime     time.Time `db:"create_time"`      // 创建时间
		LastUpdateBy   string    `db:"last_update_by"`   // 更新人
		LastUpdateTime time.Time `db:"last_update_time"` // 更新时间
		DelFlag        int64     `db:"del_flag"`         // 是否删除  -1：已删除  0：正常
		Status         int64     `db:"status"`           // 状态  1：启用  0：禁用
	}
)

func NewSysRoleModel(conn sqlx.SqlConn) *SysRoleModel {
	return &SysRoleModel{
		conn:  conn,
		table: "sys_role",
	}
}

func (m *SysRoleModel) Insert(data SysRole) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, sysRoleRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Name, data.Remark, data.CreateBy, data.LastUpdateBy, data.LastUpdateTime, data.DelFlag, data.Status)
	return ret, err
}

func (m *SysRoleModel) FindOne(id int64) (*SysRole, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", sysRoleRows, m.table)
	var resp SysRole
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

func (m *SysRoleModel) FindAll(Current int64, PageSize int64, Name string) (*[]SysRole, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", sysRoleRows, m.table)
	var resp []SysRole
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

func (m *SysRoleModel) Count() (int64, error) {
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

func (m *SysRoleModel) Update(data SysRole) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, sysRoleRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Name, data.Remark, data.LastUpdateBy, data.LastUpdateTime, data.DelFlag, data.Status, data.Id)
	return err
}

func (m *SysRoleModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
