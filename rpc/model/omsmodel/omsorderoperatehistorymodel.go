package omsmodel

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
	omsOrderOperateHistoryFieldNames          = builderx.FieldNames(&OmsOrderOperateHistory{})
	omsOrderOperateHistoryRows                = strings.Join(omsOrderOperateHistoryFieldNames, ",")
	omsOrderOperateHistoryRowsExpectAutoSet   = strings.Join(stringx.Remove(omsOrderOperateHistoryFieldNames, "id", "create_time", "update_time"), ",")
	omsOrderOperateHistoryRowsWithPlaceHolder = strings.Join(stringx.Remove(omsOrderOperateHistoryFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	OmsOrderOperateHistoryModel struct {
		conn  sqlx.SqlConn
		table string
	}

	OmsOrderOperateHistory struct {
		Id          int64     `db:"id"`
		OrderId     int64     `db:"order_id"`     // 订单id
		OperateMan  string    `db:"operate_man"`  // 操作人：用户；系统；后台管理员
		CreateTime  time.Time `db:"create_time"`  // 操作时间
		OrderStatus int64     `db:"order_status"` // 订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单
		Note        string    `db:"note"`         // 备注
	}
)

func NewOmsOrderOperateHistoryModel(conn sqlx.SqlConn) *OmsOrderOperateHistoryModel {
	return &OmsOrderOperateHistoryModel{
		conn:  conn,
		table: "oms_order_operate_history",
	}
}

func (m *OmsOrderOperateHistoryModel) Insert(data OmsOrderOperateHistory) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, omsOrderOperateHistoryRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.OrderId, data.OperateMan, data.OrderStatus, data.Note)
	return ret, err
}

func (m *OmsOrderOperateHistoryModel) FindOne(id int64) (*OmsOrderOperateHistory, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", omsOrderOperateHistoryRows, m.table)
	var resp OmsOrderOperateHistory
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

func (m *OmsOrderOperateHistoryModel) FindAll(Current int64, PageSize int64) (*[]OmsOrderOperateHistory, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", omsOrderOperateHistoryRows, m.table)
	var resp []OmsOrderOperateHistory
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

func (m *OmsOrderOperateHistoryModel) Count() (int64, error) {
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

func (m *OmsOrderOperateHistoryModel) Update(data OmsOrderOperateHistory) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, omsOrderOperateHistoryRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.OrderId, data.OperateMan, data.OrderStatus, data.Note, data.Id)
	return err
}

func (m *OmsOrderOperateHistoryModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
