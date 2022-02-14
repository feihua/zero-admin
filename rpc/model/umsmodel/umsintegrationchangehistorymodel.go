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
	umsIntegrationChangeHistoryFieldNames          = builderx.FieldNames(&UmsIntegrationChangeHistory{})
	umsIntegrationChangeHistoryRows                = strings.Join(umsIntegrationChangeHistoryFieldNames, ",")
	umsIntegrationChangeHistoryRowsExpectAutoSet   = strings.Join(stringx.Remove(umsIntegrationChangeHistoryFieldNames, "id", "create_time", "update_time"), ",")
	umsIntegrationChangeHistoryRowsWithPlaceHolder = strings.Join(stringx.Remove(umsIntegrationChangeHistoryFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	UmsIntegrationChangeHistoryModel struct {
		conn  sqlx.SqlConn
		table string
	}

	UmsIntegrationChangeHistory struct {
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

func NewUmsIntegrationChangeHistoryModel(conn sqlx.SqlConn) *UmsIntegrationChangeHistoryModel {
	return &UmsIntegrationChangeHistoryModel{
		conn:  conn,
		table: "ums_integration_change_history",
	}
}

func (m *UmsIntegrationChangeHistoryModel) Insert(data UmsIntegrationChangeHistory) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, umsIntegrationChangeHistoryRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.MemberId, data.ChangeType, data.ChangeCount, data.OperateMan, data.OperateNote, data.SourceType)
	return ret, err
}

func (m *UmsIntegrationChangeHistoryModel) FindOne(id int64) (*UmsIntegrationChangeHistory, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", umsIntegrationChangeHistoryRows, m.table)
	var resp UmsIntegrationChangeHistory
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

func (m *UmsIntegrationChangeHistoryModel) FindAll(Current int64, PageSize int64) (*[]UmsIntegrationChangeHistory, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", umsIntegrationChangeHistoryRows, m.table)
	var resp []UmsIntegrationChangeHistory
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

func (m *UmsIntegrationChangeHistoryModel) Count() (int64, error) {
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

func (m *UmsIntegrationChangeHistoryModel) Update(data UmsIntegrationChangeHistory) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, umsIntegrationChangeHistoryRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.MemberId, data.ChangeType, data.ChangeCount, data.OperateMan, data.OperateNote, data.SourceType, data.Id)
	return err
}

func (m *UmsIntegrationChangeHistoryModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
