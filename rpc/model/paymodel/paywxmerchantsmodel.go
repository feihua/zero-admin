package paymodel

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
	payWxMerchantsFieldNames          = builderx.FieldNames(&PayWxMerchants{})
	payWxMerchantsRows                = strings.Join(payWxMerchantsFieldNames, ",")
	payWxMerchantsRowsExpectAutoSet   = strings.Join(stringx.Remove(payWxMerchantsFieldNames, "id", "create_time", "update_time"), ",")
	payWxMerchantsRowsWithPlaceHolder = strings.Join(stringx.Remove(payWxMerchantsFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	PayWxMerchantsModel interface {
		Insert(data PayWxMerchants) (sql.Result, error)
		FindOne(id int64) (*PayWxMerchants, error)
		Update(data PayWxMerchants) error
		Delete(id int64) error
		FindOneByMerId(merId string, payType string) (*PayWxMerchants, error)
	}

	defaultPayWxMerchantsModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PayWxMerchants struct {
		Id         int64     `db:"id"`
		MerId      string    `db:"mer_id"`     // 本地系统商户Id,分配给调用方
		AppId      string    `db:"app_id"`     // 应用ID 微信开放平台审核通过的应用APPID
		MchId      string    `db:"mch_id"`     // 微信支付分配的商户号
		MchKey     string    `db:"mch_key"`    // 微信支付分配的商户秘钥
		PayType    string    `db:"pay_type"`   // APP:APP支付 JSAPI:小程序,公众号 MWEB:H5支付
		NotifyUrl  string    `db:"notify_url"` // 微信支付回调地址
		Remarks    string    `db:"remarks"`    // 备注
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
	}
)

func NewPayWxMerchantsModel(conn sqlx.SqlConn) PayWxMerchantsModel {
	return &defaultPayWxMerchantsModel{
		conn:  conn,
		table: "pay_wx_merchants",
	}
}

func (m *defaultPayWxMerchantsModel) Insert(data PayWxMerchants) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, payWxMerchantsRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.MerId, data.AppId, data.MchId, data.MchKey, data.PayType, data.NotifyUrl, data.Remarks)
	return ret, err
}

func (m *defaultPayWxMerchantsModel) FindOne(id int64) (*PayWxMerchants, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", payWxMerchantsRows, m.table)
	var resp PayWxMerchants
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

func (m *defaultPayWxMerchantsModel) FindOneByMerId(merId string, payType string) (*PayWxMerchants, error) {
	query := fmt.Sprintf("select %s from %s where mer_id = ? and pay_type = ? limit 1", payWxMerchantsRows, m.table)
	var resp PayWxMerchants
	err := m.conn.QueryRow(&resp, query, merId, payType)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultPayWxMerchantsModel) Update(data PayWxMerchants) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, payWxMerchantsRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.MerId, data.AppId, data.MchId, data.MchKey, data.PayType, data.NotifyUrl, data.Remarks, data.Id)
	return err
}

func (m *defaultPayWxMerchantsModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
