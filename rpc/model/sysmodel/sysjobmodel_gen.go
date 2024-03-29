// Code generated by goctl. DO NOT EDIT.

package sysmodel

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	sysJobFieldNames          = builder.RawFieldNames(&SysJob{})
	sysJobRows                = strings.Join(sysJobFieldNames, ",")
	sysJobRowsExpectAutoSet   = strings.Join(stringx.Remove(sysJobFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	sysJobRowsWithPlaceHolder = strings.Join(stringx.Remove(sysJobFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	sysJobModel interface {
		Insert(ctx context.Context, data *SysJob) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SysJob, error)
		Update(ctx context.Context, data *SysJob) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSysJobModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysJob struct {
		Id         int64          `db:"id"`          // 编号
		JobName    string         `db:"job_name"`    // 职位名称
		OrderNum   int64          `db:"order_num"`   // 排序
		CreateBy   string         `db:"create_by"`   // 创建人
		CreateTime time.Time      `db:"create_time"` // 创建时间
		UpdateBy   sql.NullString `db:"update_by"`   // 更新人
		UpdateTime sql.NullTime   `db:"update_time"` // 更新时间
		DelFlag    int64          `db:"del_flag"`    // 是否删除  -1：已删除  0：正常
		Remarks    sql.NullString `db:"remarks"`     // 备注
	}
)

func newSysJobModel(conn sqlx.SqlConn) *defaultSysJobModel {
	return &defaultSysJobModel{
		conn:  conn,
		table: "`sys_job`",
	}
}

func (m *defaultSysJobModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultSysJobModel) FindOne(ctx context.Context, id int64) (*SysJob, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysJobRows, m.table)
	var resp SysJob
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysJobModel) Insert(ctx context.Context, data *SysJob) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, sysJobRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.JobName, data.OrderNum, data.CreateBy, data.UpdateBy, data.DelFlag, data.Remarks)
	return ret, err
}

func (m *defaultSysJobModel) Update(ctx context.Context, data *SysJob) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysJobRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.JobName, data.OrderNum, data.CreateBy, data.UpdateBy, data.DelFlag, data.Remarks, data.Id)
	return err
}

func (m *defaultSysJobModel) tableName() string {
	return m.table
}
