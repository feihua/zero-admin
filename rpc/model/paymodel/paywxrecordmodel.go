package paymodel

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	payWxRecordFieldNames          = builderx.FieldNames(&PayWxRecord{})
	payWxRecordRows                = strings.Join(payWxRecordFieldNames, ",")
	payWxRecordRowsExpectAutoSet   = strings.Join(stringx.Remove(payWxRecordFieldNames, "id", "create_time", "update_time"), ",")
	payWxRecordRowsWithPlaceHolder = strings.Join(stringx.Remove(payWxRecordFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	PayWxRecordModel interface {
		Insert(data PayWxRecord) (sql.Result, error)
		FindOne(id int64) (*PayWxRecord, error)
		Update(data PayWxRecord) error
		Delete(id int64) error
	}

	defaultPayWxRecordModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PayWxRecord struct {
		Id         int64     `db:"id"`          // 主键
		BusinessId string    `db:"businessId"`  // 业务编号
		Amount     string    `db:"amount"`      // 金额
		PayType    int64     `db:"pay_type"`    // 支付类型(1:app支付 2:小程序支付 3:h5支付 4:公众号支付)
		Remarks    string    `db:"remarks"`     // 备注
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
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, payWxRecordRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.BusinessId, data.Amount, data.PayType, data.Remarks)
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

func (m *defaultPayWxRecordModel) Update(data PayWxRecord) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, payWxRecordRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.BusinessId, data.Amount, data.PayType, data.Remarks, data.Id)
	return err
}

func (m *defaultPayWxRecordModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
