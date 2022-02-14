package omsmodel

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/builderx"
)

var (
	omsCartItemFieldNames          = builderx.FieldNames(&OmsCartItem{})
	omsCartItemRows                = strings.Join(omsCartItemFieldNames, ",")
	omsCartItemRowsExpectAutoSet   = strings.Join(stringx.Remove(omsCartItemFieldNames, "id", "create_time", "update_time"), ",")
	omsCartItemRowsWithPlaceHolder = strings.Join(stringx.Remove(omsCartItemFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	OmsCartItemModel struct {
		conn  sqlx.SqlConn
		table string
	}

	OmsCartItem struct {
		Id                int64     `db:"id"`
		ProductId         int64     `db:"product_id"`
		ProductSkuId      int64     `db:"product_sku_id"`
		MemberId          int64     `db:"member_id"`
		Quantity          int64     `db:"quantity"`            // 购买数量
		Price             float64   `db:"price"`               // 添加到购物车的价格
		ProductPic        string    `db:"product_pic"`         // 商品主图
		ProductName       string    `db:"product_name"`        // 商品名称
		ProductSubTitle   string    `db:"product_sub_title"`   // 商品副标题（卖点）
		ProductSkuCode    string    `db:"product_sku_code"`    // 商品sku条码
		MemberNickname    string    `db:"member_nickname"`     // 会员昵称
		CreateDate        time.Time `db:"create_date"`         // 创建时间
		ModifyDate        time.Time `db:"modify_date"`         // 修改时间
		DeleteStatus      int64     `db:"delete_status"`       // 是否删除
		ProductCategoryId int64     `db:"product_category_id"` // 商品分类
		ProductBrand      string    `db:"product_brand"`
		ProductSn         string    `db:"product_sn"`
		ProductAttr       string    `db:"product_attr"` // 商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]
	}
)

func NewOmsCartItemModel(conn sqlx.SqlConn) *OmsCartItemModel {
	return &OmsCartItemModel{
		conn:  conn,
		table: "oms_cart_item",
	}
}

func (m *OmsCartItemModel) Insert(data OmsCartItem) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, omsCartItemRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.ProductId, data.ProductSkuId, data.MemberId, data.Quantity, data.Price, data.ProductPic, data.ProductName, data.ProductSubTitle, data.ProductSkuCode, data.MemberNickname, data.CreateDate, data.ModifyDate, data.DeleteStatus, data.ProductCategoryId, data.ProductBrand, data.ProductSn, data.ProductAttr)
	return ret, err
}

func (m *OmsCartItemModel) FindOne(id int64) (*OmsCartItem, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", omsCartItemRows, m.table)
	var resp OmsCartItem
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

func (m *OmsCartItemModel) FindAll(Current int64, PageSize int64) (*[]OmsCartItem, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", omsCartItemRows, m.table)
	var resp []OmsCartItem
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

func (m *OmsCartItemModel) Count() (int64, error) {
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

func (m *OmsCartItemModel) Update(data OmsCartItem) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, omsCartItemRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.ProductId, data.ProductSkuId, data.MemberId, data.Quantity, data.Price, data.ProductPic, data.ProductName, data.ProductSubTitle, data.ProductSkuCode, data.MemberNickname, data.CreateDate, data.ModifyDate, data.DeleteStatus, data.ProductCategoryId, data.ProductBrand, data.ProductSn, data.ProductAttr, data.Id)
	return err
}

func (m *OmsCartItemModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
