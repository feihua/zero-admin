package sysmodel

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
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
		Count(ctx context.Context, in *sysclient.UserListReq) (int64, error)
		FindAll(ctx context.Context, in *sysclient.UserListReq) (*[]SysUserList, error)
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

func (m *customSysUserModel) FindAll(ctx context.Context, in *sysclient.UserListReq) (*[]SysUserList, error) {
	where := "1=1"
	if len(in.Name) > 0 {
		where = where + fmt.Sprintf(" AND sys_user.name like '%%%s%%'", in.Name)
	}
	if len(in.Mobile) > 0 {
		where = where + fmt.Sprintf(" AND sys_user.mobile like '%%%s%%'", in.Mobile)
	}
	if in.Status != 2 {
		where = where + fmt.Sprintf(" AND sys_user.status = %d", in.Status)
	}
	where = where + fmt.Sprint(" ORDER BY create_time DESC")
	query := fmt.Sprintf("select sys_user.*, ifnull(sj.job_name,'') as job_name, ifnull(sd.name ,'')as dept_name, ifnull(sys_role.name,'') as role_name,ifnull(sys_role.id ,'0')as role_id from sys_user   left join sys_user_role sur on sys_user.id = sur.user_id   left join sys_role on sur.role_id = sys_role.id    left join sys_job sj on sys_user.job_id = sj.id left join sys_dept sd on sys_user.dept_id = sd.id where %s limit ?,?", where)
	var resp []SysUserList
	err := m.conn.QueryRows(&resp, query, (in.Current-1)*in.PageSize, in.PageSize)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customSysUserModel) Count(ctx context.Context, in *sysclient.UserListReq) (int64, error) {
	where := "1=1"
	if len(in.Name) > 0 {
		where = where + fmt.Sprintf(" AND name like '%%%s%%'", in.Name)
	}
	if len(in.Mobile) > 0 {
		where = where + fmt.Sprintf(" AND mobile like '%%%s%%'", in.Mobile)
	}
	if in.Status != 2 {
		where = where + fmt.Sprintf(" AND status = %d", in.Status)
	}
	query := fmt.Sprintf("select count(*) as count from %s where %s", m.table, where)

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
	// 拼接占位符 "?"
	placeholders := make([]string, len(ids))
	for i := range ids {
		placeholders[i] = "?"
	}

	// 构建删除语句
	query := fmt.Sprintf("DELETE FROM %s WHERE id IN (%s)", m.table, strings.Join(placeholders, ","))

	// 构建参数列表
	args := make([]interface{}, len(ids))
	for i, id := range ids {
		args[i] = id
	}

	// 执行删除语句
	_, err := m.conn.ExecCtx(ctx, query, args...)
	return err
}
