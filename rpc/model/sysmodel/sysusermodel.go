package sysmodel

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	sysUserFieldNames          = builderx.FieldNames(&SysUser{})
	sysUserRows                = strings.Join(sysUserFieldNames, ",")
	sysUserRowsExpectAutoSet   = strings.Join(stringx.Remove(sysUserFieldNames, "id", "create_time", "update_time"), ",")
	sysUserRowsWithPlaceHolder = strings.Join(stringx.Remove(sysUserFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	SysUserModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysUser struct {
		Id             int64     `db:"id"`               // 编号
		Name           string    `db:"name"`             // 用户名
		NickName       string    `db:"nick_name"`        // 昵称
		Avatar         string    `db:"avatar"`           // 头像
		Password       string    `db:"password"`         // 密码
		Salt           string    `db:"salt"`             // 加密盐
		Email          string    `db:"email"`            // 邮箱
		Mobile         string    `db:"mobile"`           // 手机号
		Status         int64     `db:"status"`           // 状态  0：禁用   1：正常
		DeptId         int64     `db:"dept_id"`          // 机构ID
		CreateBy       string    `db:"create_by"`        // 创建人
		CreateTime     time.Time `db:"create_time"`      // 创建时间
		LastUpdateBy   string    `db:"last_update_by"`   // 更新人
		LastUpdateTime time.Time `db:"last_update_time"` // 更新时间
		DelFlag        int64     `db:"del_flag"`         // 是否删除  -1：已删除  0：正常
	}
)

func NewSysUserModel(conn sqlx.SqlConn) *SysUserModel {
	return &SysUserModel{
		conn:  conn,
		table: "sys_user",
	}
}

func (m *SysUserModel) Insert(data SysUser) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, sysUserRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Name, data.NickName, data.Avatar, data.Password, data.Salt, data.Email, data.Mobile, data.Status, data.DeptId, data.CreateBy, data.LastUpdateBy, data.LastUpdateTime, data.DelFlag)
	return ret, err
}

func (m *SysUserModel) FindOne(id int64) (*SysUser, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", sysUserRows, m.table)
	var resp SysUser
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

func (m *SysUserModel) FindAll(Current int64, PageSize int64) (*[]SysUser, error) {

	if Current < 1 {
		Current = 1
	}
	if PageSize < 1 {
		PageSize = 20
	}
	query := fmt.Sprintf("select %s from %s limit ?,?", sysUserRows, m.table)
	var resp []SysUser
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

func (m *SysUserModel) Count() (int64, error) {
	query := fmt.Sprintf("select count(%s) from %s", "id", m.table)

	var count int64
	result, err := m.conn.Exec(query, count)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	return 10, nil
}

func (m *SysUserModel) FindOneByName(name string) (*SysUser, error) {
	var resp SysUser
	query := fmt.Sprintf("select %s from %s where name = ? limit 1", sysUserRows, m.table)
	err := m.conn.QueryRow(&resp, query, name)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *SysUserModel) Update(data SysUser) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, sysUserRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Name, data.NickName, data.Avatar, data.Password, data.Salt, data.Email, data.Mobile, data.Status, data.DeptId, data.CreateBy, data.LastUpdateBy, data.LastUpdateTime, data.DelFlag, data.Id)
	return err
}

func (m *SysUserModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
