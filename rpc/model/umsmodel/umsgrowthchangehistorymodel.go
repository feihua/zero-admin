package umsmodel

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
	umsGrowthChangeHistoryFieldNames          = builderx.FieldNames(&UmsGrowthChangeHistory{})
	umsGrowthChangeHistoryRows                = strings.Join(umsGrowthChangeHistoryFieldNames, ",")
	umsGrowthChangeHistoryRowsExpectAutoSet   = strings.Join(stringx.Remove(umsGrowthChangeHistoryFieldNames, "id", "create_time", "update_time"), ",")
	umsGrowthChangeHistoryRowsWithPlaceHolder = strings.Join(stringx.Remove(umsGrowthChangeHistoryFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	UmsGrowthChangeHistoryModel struct {
		conn  sqlx.SqlConn
		table string
	}

	UmsGrowthChangeHistory struct {
		Id          int64     `db:"id"`
		MemberId    int64     `db:"member_id"`
		CreateTime  time.Time `db:"create_time"`
		ChangeType  int64     `db:"change_type"`  // 改变类型：0->增加；1->减少
		ChangeCount int64     `db:"change_count"` // 积分改变数量
		OperateMan  string    `db:"operate_man"`  // 操作人员
		OperateNote string    `db:"operate_note"` // 操作备注
		SourceType  int64     `db:"source_type"`  // 积分来源：0->购物；1->管理员修改
	}
)

func NewUmsGrowthChangeHistoryModel(conn sqlx.SqlConn) *UmsGrowthChangeHistoryModel {
	return &UmsGrowthChangeHistoryModel{
		conn:  conn,
		table: "ums_growth_change_history",
	}
}

func (m *UmsGrowthChangeHistoryModel) Insert(data UmsGrowthChangeHistory) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, umsGrowthChangeHistoryRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.MemberId, data.ChangeType, data.ChangeCount, data.OperateMan, data.OperateNote, data.SourceType)
	return ret, err
}

func (m *UmsGrowthChangeHistoryModel) FindOne(id int64) (*UmsGrowthChangeHistory, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", umsGrowthChangeHistoryRows, m.table)
	var resp UmsGrowthChangeHistory
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

func (m *UmsGrowthChangeHistoryModel) FindAll(Current int64, PageSize int64) (*[]UmsGrowthChangeHistory, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", umsGrowthChangeHistoryRows, m.table)
	var resp []UmsGrowthChangeHistory
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

func (m *UmsGrowthChangeHistoryModel) Count() (int64, error) {
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

func (m *UmsGrowthChangeHistoryModel) Update(data UmsGrowthChangeHistory) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, umsGrowthChangeHistoryRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.MemberId, data.ChangeType, data.ChangeCount, data.OperateMan, data.OperateNote, data.SourceType, data.Id)
	return err
}

func (m *UmsGrowthChangeHistoryModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
