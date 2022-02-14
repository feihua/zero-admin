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
	omsOrderReturnApplyFieldNames          = builderx.FieldNames(&OmsOrderReturnApply{})
	omsOrderReturnApplyRows                = strings.Join(omsOrderReturnApplyFieldNames, ",")
	omsOrderReturnApplyRowsExpectAutoSet   = strings.Join(stringx.Remove(omsOrderReturnApplyFieldNames, "id", "create_time", "update_time"), ",")
	omsOrderReturnApplyRowsWithPlaceHolder = strings.Join(stringx.Remove(omsOrderReturnApplyFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	OmsOrderReturnApplyModel struct {
		conn  sqlx.SqlConn
		table string
	}

	OmsOrderReturnApply struct {
		Id               int64     `db:"id"`
		OrderId          int64     `db:"order_id"`           // 订单id
		CompanyAddressId int64     `db:"company_address_id"` // 收货地址表id
		ProductId        int64     `db:"product_id"`         // 退货商品id
		OrderSn          string    `db:"order_sn"`           // 订单编号
		CreateTime       time.Time `db:"create_time"`        // 申请时间
		MemberUsername   string    `db:"member_username"`    // 会员用户名
		ReturnAmount     float64   `db:"return_amount"`      // 退款金额
		ReturnName       string    `db:"return_name"`        // 退货人姓名
		ReturnPhone      string    `db:"return_phone"`       // 退货人电话
		Status           int64     `db:"status"`             // 申请状态：0->待处理；1->退货中；2->已完成；3->已拒绝
		HandleTime       time.Time `db:"handle_time"`        // 处理时间
		ProductPic       string    `db:"product_pic"`        // 商品图片
		ProductName      string    `db:"product_name"`       // 商品名称
		ProductBrand     string    `db:"product_brand"`      // 商品品牌
		ProductAttr      string    `db:"product_attr"`       // 商品销售属性：颜色：红色；尺码：xl;
		ProductCount     int64     `db:"product_count"`      // 退货数量
		ProductPrice     float64   `db:"product_price"`      // 商品单价
		ProductRealPrice float64   `db:"product_real_price"` // 商品实际支付单价
		Reason           string    `db:"reason"`             // 原因
		Description      string    `db:"description"`        // 描述
		ProofPics        string    `db:"proof_pics"`         // 凭证图片，以逗号隔开
		HandleNote       string    `db:"handle_note"`        // 处理备注
		HandleMan        string    `db:"handle_man"`         // 处理人员
		ReceiveMan       string    `db:"receive_man"`        // 收货人
		ReceiveTime      time.Time `db:"receive_time"`       // 收货时间
		ReceiveNote      string    `db:"receive_note"`       // 收货备注
	}
)

func NewOmsOrderReturnApplyModel(conn sqlx.SqlConn) *OmsOrderReturnApplyModel {
	return &OmsOrderReturnApplyModel{
		conn:  conn,
		table: "oms_order_return_apply",
	}
}

func (m *OmsOrderReturnApplyModel) Insert(data OmsOrderReturnApply) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, omsOrderReturnApplyRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.OrderId, data.CompanyAddressId, data.ProductId, data.OrderSn, data.MemberUsername, data.ReturnAmount, data.ReturnName, data.ReturnPhone, data.Status, data.HandleTime, data.ProductPic, data.ProductName, data.ProductBrand, data.ProductAttr, data.ProductCount, data.ProductPrice, data.ProductRealPrice, data.Reason, data.Description, data.ProofPics, data.HandleNote, data.HandleMan, data.ReceiveMan, data.ReceiveTime, data.ReceiveNote)
	return ret, err
}

func (m *OmsOrderReturnApplyModel) FindOne(id int64) (*OmsOrderReturnApply, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", omsOrderReturnApplyRows, m.table)
	var resp OmsOrderReturnApply
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

func (m *OmsOrderReturnApplyModel) FindAll(Current int64, PageSize int64) (*[]OmsOrderReturnApply, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", omsOrderReturnApplyRows, m.table)
	var resp []OmsOrderReturnApply
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

func (m *OmsOrderReturnApplyModel) Count() (int64, error) {
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

func (m *OmsOrderReturnApplyModel) Update(data OmsOrderReturnApply) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, omsOrderReturnApplyRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.OrderId, data.CompanyAddressId, data.ProductId, data.OrderSn, data.MemberUsername, data.ReturnAmount, data.ReturnName, data.ReturnPhone, data.Status, data.HandleTime, data.ProductPic, data.ProductName, data.ProductBrand, data.ProductAttr, data.ProductCount, data.ProductPrice, data.ProductRealPrice, data.Reason, data.Description, data.ProofPics, data.HandleNote, data.HandleMan, data.ReceiveMan, data.ReceiveTime, data.ReceiveNote, data.Id)
	return err
}

func (m *OmsOrderReturnApplyModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
