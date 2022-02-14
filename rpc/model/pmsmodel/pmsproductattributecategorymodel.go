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
	pmsProductAttributeCategoryFieldNames          = builderx.FieldNames(&PmsProductAttributeCategory{})
	pmsProductAttributeCategoryRows                = strings.Join(pmsProductAttributeCategoryFieldNames, ",")
	pmsProductAttributeCategoryRowsExpectAutoSet   = strings.Join(stringx.Remove(pmsProductAttributeCategoryFieldNames, "id", "create_time", "update_time"), ",")
	pmsProductAttributeCategoryRowsWithPlaceHolder = strings.Join(stringx.Remove(pmsProductAttributeCategoryFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	PmsProductAttributeCategoryModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PmsProductAttributeCategory struct {
		Id             int64  `db:"id"`
		Name           string `db:"name"`
		AttributeCount int64  `db:"attribute_count"` // 属性数量
		ParamCount     int64  `db:"param_count"`     // 参数数量
	}
)

func NewPmsProductAttributeCategoryModel(conn sqlx.SqlConn) *PmsProductAttributeCategoryModel {
	return &PmsProductAttributeCategoryModel{
		conn:  conn,
		table: "pms_product_attribute_category",
	}
}

func (m *PmsProductAttributeCategoryModel) Insert(data PmsProductAttributeCategory) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, pmsProductAttributeCategoryRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Name, data.AttributeCount, data.ParamCount)
	return ret, err
}

func (m *PmsProductAttributeCategoryModel) FindOne(id int64) (*PmsProductAttributeCategory, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", pmsProductAttributeCategoryRows, m.table)
	var resp PmsProductAttributeCategory
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

func (m *PmsProductAttributeCategoryModel) FindAll(Current int64, PageSize int64) (*[]PmsProductAttributeCategory, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", pmsProductAttributeCategoryRows, m.table)
	var resp []PmsProductAttributeCategory
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

func (m *PmsProductAttributeCategoryModel) Count() (int64, error) {
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

func (m *PmsProductAttributeCategoryModel) Update(data PmsProductAttributeCategory) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, pmsProductAttributeCategoryRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Name, data.AttributeCount, data.ParamCount, data.Id)
	return err
}

func (m *PmsProductAttributeCategoryModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
