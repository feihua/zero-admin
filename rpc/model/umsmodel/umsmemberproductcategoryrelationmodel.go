package umsmodel

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	umsMemberProductCategoryRelationFieldNames          = builderx.FieldNames(&UmsMemberProductCategoryRelation{})
	umsMemberProductCategoryRelationRows                = strings.Join(umsMemberProductCategoryRelationFieldNames, ",")
	umsMemberProductCategoryRelationRowsExpectAutoSet   = strings.Join(stringx.Remove(umsMemberProductCategoryRelationFieldNames, "id", "create_time", "update_time"), ",")
	umsMemberProductCategoryRelationRowsWithPlaceHolder = strings.Join(stringx.Remove(umsMemberProductCategoryRelationFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	UmsMemberProductCategoryRelationModel struct {
		conn  sqlx.SqlConn
		table string
	}

	UmsMemberProductCategoryRelation struct {
		Id                int64 `db:"id"`
		MemberId          int64 `db:"member_id"`
		ProductCategoryId int64 `db:"product_category_id"`
	}
)

func NewUmsMemberProductCategoryRelationModel(conn sqlx.SqlConn) *UmsMemberProductCategoryRelationModel {
	return &UmsMemberProductCategoryRelationModel{
		conn:  conn,
		table: "ums_member_product_category_relation",
	}
}

func (m *UmsMemberProductCategoryRelationModel) Insert(data UmsMemberProductCategoryRelation) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, umsMemberProductCategoryRelationRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.MemberId, data.ProductCategoryId)
	return ret, err
}

func (m *UmsMemberProductCategoryRelationModel) FindOne(id int64) (*UmsMemberProductCategoryRelation, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", umsMemberProductCategoryRelationRows, m.table)
	var resp UmsMemberProductCategoryRelation
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

func (m *UmsMemberProductCategoryRelationModel) FindAll(Current int64, PageSize int64) (*[]UmsMemberProductCategoryRelation, error) {

	if Current < 1 {
		Current = 1
	}
	if PageSize < 1 {
		PageSize = 20
	}
	query := fmt.Sprintf("select %s from %s limit ?,?", umsMemberProductCategoryRelationRows, m.table)
	var resp []UmsMemberProductCategoryRelation
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

func (m *UmsMemberProductCategoryRelationModel) Update(data UmsMemberProductCategoryRelation) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, umsMemberProductCategoryRelationRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.MemberId, data.ProductCategoryId, data.Id)
	return err
}

func (m *UmsMemberProductCategoryRelationModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
