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
	pmsSkuStockFieldNames          = builderx.FieldNames(&PmsSkuStock{})
	pmsSkuStockRows                = strings.Join(pmsSkuStockFieldNames, ",")
	pmsSkuStockRowsExpectAutoSet   = strings.Join(stringx.Remove(pmsSkuStockFieldNames, "id", "create_time", "update_time"), ",")
	pmsSkuStockRowsWithPlaceHolder = strings.Join(stringx.Remove(pmsSkuStockFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	PmsSkuStockModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PmsSkuStock struct {
		Id             int64   `db:"id"`
		ProductId      int64   `db:"product_id"`
		SkuCode        string  `db:"sku_code"` // sku编码
		Price          float64 `db:"price"`
		Stock          int64   `db:"stock"`           // 库存
		LowStock       int64   `db:"low_stock"`       // 预警库存
		Pic            string  `db:"pic"`             // 展示图片
		Sale           int64   `db:"sale"`            // 销量
		PromotionPrice float64 `db:"promotion_price"` // 单品促销价格
		LockStock      int64   `db:"lock_stock"`      // 锁定库存
		SpData         string  `db:"sp_data"`         // 商品销售属性，json格式
	}
)

func NewPmsSkuStockModel(conn sqlx.SqlConn) *PmsSkuStockModel {
	return &PmsSkuStockModel{
		conn:  conn,
		table: "pms_sku_stock",
	}
}

func (m *PmsSkuStockModel) Insert(data PmsSkuStock) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, pmsSkuStockRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.ProductId, data.SkuCode, data.Price, data.Stock, data.LowStock, data.Pic, data.Sale, data.PromotionPrice, data.LockStock, data.SpData)
	return ret, err
}

func (m *PmsSkuStockModel) FindOne(id int64) (*PmsSkuStock, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", pmsSkuStockRows, m.table)
	var resp PmsSkuStock
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

func (m *PmsSkuStockModel) FindAll(Current int64, PageSize int64) (*[]PmsSkuStock, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", pmsSkuStockRows, m.table)
	var resp []PmsSkuStock
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

func (m *PmsSkuStockModel) Count() (int64, error) {
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

func (m *PmsSkuStockModel) Update(data PmsSkuStock) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, pmsSkuStockRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.ProductId, data.SkuCode, data.Price, data.Stock, data.LowStock, data.Pic, data.Sale, data.PromotionPrice, data.LockStock, data.SpData, data.Id)
	return err
}

func (m *PmsSkuStockModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
