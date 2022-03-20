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
	sysJobFieldNames          = builderx.FieldNames(&SysJob{})
	sysJobRows                = strings.Join(sysJobFieldNames, ",")
	sysJobRowsExpectAutoSet   = strings.Join(stringx.Remove(sysJobFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	sysJobRowsWithPlaceHolder = strings.Join(stringx.Remove(sysJobFieldNames, "`id`", "`create_by`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	SysJobModel interface {
		Insert(data SysJob) (sql.Result, error)
		FindOne(id int64) (*SysJob, error)
		FindAll(Current int64, PageSize int64) (*[]SysJob, error)
		Count() (int64, error)
		Update(data SysJob) error
		Delete(id int64) error
	}

	defaultSysJobModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysJob struct {
		Id             int64     `db:"id"`               // 编号
		JobName        string    `db:"job_name"`         // 职位名称
		OrderNum       int64     `db:"order_num"`        // 排序
		Remarks        string    `db:"remarks"`          // 备注
		CreateBy       string    `db:"create_by"`        // 创建人
		CreateTime     time.Time `db:"create_time"`      // 创建时间
		LastUpdateBy   string    `db:"last_update_by"`   // 更新人
		LastUpdateTime time.Time `db:"last_update_time"` // 更新时间
		DelFlag        int64     `db:"del_flag"`         // 是否删除  -1：已删除  0：正常
	}
)

func NewSysJobModel(conn sqlx.SqlConn) SysJobModel {
	return &defaultSysJobModel{
		conn:  conn,
		table: "sys_job",
	}
}

func (m *defaultSysJobModel) Insert(data SysJob) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, sysJobRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.JobName, data.OrderNum, data.Remarks, data.CreateBy, data.LastUpdateBy, data.LastUpdateTime, data.DelFlag)
	return ret, err
}

func (m *defaultSysJobModel) FindOne(id int64) (*SysJob, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", sysJobRows, m.table)
	var resp SysJob
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

func (m *defaultSysJobModel) FindAll(Current int64, PageSize int64) (*[]SysJob, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", sysJobRows, m.table)
	var resp []SysJob
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

func (m *defaultSysJobModel) Count() (int64, error) {
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

func (m *defaultSysJobModel) Update(data SysJob) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, sysJobRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.JobName, data.OrderNum, data.Remarks, data.LastUpdateBy, data.LastUpdateTime, data.DelFlag, data.Id)
	return err
}

func (m *defaultSysJobModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
