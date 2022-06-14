package umsmodel

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
	umsMemberReceiveAddressFieldNames          = builderx.FieldNames(&UmsMemberReceiveAddress{})
	umsMemberReceiveAddressRows                = strings.Join(umsMemberReceiveAddressFieldNames, ",")
	umsMemberReceiveAddressRowsExpectAutoSet   = strings.Join(stringx.Remove(umsMemberReceiveAddressFieldNames, "id", "create_time", "update_time"), ",")
	umsMemberReceiveAddressRowsWithPlaceHolder = strings.Join(stringx.Remove(umsMemberReceiveAddressFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	UmsMemberReceiveAddressModel struct {
		conn  sqlx.SqlConn
		table string
	}

	UmsMemberReceiveAddress struct {
		Id            int64  `db:"id"`
		MemberId      int64  `db:"member_id"`
		Name          string `db:"name"` // 收货人名称
		PhoneNumber   string `db:"phone_number"`
		DefaultStatus int64  `db:"default_status"` // 是否为默认
		PostCode      string `db:"post_code"`      // 邮政编码
		Province      string `db:"province"`       // 省份/直辖市
		City          string `db:"city"`           // 城市
		Region        string `db:"region"`         // 区
		DetailAddress string `db:"detail_address"` // 详细地址(街道)
	}
)

func NewUmsMemberReceiveAddressModel(conn sqlx.SqlConn) *UmsMemberReceiveAddressModel {
	return &UmsMemberReceiveAddressModel{
		conn:  conn,
		table: "ums_member_receive_address",
	}
}

func (m *UmsMemberReceiveAddressModel) Insert(data UmsMemberReceiveAddress) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, umsMemberReceiveAddressRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.MemberId, data.Name, data.PhoneNumber, data.DefaultStatus, data.PostCode, data.Province, data.City, data.Region, data.DetailAddress)
	return ret, err
}

func (m *UmsMemberReceiveAddressModel) FindOne(id int64) (*UmsMemberReceiveAddress, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", umsMemberReceiveAddressRows, m.table)
	var resp UmsMemberReceiveAddress
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

func (m *UmsMemberReceiveAddressModel) FindByIdAndMemberId(id int64, memberId int64) (*UmsMemberReceiveAddress, error) {
	query := fmt.Sprintf("select %s from %s where id = ? and member_id = ? limit 1", umsMemberReceiveAddressRows, m.table)
	var resp UmsMemberReceiveAddress
	err := m.conn.QueryRow(&resp, query, id, memberId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *UmsMemberReceiveAddressModel) FindAll(Current int64, PageSize int64) (*[]UmsMemberReceiveAddress, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", umsMemberReceiveAddressRows, m.table)
	var resp []UmsMemberReceiveAddress
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

func (m *UmsMemberReceiveAddressModel) Count() (int64, error) {
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

func (m *UmsMemberReceiveAddressModel) Update(data UmsMemberReceiveAddress) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, umsMemberReceiveAddressRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.MemberId, data.Name, data.PhoneNumber, data.DefaultStatus, data.PostCode, data.Province, data.City, data.Region, data.DetailAddress, data.Id)
	return err
}

func (m *UmsMemberReceiveAddressModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
