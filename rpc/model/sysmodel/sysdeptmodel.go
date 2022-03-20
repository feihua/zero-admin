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
	sysDeptFieldNames          = builderx.FieldNames(&SysDept{})
	sysDeptRows                = strings.Join(sysDeptFieldNames, ",")
	sysDeptRowsExpectAutoSet   = strings.Join(stringx.Remove(sysDeptFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	sysDeptRowsWithPlaceHolder = strings.Join(stringx.Remove(sysDeptFieldNames, "`id`", "`create_by`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	SysDeptModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysDept struct {
		Id             int64     `db:"id"`               // 编号
		Name           string    `db:"name"`             // 机构名称
		ParentId       int64     `db:"parent_id"`        // 上级机构ID，一级机构为0
		OrderNum       int64     `db:"order_num"`        // 排序
		CreateBy       string    `db:"create_by"`        // 创建人
		CreateTime     time.Time `db:"create_time"`      // 创建时间
		LastUpdateBy   string    `db:"last_update_by"`   // 更新人
		LastUpdateTime time.Time `db:"last_update_time"` // 更新时间
		DelFlag        int64     `db:"del_flag"`         // 是否删除  -1：已删除  0：正常
	}
)

func NewSysDeptModel(conn sqlx.SqlConn) *SysDeptModel {
	return &SysDeptModel{
		conn:  conn,
		table: "sys_dept",
	}
}

func (m *SysDeptModel) Insert(data SysDept) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, sysDeptRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Name, data.ParentId, data.OrderNum, data.CreateBy, data.LastUpdateBy, data.LastUpdateTime, data.DelFlag)
	return ret, err
}

func (m *SysDeptModel) FindOne(id int64) (*SysDept, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", sysDeptRows, m.table)
	var resp SysDept
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

func (m *SysDeptModel) FindAll(Current int64, PageSize int64) (*[]SysDept, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", sysDeptRows, m.table)
	var resp []SysDept
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

func (m *SysDeptModel) Count() (int64, error) {
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
func (m *SysDeptModel) Update(data SysDept) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, sysDeptRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Name, data.ParentId, data.OrderNum, data.LastUpdateBy, data.LastUpdateTime, data.DelFlag, data.Id)
	return err
}

func (m *SysDeptModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
