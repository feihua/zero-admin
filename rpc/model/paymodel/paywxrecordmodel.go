package paymodel

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
	payWxRecordFieldNames          = builderx.FieldNames(&PayWxRecord{})
	payWxRecordRows                = strings.Join(payWxRecordFieldNames, ",")
	payWxRecordRowsExpectAutoSet   = strings.Join(stringx.Remove(payWxRecordFieldNames, "id", "create_time", "update_time"), ",")
	payWxRecordRowsWithPlaceHolder = strings.Join(stringx.Remove(payWxRecordFieldNames, "id", "pay_type", "create_time", "update_time"), "=?,") + "=?"
)

type (
	PayWxRecordModel interface {
		Insert(data PayWxRecord) (sql.Result, error)
		FindOne(id int64) (*PayWxRecord, error)
		FindOneByBusinessId(businessId string) (*PayWxRecord, error)
		Update(data PayWxRecord) error
		Delete(id int64) error
	}

	defaultPayWxRecordModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PayWxRecord struct {
		Id         int64     `db:"id"`         // 主键
		BusinessId string    `db:"businessId"` // 业务编号
		Amount     string    `db:"amount"`     // 金额
		PayType    string    `db:"pay_type"`   // 支付类型(APP:APP支付 JSAPI:小程序,公众号 MWEB:H5支付)
		Remarks    string    `db:"remarks"`    // 备注
		ReturnCode string    `db:"return_code"`
		ReturnMsg  string    `db:"return_msg"`
		ResultCode string    `db:"result_code"`
		ResultMsg  string    `db:"result_msg"`
		PayStatus  int64     `db:"pay_status"`  // 0：初始化 1：已发送 2：成功 3：失败
		CreateTime time.Time `db:"create_time"` // 创建时间
		UpdateTime time.Time `db:"update_time"` // 更新时间
	}
)

func NewPayWxRecordModel(conn sqlx.SqlConn) PayWxRecordModel {
	return &defaultPayWxRecordModel{
		conn:  conn,
		table: "pay_wx_record",
	}
}

func (m *defaultPayWxRecordModel) Insert(data PayWxRecord) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, payWxRecordRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.BusinessId, data.Amount, data.PayType, data.Remarks, data.ReturnCode, data.ReturnMsg, data.ResultCode, data.ResultMsg, data.PayStatus)
	return ret, err
}

func (m *defaultPayWxRecordModel) FindOne(id int64) (*PayWxRecord, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", payWxRecordRows, m.table)
	var resp PayWxRecord
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

func (m *defaultPayWxRecordModel) FindOneByBusinessId(businessId string) (*PayWxRecord, error) {
	query := fmt.Sprintf("select %s from %s where businessId = ? limit 1", payWxRecordRows, m.table)
	var resp PayWxRecord
	err := m.conn.QueryRow(&resp, query, businessId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultPayWxRecordModel) Update(data PayWxRecord) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, payWxRecordRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.BusinessId, data.Amount, data.Remarks, data.ReturnCode, data.ReturnMsg, data.ResultCode, data.ResultMsg, data.PayStatus, data.Id)
	return err
}

func (m *defaultPayWxRecordModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
