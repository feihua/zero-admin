package omsmodel

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
	omsOrderItemFieldNames          = builderx.FieldNames(&OmsOrderItem{})
	omsOrderItemRows                = strings.Join(omsOrderItemFieldNames, ",")
	omsOrderItemRowsExpectAutoSet   = strings.Join(stringx.Remove(omsOrderItemFieldNames, "id", "create_time", "update_time"), ",")
	omsOrderItemRowsWithPlaceHolder = strings.Join(stringx.Remove(omsOrderItemFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	OmsOrderItemModel struct {
		conn  sqlx.SqlConn
		table string
	}

	OmsOrderItem struct {
		Id                int64   `db:"id"`
		OrderId           int64   `db:"order_id"` // 订单id
		OrderSn           string  `db:"order_sn"` // 订单编号
		ProductId         int64   `db:"product_id"`
		ProductPic        string  `db:"product_pic"`
		ProductName       string  `db:"product_name"`
		ProductBrand      string  `db:"product_brand"`
		ProductSn         string  `db:"product_sn"`
		ProductPrice      float64 `db:"product_price"`       // 销售价格
		ProductQuantity   int64   `db:"product_quantity"`    // 购买数量
		ProductSkuId      int64   `db:"product_sku_id"`      // 商品sku编号
		ProductSkuCode    string  `db:"product_sku_code"`    // 商品sku条码
		ProductCategoryId int64   `db:"product_category_id"` // 商品分类id
		PromotionName     string  `db:"promotion_name"`      // 商品促销名称
		PromotionAmount   float64 `db:"promotion_amount"`    // 商品促销分解金额
		CouponAmount      float64 `db:"coupon_amount"`       // 优惠券优惠分解金额
		IntegrationAmount float64 `db:"integration_amount"`  // 积分优惠分解金额
		RealAmount        float64 `db:"real_amount"`         // 该商品经过优惠后的分解金额
		GiftIntegration   int64   `db:"gift_integration"`
		GiftGrowth        int64   `db:"gift_growth"`
		ProductAttr       string  `db:"product_attr"` // 商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]
	}
)

func NewOmsOrderItemModel(conn sqlx.SqlConn) *OmsOrderItemModel {
	return &OmsOrderItemModel{
		conn:  conn,
		table: "oms_order_item",
	}
}

func (m *OmsOrderItemModel) Insert(data OmsOrderItem) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, omsOrderItemRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.OrderId, data.OrderSn, data.ProductId, data.ProductPic, data.ProductName, data.ProductBrand, data.ProductSn, data.ProductPrice, data.ProductQuantity, data.ProductSkuId, data.ProductSkuCode, data.ProductCategoryId, data.PromotionName, data.PromotionAmount, data.CouponAmount, data.IntegrationAmount, data.RealAmount, data.GiftIntegration, data.GiftGrowth, data.ProductAttr)
	return ret, err
}

func (m *OmsOrderItemModel) FindOne(id int64) (*OmsOrderItem, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", omsOrderItemRows, m.table)
	var resp OmsOrderItem
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

func (m *OmsOrderItemModel) FindAll(Current int64, PageSize int64) (*[]OmsOrderItem, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", omsOrderItemRows, m.table)
	var resp []OmsOrderItem
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

func (m *OmsOrderItemModel) Count() (int64, error) {
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

func (m *OmsOrderItemModel) Update(data OmsOrderItem) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, omsOrderItemRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.OrderId, data.OrderSn, data.ProductId, data.ProductPic, data.ProductName, data.ProductBrand, data.ProductSn, data.ProductPrice, data.ProductQuantity, data.ProductSkuId, data.ProductSkuCode, data.ProductCategoryId, data.PromotionName, data.PromotionAmount, data.CouponAmount, data.IntegrationAmount, data.RealAmount, data.GiftIntegration, data.GiftGrowth, data.ProductAttr, data.Id)
	return err
}

func (m *OmsOrderItemModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
