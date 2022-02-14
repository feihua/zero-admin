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
	omsOrderFieldNames          = builderx.FieldNames(&OmsOrder{})
	omsOrderRows                = strings.Join(omsOrderFieldNames, ",")
	omsOrderRowsExpectAutoSet   = strings.Join(stringx.Remove(omsOrderFieldNames, "id", "create_time", "update_time"), ",")
	omsOrderRowsWithPlaceHolder = strings.Join(stringx.Remove(omsOrderFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	OmsOrderModel struct {
		conn  sqlx.SqlConn
		table string
	}

	OmsOrder struct {
		Id                    int64     `db:"id"` // 订单id
		MemberId              int64     `db:"member_id"`
		CouponId              int64     `db:"coupon_id"`
		OrderSn               string    `db:"order_sn"`                // 订单编号
		CreateTime            time.Time `db:"create_time"`             // 提交时间
		MemberUsername        string    `db:"member_username"`         // 用户帐号
		TotalAmount           float64   `db:"total_amount"`            // 订单总金额
		PayAmount             float64   `db:"pay_amount"`              // 应付金额（实际支付金额）
		FreightAmount         float64   `db:"freight_amount"`          // 运费金额
		PromotionAmount       float64   `db:"promotion_amount"`        // 促销优化金额（促销价、满减、阶梯价）
		IntegrationAmount     float64   `db:"integration_amount"`      // 积分抵扣金额
		CouponAmount          float64   `db:"coupon_amount"`           // 优惠券抵扣金额
		DiscountAmount        float64   `db:"discount_amount"`         // 管理员后台调整订单使用的折扣金额
		PayType               int64     `db:"pay_type"`                // 支付方式：0->未支付；1->支付宝；2->微信
		SourceType            int64     `db:"source_type"`             // 订单来源：0->PC订单；1->app订单
		Status                int64     `db:"status"`                  // 订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单
		OrderType             int64     `db:"order_type"`              // 订单类型：0->正常订单；1->秒杀订单
		DeliveryCompany       string    `db:"delivery_company"`        // 物流公司(配送方式)
		DeliverySn            string    `db:"delivery_sn"`             // 物流单号
		AutoConfirmDay        int64     `db:"auto_confirm_day"`        // 自动确认时间（天）
		Integration           int64     `db:"integration"`             // 可以获得的积分
		Growth                int64     `db:"growth"`                  // 可以活动的成长值
		PromotionInfo         string    `db:"promotion_info"`          // 活动信息
		BillType              int64     `db:"bill_type"`               // 发票类型：0->不开发票；1->电子发票；2->纸质发票
		BillHeader            string    `db:"bill_header"`             // 发票抬头
		BillContent           string    `db:"bill_content"`            // 发票内容
		BillReceiverPhone     string    `db:"bill_receiver_phone"`     // 收票人电话
		BillReceiverEmail     string    `db:"bill_receiver_email"`     // 收票人邮箱
		ReceiverName          string    `db:"receiver_name"`           // 收货人姓名
		ReceiverPhone         string    `db:"receiver_phone"`          // 收货人电话
		ReceiverPostCode      string    `db:"receiver_post_code"`      // 收货人邮编
		ReceiverProvince      string    `db:"receiver_province"`       // 省份/直辖市
		ReceiverCity          string    `db:"receiver_city"`           // 城市
		ReceiverRegion        string    `db:"receiver_region"`         // 区
		ReceiverDetailAddress string    `db:"receiver_detail_address"` // 详细地址
		Note                  string    `db:"note"`                    // 订单备注
		ConfirmStatus         int64     `db:"confirm_status"`          // 确认收货状态：0->未确认；1->已确认
		DeleteStatus          int64     `db:"delete_status"`           // 删除状态：0->未删除；1->已删除
		UseIntegration        int64     `db:"use_integration"`         // 下单时使用的积分
		PaymentTime           time.Time `db:"payment_time"`            // 支付时间
		DeliveryTime          time.Time `db:"delivery_time"`           // 发货时间
		ReceiveTime           time.Time `db:"receive_time"`            // 确认收货时间
		CommentTime           time.Time `db:"comment_time"`            // 评价时间
		ModifyTime            time.Time `db:"modify_time"`             // 修改时间
	}
)

func NewOmsOrderModel(conn sqlx.SqlConn) *OmsOrderModel {
	return &OmsOrderModel{
		conn:  conn,
		table: "oms_order",
	}
}

func (m *OmsOrderModel) Insert(data OmsOrder) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, omsOrderRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.MemberId, data.CouponId, data.OrderSn, data.MemberUsername, data.TotalAmount, data.PayAmount, data.FreightAmount, data.PromotionAmount, data.IntegrationAmount, data.CouponAmount, data.DiscountAmount, data.PayType, data.SourceType, data.Status, data.OrderType, data.DeliveryCompany, data.DeliverySn, data.AutoConfirmDay, data.Integration, data.Growth, data.PromotionInfo, data.BillType, data.BillHeader, data.BillContent, data.BillReceiverPhone, data.BillReceiverEmail, data.ReceiverName, data.ReceiverPhone, data.ReceiverPostCode, data.ReceiverProvince, data.ReceiverCity, data.ReceiverRegion, data.ReceiverDetailAddress, data.Note, data.ConfirmStatus, data.DeleteStatus, data.UseIntegration, data.PaymentTime, data.DeliveryTime, data.ReceiveTime, data.CommentTime, data.ModifyTime)
	return ret, err
}

func (m *OmsOrderModel) FindOne(id int64) (*OmsOrder, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", omsOrderRows, m.table)
	var resp OmsOrder
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

func (m *OmsOrderModel) FindAll(Current int64, PageSize int64) (*[]OmsOrder, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", omsOrderRows, m.table)
	var resp []OmsOrder
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

func (m *OmsOrderModel) Count() (int64, error) {
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

func (m *OmsOrderModel) Update(data OmsOrder) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, omsOrderRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.MemberId, data.CouponId, data.OrderSn, data.MemberUsername, data.TotalAmount, data.PayAmount, data.FreightAmount, data.PromotionAmount, data.IntegrationAmount, data.CouponAmount, data.DiscountAmount, data.PayType, data.SourceType, data.Status, data.OrderType, data.DeliveryCompany, data.DeliverySn, data.AutoConfirmDay, data.Integration, data.Growth, data.PromotionInfo, data.BillType, data.BillHeader, data.BillContent, data.BillReceiverPhone, data.BillReceiverEmail, data.ReceiverName, data.ReceiverPhone, data.ReceiverPostCode, data.ReceiverProvince, data.ReceiverCity, data.ReceiverRegion, data.ReceiverDetailAddress, data.Note, data.ConfirmStatus, data.DeleteStatus, data.UseIntegration, data.PaymentTime, data.DeliveryTime, data.ReceiveTime, data.CommentTime, data.ModifyTime, data.Id)
	return err
}

func (m *OmsOrderModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
