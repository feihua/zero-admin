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
	omsOrderReturnReasonFieldNames          = builderx.FieldNames(&OmsOrderReturnReason{})
	omsOrderReturnReasonRows                = strings.Join(omsOrderReturnReasonFieldNames, ",")
	omsOrderReturnReasonRowsExpectAutoSet   = strings.Join(stringx.Remove(omsOrderReturnReasonFieldNames, "id", "create_time", "update_time"), ",")
	omsOrderReturnReasonRowsWithPlaceHolder = strings.Join(stringx.Remove(omsOrderReturnReasonFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	OmsOrderReturnReasonModel struct {
		conn  sqlx.SqlConn
		table string
	}

	OmsOrderReturnReason struct {
		Id         int64     `db:"id"`
		Name       string    `db:"name"` // 退货类型
		Sort       int64     `db:"sort"`
		Status     int64     `db:"status"`      // 状态：0->不启用；1->启用
		CreateTime time.Time `db:"create_time"` // 添加时间
	}
)

func NewOmsOrderReturnReasonModel(conn sqlx.SqlConn) *OmsOrderReturnReasonModel {
	return &OmsOrderReturnReasonModel{
		conn:  conn,
		table: "oms_order_return_reason",
	}
}

func (m *OmsOrderReturnReasonModel) Insert(data OmsOrderReturnReason) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, omsOrderReturnReasonRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Name, data.Sort, data.Status)
	return ret, err
}

func (m *OmsOrderReturnReasonModel) FindOne(id int64) (*OmsOrderReturnReason, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", omsOrderReturnReasonRows, m.table)
	var resp OmsOrderReturnReason
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

func (m *OmsOrderReturnReasonModel) FindAll(Current int64, PageSize int64) (*[]OmsOrderReturnReason, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", omsOrderReturnReasonRows, m.table)
	var resp []OmsOrderReturnReason
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

func (m *OmsOrderReturnReasonModel) Count() (int64, error) {
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

func (m *OmsOrderReturnReasonModel) Update(data OmsOrderReturnReason) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, omsOrderReturnReasonRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Name, data.Sort, data.Status, data.Id)
	return err
}

func (m *OmsOrderReturnReasonModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
