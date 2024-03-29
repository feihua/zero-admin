// Code generated by goctl. DO NOT EDIT.

package pmsmodel

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	pmsSkuStockFieldNames          = builder.RawFieldNames(&PmsSkuStock{})
	pmsSkuStockRows                = strings.Join(pmsSkuStockFieldNames, ",")
	pmsSkuStockRowsExpectAutoSet   = strings.Join(stringx.Remove(pmsSkuStockFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	pmsSkuStockRowsWithPlaceHolder = strings.Join(stringx.Remove(pmsSkuStockFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	pmsSkuStockModel interface {
		Insert(ctx context.Context, data *PmsSkuStock) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*PmsSkuStock, error)
		Update(ctx context.Context, data *PmsSkuStock) error
		Delete(ctx context.Context, id int64) error
	}

	defaultPmsSkuStockModel struct {
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

func newPmsSkuStockModel(conn sqlx.SqlConn) *defaultPmsSkuStockModel {
	return &defaultPmsSkuStockModel{
		conn:  conn,
		table: "`pms_sku_stock`",
	}
}

func (m *defaultPmsSkuStockModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultPmsSkuStockModel) FindOne(ctx context.Context, id int64) (*PmsSkuStock, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", pmsSkuStockRows, m.table)
	var resp PmsSkuStock
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultPmsSkuStockModel) Insert(ctx context.Context, data *PmsSkuStock) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, pmsSkuStockRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.ProductId, data.SkuCode, data.Price, data.Stock, data.LowStock, data.Pic, data.Sale, data.PromotionPrice, data.LockStock, data.SpData)
	return ret, err
}

func (m *defaultPmsSkuStockModel) Update(ctx context.Context, data *PmsSkuStock) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, pmsSkuStockRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.ProductId, data.SkuCode, data.Price, data.Stock, data.LowStock, data.Pic, data.Sale, data.PromotionPrice, data.LockStock, data.SpData, data.Id)
	return err
}

func (m *defaultPmsSkuStockModel) tableName() string {
	return m.table
}
