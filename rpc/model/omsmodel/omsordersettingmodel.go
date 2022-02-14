package omsmodel

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/builderx"
)

var (
	omsOrderSettingFieldNames          = builderx.FieldNames(&OmsOrderSetting{})
	omsOrderSettingRows                = strings.Join(omsOrderSettingFieldNames, ",")
	omsOrderSettingRowsExpectAutoSet   = strings.Join(stringx.Remove(omsOrderSettingFieldNames, "id", "create_time", "update_time"), ",")
	omsOrderSettingRowsWithPlaceHolder = strings.Join(stringx.Remove(omsOrderSettingFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	OmsOrderSettingModel struct {
		conn  sqlx.SqlConn
		table string
	}

	OmsOrderSetting struct {
		Id                  int64 `db:"id"`
		FlashOrderOvertime  int64 `db:"flash_order_overtime"`  // 秒杀订单超时关闭时间(分)
		NormalOrderOvertime int64 `db:"normal_order_overtime"` // 正常订单超时时间(分)
		ConfirmOvertime     int64 `db:"confirm_overtime"`      // 发货后自动确认收货时间（天）
		FinishOvertime      int64 `db:"finish_overtime"`       // 自动完成交易时间，不能申请售后（天）
		CommentOvertime     int64 `db:"comment_overtime"`      // 订单完成后自动好评时间（天）
	}
)

func NewOmsOrderSettingModel(conn sqlx.SqlConn) *OmsOrderSettingModel {
	return &OmsOrderSettingModel{
		conn:  conn,
		table: "oms_order_setting",
	}
}

func (m *OmsOrderSettingModel) Insert(data OmsOrderSetting) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, omsOrderSettingRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.FlashOrderOvertime, data.NormalOrderOvertime, data.ConfirmOvertime, data.FinishOvertime, data.CommentOvertime)
	return ret, err
}

func (m *OmsOrderSettingModel) FindOne(id int64) (*OmsOrderSetting, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", omsOrderSettingRows, m.table)
	var resp OmsOrderSetting
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

func (m *OmsOrderSettingModel) FindAll(Current int64, PageSize int64) (*[]OmsOrderSetting, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", omsOrderSettingRows, m.table)
	var resp []OmsOrderSetting
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

func (m *OmsOrderSettingModel) Count() (int64, error) {
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

func (m *OmsOrderSettingModel) Update(data OmsOrderSetting) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, omsOrderSettingRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.FlashOrderOvertime, data.NormalOrderOvertime, data.ConfirmOvertime, data.FinishOvertime, data.CommentOvertime, data.Id)
	return err
}

func (m *OmsOrderSettingModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
