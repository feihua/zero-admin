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
	pmsProductCategoryFieldNames          = builderx.FieldNames(&PmsProductCategory{})
	pmsProductCategoryRows                = strings.Join(pmsProductCategoryFieldNames, ",")
	pmsProductCategoryRowsExpectAutoSet   = strings.Join(stringx.Remove(pmsProductCategoryFieldNames, "id", "create_time", "update_time"), ",")
	pmsProductCategoryRowsWithPlaceHolder = strings.Join(stringx.Remove(pmsProductCategoryFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	PmsProductCategoryModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PmsProductCategory struct {
		Id           int64  `db:"id"`
		ParentId     int64  `db:"parent_id"` // 上机分类的编号：0表示一级分类
		Name         string `db:"name"`
		Level        int64  `db:"level"` // 分类级别：0->1级；1->2级
		ProductCount int64  `db:"product_count"`
		ProductUnit  string `db:"product_unit"`
		NavStatus    int64  `db:"nav_status"`  // 是否显示在导航栏：0->不显示；1->显示
		ShowStatus   int64  `db:"show_status"` // 显示状态：0->不显示；1->显示
		Sort         int64  `db:"sort"`
		Icon         string `db:"icon"` // 图标
		Keywords     string `db:"keywords"`
		Description  string `db:"description"` // 描述
	}
)

func NewPmsProductCategoryModel(conn sqlx.SqlConn) *PmsProductCategoryModel {
	return &PmsProductCategoryModel{
		conn:  conn,
		table: "pms_product_category",
	}
}

func (m *PmsProductCategoryModel) Insert(data PmsProductCategory) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, pmsProductCategoryRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.ParentId, data.Name, data.Level, data.ProductCount, data.ProductUnit, data.NavStatus, data.ShowStatus, data.Sort, data.Icon, data.Keywords, data.Description)
	return ret, err
}

func (m *PmsProductCategoryModel) FindOne(id int64) (*PmsProductCategory, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", pmsProductCategoryRows, m.table)
	var resp PmsProductCategory
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

func (m *PmsProductCategoryModel) FindByParentId(parentId int64) (*[]PmsProductCategory, error) {

	query := fmt.Sprintf("select %s from %s  where parent_id = ?", pmsProductCategoryRows, m.table)
	var resp []PmsProductCategory
	err := m.conn.QueryRows(&resp, query, parentId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *PmsProductCategoryModel) FindAll(Current int64, PageSize int64) (*[]PmsProductCategory, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", pmsProductCategoryRows, m.table)
	var resp []PmsProductCategory
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

func (m *PmsProductCategoryModel) Count() (int64, error) {
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

func (m *PmsProductCategoryModel) Update(data PmsProductCategory) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, pmsProductCategoryRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.ParentId, data.Name, data.Level, data.ProductCount, data.ProductUnit, data.NavStatus, data.ShowStatus, data.Sort, data.Icon, data.Keywords, data.Description, data.Id)
	return err
}

func (m *PmsProductCategoryModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
