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
	omsCompanyAddressFieldNames          = builderx.FieldNames(&OmsCompanyAddress{})
	omsCompanyAddressRows                = strings.Join(omsCompanyAddressFieldNames, ",")
	omsCompanyAddressRowsExpectAutoSet   = strings.Join(stringx.Remove(omsCompanyAddressFieldNames, "id", "create_time", "update_time"), ",")
	omsCompanyAddressRowsWithPlaceHolder = strings.Join(stringx.Remove(omsCompanyAddressFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	OmsCompanyAddressModel struct {
		conn  sqlx.SqlConn
		table string
	}

	OmsCompanyAddress struct {
		Id            int64  `db:"id"`
		AddressName   string `db:"address_name"`   // 地址名称
		SendStatus    int64  `db:"send_status"`    // 默认发货地址：0->否；1->是
		ReceiveStatus int64  `db:"receive_status"` // 是否默认收货地址：0->否；1->是
		Name          string `db:"name"`           // 收发货人姓名
		Phone         string `db:"phone"`          // 收货人电话
		Province      string `db:"province"`       // 省/直辖市
		City          string `db:"city"`           // 市
		Region        string `db:"region"`         // 区
		DetailAddress string `db:"detail_address"` // 详细地址
	}
)

func NewOmsCompanyAddressModel(conn sqlx.SqlConn) *OmsCompanyAddressModel {
	return &OmsCompanyAddressModel{
		conn:  conn,
		table: "oms_company_address",
	}
}

func (m *OmsCompanyAddressModel) Insert(data OmsCompanyAddress) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, omsCompanyAddressRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.AddressName, data.SendStatus, data.ReceiveStatus, data.Name, data.Phone, data.Province, data.City, data.Region, data.DetailAddress)
	return ret, err
}

func (m *OmsCompanyAddressModel) FindOne(id int64) (*OmsCompanyAddress, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", omsCompanyAddressRows, m.table)
	var resp OmsCompanyAddress
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

func (m *OmsCompanyAddressModel) FindAll(Current int64, PageSize int64) (*[]OmsCompanyAddress, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", omsCompanyAddressRows, m.table)
	var resp []OmsCompanyAddress
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

func (m *OmsCompanyAddressModel) Count() (int64, error) {
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

func (m *OmsCompanyAddressModel) Update(data OmsCompanyAddress) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, omsCompanyAddressRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.AddressName, data.SendStatus, data.ReceiveStatus, data.Name, data.Phone, data.Province, data.City, data.Region, data.DetailAddress, data.Id)
	return err
}

func (m *OmsCompanyAddressModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
