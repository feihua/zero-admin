package smsmodel

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
	smsCouponHistoryFieldNames          = builderx.FieldNames(&SmsCouponHistory{})
	smsCouponHistoryRows                = strings.Join(smsCouponHistoryFieldNames, ",")
	smsCouponHistoryRowsExpectAutoSet   = strings.Join(stringx.Remove(smsCouponHistoryFieldNames, "id", "create_time", "update_time"), ",")
	smsCouponHistoryRowsWithPlaceHolder = strings.Join(stringx.Remove(smsCouponHistoryFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	SmsCouponHistoryModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SmsCouponHistory struct {
		Id             int64     `db:"id"`
		CouponId       int64     `db:"coupon_id"`
		MemberId       int64     `db:"member_id"`
		CouponCode     string    `db:"coupon_code"`
		MemberNickname string    `db:"member_nickname"` // 领取人昵称
		GetType        int64     `db:"get_type"`        // 获取类型：0->后台赠送；1->主动获取
		CreateTime     time.Time `db:"create_time"`
		UseStatus      int64     `db:"use_status"` // 使用状态：0->未使用；1->已使用；2->已过期
		UseTime        time.Time `db:"use_time"`   // 使用时间
		OrderId        int64     `db:"order_id"`   // 订单编号
		OrderSn        string    `db:"order_sn"`   // 订单号码
	}
)

func NewSmsCouponHistoryModel(conn sqlx.SqlConn) *SmsCouponHistoryModel {
	return &SmsCouponHistoryModel{
		conn:  conn,
		table: "sms_coupon_history",
	}
}

func (m *SmsCouponHistoryModel) Insert(data SmsCouponHistory) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, smsCouponHistoryRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.CouponId, data.MemberId, data.CouponCode, data.MemberNickname, data.GetType, data.UseStatus, data.UseTime, data.OrderId, data.OrderSn)
	return ret, err
}

func (m *SmsCouponHistoryModel) FindOne(id int64) (*SmsCouponHistory, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", smsCouponHistoryRows, m.table)
	var resp SmsCouponHistory
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

func (m *SmsCouponHistoryModel) FindAll(Current int64, PageSize int64) (*[]SmsCouponHistory, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", smsCouponHistoryRows, m.table)
	var resp []SmsCouponHistory
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

func (m *SmsCouponHistoryModel) Count() (int64, error) {
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

func (m *SmsCouponHistoryModel) Update(data SmsCouponHistory) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, smsCouponHistoryRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.CouponId, data.MemberId, data.CouponCode, data.MemberNickname, data.GetType, data.UseStatus, data.UseTime, data.OrderId, data.OrderSn, data.Id)
	return err
}

func (m *SmsCouponHistoryModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
