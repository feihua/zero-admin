package pmsmodel

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PmsProductModel = (*customPmsProductModel)(nil)

type (
	// PmsProductModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsProductModel.
	PmsProductModel interface {
		pmsProductModel
		ProductListByCategoryId(CategoryId int64, Current int64, PageSize int64) (*[]PmsProduct, error)
		CountByCategoryId(ctx context.Context, categoryId int64) (int64, error)
		ProductListByIds(Ids []int64) (*[]PmsProduct, error)
	}

	customPmsProductModel struct {
		*defaultPmsProductModel
	}
)

// NewPmsProductModel returns a model for the database table.
func NewPmsProductModel(conn sqlx.SqlConn, c cache.CacheConf) PmsProductModel {
	return &customPmsProductModel{
		defaultPmsProductModel: newPmsProductModel(conn, c),
	}

}

func (m *customPmsProductModel) ProductListByCategoryId(CategoryId int64, Current int64, PageSize int64) (*[]PmsProduct, error) {

	query := fmt.Sprintf("select %s from %s where product_category_id = ?  limit ?,?", pmsProductRows, m.table)
	var resp []PmsProduct
	err := m.QueryRowsNoCache(&resp, query, CategoryId, (Current-1)*PageSize, PageSize)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
func (m *customPmsProductModel) CountByCategoryId(ctx context.Context, categoryId int64) (int64, error) {
	query := fmt.Sprintf("select count(*) as count from %s where product_category_id=?", m.table)
	var count int64
	err := m.QueryRowNoCacheCtx(ctx, &count, query, categoryId)
	switch err {
	case nil:
		return count, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}

func (m *customPmsProductModel) ProductListByIds(Ids []int64) (*[]PmsProduct, error) {

	ids, err := json.Marshal(Ids)
	if err != nil {
		panic(err)
	}
	var idsString = string(ids)
	// 最后如果不要中括号，直接trim掉即可
	idsString = strings.Trim(idsString, "[]")

	query := fmt.Sprintf("select %s from %s where id in(?)", pmsProductRows, m.table)

	// 动态填充id
	// query, args, _ := officialSqlx.In("SELECT %s FROM %s WHERE id IN (?)", pmsProductRows, m.table, Ids)
	// if err != nil {
	// 	return
	// }

	// sqlx.In 返回带 `?` bindvar的查询语句, 我们使用Rebind()重新绑定它
	// query = DB.Rebind(query)

	// err = DB.Select(&resp, query, args...)

	var resp []PmsProduct
	err = m.QueryRowsNoCache(&resp, query, idsString)

	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
