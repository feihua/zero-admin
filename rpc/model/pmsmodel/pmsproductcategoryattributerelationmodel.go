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
	pmsProductCategoryAttributeRelationFieldNames          = builderx.FieldNames(&PmsProductCategoryAttributeRelation{})
	pmsProductCategoryAttributeRelationRows                = strings.Join(pmsProductCategoryAttributeRelationFieldNames, ",")
	pmsProductCategoryAttributeRelationRowsExpectAutoSet   = strings.Join(stringx.Remove(pmsProductCategoryAttributeRelationFieldNames, "id", "create_time", "update_time"), ",")
	pmsProductCategoryAttributeRelationRowsWithPlaceHolder = strings.Join(stringx.Remove(pmsProductCategoryAttributeRelationFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	PmsProductCategoryAttributeRelationModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PmsProductCategoryAttributeRelation struct {
		Id                 int64 `db:"id"`
		ProductCategoryId  int64 `db:"product_category_id"`
		ProductAttributeId int64 `db:"product_attribute_id"`
	}
)

func NewPmsProductCategoryAttributeRelationModel(conn sqlx.SqlConn) *PmsProductCategoryAttributeRelationModel {
	return &PmsProductCategoryAttributeRelationModel{
		conn:  conn,
		table: "pms_product_category_attribute_relation",
	}
}

func (m *PmsProductCategoryAttributeRelationModel) Insert(data PmsProductCategoryAttributeRelation) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, pmsProductCategoryAttributeRelationRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.ProductCategoryId, data.ProductAttributeId)
	return ret, err
}

func (m *PmsProductCategoryAttributeRelationModel) FindOne(id int64) (*PmsProductCategoryAttributeRelation, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", pmsProductCategoryAttributeRelationRows, m.table)
	var resp PmsProductCategoryAttributeRelation
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

func (m *PmsProductCategoryAttributeRelationModel) FindAll(Current int64, PageSize int64) (*[]PmsProductCategoryAttributeRelation, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", pmsProductCategoryAttributeRelationRows, m.table)
	var resp []PmsProductCategoryAttributeRelation
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

func (m *PmsProductCategoryAttributeRelationModel) Count() (int64, error) {
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

func (m *PmsProductCategoryAttributeRelationModel) Update(data PmsProductCategoryAttributeRelation) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, pmsProductCategoryAttributeRelationRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.ProductCategoryId, data.ProductAttributeId, data.Id)
	return err
}

func (m *PmsProductCategoryAttributeRelationModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
