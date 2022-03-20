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
	sysConfigFieldNames          = builderx.FieldNames(&SysConfig{})
	sysConfigRows                = strings.Join(sysConfigFieldNames, ",")
	sysConfigRowsExpectAutoSet   = strings.Join(stringx.Remove(sysConfigFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	sysConfigRowsWithPlaceHolder = strings.Join(stringx.Remove(sysConfigFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	SysConfigModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysConfig struct {
		Id             int64     `db:"id"`               // 编号
		Value          string    `db:"value"`            // 数据值
		Label          string    `db:"label"`            // 标签名
		Type           string    `db:"type"`             // 类型
		Description    string    `db:"description"`      // 描述
		Sort           float64   `db:"sort"`             // 排序（升序）
		CreateBy       string    `db:"create_by"`        // 创建人
		CreateTime     time.Time `db:"create_time"`      // 创建时间
		LastUpdateBy   string    `db:"last_update_by"`   // 更新人
		LastUpdateTime time.Time `db:"last_update_time"` // 更新时间
		Remarks        string    `db:"remarks"`          // 备注信息
		DelFlag        int64     `db:"del_flag"`         // 是否删除  -1：已删除  0：正常
	}
)

func NewSysConfigModel(conn sqlx.SqlConn) *SysConfigModel {
	return &SysConfigModel{
		conn:  conn,
		table: "sys_config",
	}
}

func (m *SysConfigModel) Insert(data SysConfig) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, sysConfigRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Value, data.Label, data.Type, data.Description, data.Sort, data.CreateBy, data.LastUpdateBy, data.LastUpdateTime, data.Remarks, data.DelFlag)
	return ret, err
}

func (m *SysConfigModel) FindOne(id int64) (*SysConfig, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", sysConfigRows, m.table)
	var resp SysConfig
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

func (m *SysConfigModel) FindAll(Current int64, PageSize int64) (*[]SysConfig, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", sysConfigRows, m.table)
	var resp []SysConfig
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

func (m *SysConfigModel) Count() (int64, error) {
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

func (m *SysConfigModel) Update(data SysConfig) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, sysConfigRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Value, data.Label, data.Type, data.Description, data.Sort, data.CreateBy, data.LastUpdateBy, data.LastUpdateTime, data.Remarks, data.DelFlag, data.Id)
	return err
}

func (m *SysConfigModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
