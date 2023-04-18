package pmsmodel

import (
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PmsProductCategoryModel = (*customPmsProductCategoryModel)(nil)

type (
	// PmsProductCategoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsProductCategoryModel.
	PmsProductCategoryModel interface {
		pmsProductCategoryModel
		FindByParentId(parentId int64) (*[]PmsProductCategory, error)
		FindAll(Current int64, PageSize int64) (*[]PmsProductCategory, error)
		Count() (int64, error)
	}

	customPmsProductCategoryModel struct {
		*defaultPmsProductCategoryModel
	}
)

// NewPmsProductCategoryModel returns a model for the database table.
func NewPmsProductCategoryModel(conn sqlx.SqlConn, c cache.CacheConf) PmsProductCategoryModel {
	return &customPmsProductCategoryModel{
		defaultPmsProductCategoryModel: newPmsProductCategoryModel(conn, c),
	}
}

func (m *customPmsProductCategoryModel) FindByParentId(parentId int64) (*[]PmsProductCategory, error) {

	query := fmt.Sprintf("select %s from %s  where parent_id = ?", pmsProductCategoryRows, m.table)
	var resp []PmsProductCategory
	err := m.QueryRowsNoCache(&resp, query, parentId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customPmsProductCategoryModel) FindAll(Current int64, PageSize int64) (*[]PmsProductCategory, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", pmsProductCategoryRows, m.table)
	var resp []PmsProductCategory
	err := m.QueryRowsNoCache(&resp, query, (Current-1)*PageSize, PageSize)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customPmsProductCategoryModel) Count() (int64, error) {
	query := fmt.Sprintf("select count(*) as count from %s", m.table)

	var count int64
	err := m.QueryRowNoCache(&count, query)

	switch err {
	case nil:
		return count, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}
