package pmsmodel

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
	pmsProductAttributeValueFieldNames          = builderx.FieldNames(&PmsProductAttributeValue{})
	pmsProductAttributeValueRows                = strings.Join(pmsProductAttributeValueFieldNames, ",")
	pmsProductAttributeValueRowsExpectAutoSet   = strings.Join(stringx.Remove(pmsProductAttributeValueFieldNames, "id", "create_time", "update_time"), ",")
	pmsProductAttributeValueRowsWithPlaceHolder = strings.Join(stringx.Remove(pmsProductAttributeValueFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	PmsProductAttributeValueModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PmsProductAttributeValue struct {
		Id                 int64  `db:"id"`
		ProductId          int64  `db:"product_id"`
		ProductAttributeId int64  `db:"product_attribute_id"`
		Value              string `db:"value"` // 手动添加规格或参数的值，参数单值，规格有多个时以逗号隔开
	}
)

func NewPmsProductAttributeValueModel(conn sqlx.SqlConn) *PmsProductAttributeValueModel {
	return &PmsProductAttributeValueModel{
		conn:  conn,
		table: "pms_product_attribute_value",
	}
}

func (m *PmsProductAttributeValueModel) Insert(data PmsProductAttributeValue) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, pmsProductAttributeValueRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.ProductId, data.ProductAttributeId, data.Value)
	return ret, err
}

func (m *PmsProductAttributeValueModel) FindOne(id int64) (*PmsProductAttributeValue, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", pmsProductAttributeValueRows, m.table)
	var resp PmsProductAttributeValue
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

func (m *PmsProductAttributeValueModel) FindAll(Current int64, PageSize int64) (*[]PmsProductAttributeValue, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", pmsProductAttributeValueRows, m.table)
	var resp []PmsProductAttributeValue
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

func (m *PmsProductAttributeValueModel) Count() (int64, error) {
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

func (m *PmsProductAttributeValueModel) Update(data PmsProductAttributeValue) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, pmsProductAttributeValueRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.ProductId, data.ProductAttributeId, data.Value, data.Id)
	return err
}

func (m *PmsProductAttributeValueModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
