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
	sysUserFieldNames          = builderx.FieldNames(&SysUser{})
	sysUserRows                = strings.Join(sysUserFieldNames, ",")
	sysUserRowsExpectAutoSet   = strings.Join(stringx.Remove(sysUserFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	sysUserRowsWithPlaceHolder = strings.Join(stringx.Remove(sysUserFieldNames, "`id`", "`create_by`", "`create_time`", "`update_time`"), "=?,") + "=?"
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
		JobId          int64     `db:"job_id"`           // 岗位ID
		CreateBy       string    `db:"create_by"`        // 创建人
		CreateTime     time.Time `db:"create_time"`      // 创建时间
		LastUpdateBy   string    `db:"last_update_by"`   // 更新人
		LastUpdateTime time.Time `db:"last_update_time"` // 更新时间
		DelFlag        int64     `db:"del_flag"`         // 是否删除  -1：已删除  0：正常

	}

	SysUserList struct {
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
		JobId          int64     `db:"job_id"`           // 岗位ID
		CreateBy       string    `db:"create_by"`        // 创建人
		CreateTime     time.Time `db:"create_time"`      // 创建时间
		LastUpdateBy   string    `db:"last_update_by"`   // 更新人
		LastUpdateTime time.Time `db:"last_update_time"` // 更新时间
		DelFlag        int64     `db:"del_flag"`         // 是否删除  -1：已删除  0：正常
		JobName        string    `db:"job_name"`
		DeptName       string    `db:"dept_name"`
		RoleName       string    `db:"role_name"`
		RoleId         int64     `db:"role_id"`
	}
)

func NewSysUserModel(conn sqlx.SqlConn) *SysUserModel {
	return &SysUserModel{
		conn:  conn,
		table: "sys_user",
	}
}

func (m *SysUserModel) Insert(data SysUser) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, sysUserRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Name, data.NickName, data.Avatar, data.Password, data.Salt, data.Email, data.Mobile, data.Status, data.DeptId, data.JobId, data.CreateBy, data.LastUpdateBy, data.LastUpdateTime, data.DelFlag)
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

func (m *SysUserModel) FindAll(Current int64, PageSize int64) (*[]SysUserList, error) {

	//query := fmt.Sprintf("select %s from %s limit ?,?", sysUserRows, m.table)
	query := "select sys_user.*, ifnull(sj.job_name,'') as job_name, ifnull(sd.name ,'')as dept_name, ifnull(sys_role.name,'') as role_name,ifnull(sys_role.id ,'1')as role_id from sys_user   left join sys_user_role sur on sys_user.id = sur.user_id   left join sys_role on sur.role_id = sys_role.id    left join sys_job sj on sys_user.job_id = sj.id left join sys_dept sd on sys_user.dept_id = sd.id limit ?,?"
	var resp []SysUserList
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
	_, err := m.conn.Exec(query, data.Name, data.NickName, data.Avatar, data.Password, data.Salt, data.Email, data.Mobile, data.Status, data.DeptId, data.JobId, data.LastUpdateBy, data.LastUpdateTime, data.DelFlag, data.Id)
	return err
}

func (m *SysUserModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
