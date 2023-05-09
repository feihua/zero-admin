package sysmodel

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
	"time"
)

var _ SysUserModel = (*customSysUserModel)(nil)

type (
	// SysUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysUserModel.
	SysUserModel interface {
		sysUserModel
		Count(ctx context.Context) (int64, error)
		FindAll(ctx context.Context, Current int64, PageSize int64) (*[]SysUserList, error)
		DeleteByIds(ctx context.Context, ids []int64) error
	}

	customSysUserModel struct {
		*defaultSysUserModel
	}

	SysUserList struct {
		Id         int64          `db:"id"`          // 编号
		Name       string         `db:"name"`        // 用户名
		NickName   sql.NullString `db:"nick_name"`   // 昵称
		Avatar     sql.NullString `db:"avatar"`      // 头像
		Password   string         `db:"password"`    // 密码
		Salt       string         `db:"salt"`        // 加密盐
		Email      sql.NullString `db:"email"`       // 邮箱
		Mobile     sql.NullString `db:"mobile"`      // 手机号
		Status     int64          `db:"status"`      // 状态  0：禁用   1：正常
		DeptId     int64          `db:"dept_id"`     // 机构ID
		CreateBy   string         `db:"create_by"`   // 创建人
		CreateTime time.Time      `db:"create_time"` // 创建时间
		UpdateBy   sql.NullString `db:"update_by"`   // 更新人
		UpdateTime sql.NullTime   `db:"update_time"` // 更新时间
		DelFlag    int64          `db:"del_flag"`    // 是否删除  -1：已删除  0：正常
		JobId      int64          `db:"job_id"`      // 岗位Id
		JobName    string         `db:"job_name"`
		DeptName   string         `db:"dept_name"`
		RoleName   string         `db:"role_name"`
		RoleId     int64          `db:"role_id"`
	}
)

// NewSysUserModel returns a model for the database table.
func NewSysUserModel(conn sqlx.SqlConn) SysUserModel {
	return &customSysUserModel{
		defaultSysUserModel: newSysUserModel(conn),
	}
}

func (m *customSysUserModel) FindAll(ctx context.Context, Current int64, PageSize int64) (*[]SysUserList, error) {

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

func (m *customSysUserModel) Count(ctx context.Context) (int64, error) {
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

func (m *customSysUserModel) DeleteByIds(ctx context.Context, ids []int64) error {
	query := fmt.Sprintf("delete from %s where `id` in (?)", m.table)
	_, err := m.conn.ExecCtx(ctx, query, strings.Replace(strings.Trim(fmt.Sprint(ids), "[]"), " ", ",", -1))
	return err
}
