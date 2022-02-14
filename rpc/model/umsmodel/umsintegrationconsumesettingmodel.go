package umsmodel

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
	umsIntegrationConsumeSettingFieldNames          = builderx.FieldNames(&UmsIntegrationConsumeSetting{})
	umsIntegrationConsumeSettingRows                = strings.Join(umsIntegrationConsumeSettingFieldNames, ",")
	umsIntegrationConsumeSettingRowsExpectAutoSet   = strings.Join(stringx.Remove(umsIntegrationConsumeSettingFieldNames, "id", "create_time", "update_time"), ",")
	umsIntegrationConsumeSettingRowsWithPlaceHolder = strings.Join(stringx.Remove(umsIntegrationConsumeSettingFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	UmsIntegrationConsumeSettingModel struct {
		conn  sqlx.SqlConn
		table string
	}

	UmsIntegrationConsumeSetting struct {
		Id                 int64 `db:"id"`
		DeductionPerAmount int64 `db:"deduction_per_amount"`  // 每一元需要抵扣的积分数量
		MaxPercentPerOrder int64 `db:"max_percent_per_order"` // 每笔订单最高抵用百分比
		UseUnit            int64 `db:"use_unit"`              // 每次使用积分最小单位100
		CouponStatus       int64 `db:"coupon_status"`         // 是否可以和优惠券同用；0->不可以；1->可以
	}
)

func NewUmsIntegrationConsumeSettingModel(conn sqlx.SqlConn) *UmsIntegrationConsumeSettingModel {
	return &UmsIntegrationConsumeSettingModel{
		conn:  conn,
		table: "ums_integration_consume_setting",
	}
}

func (m *UmsIntegrationConsumeSettingModel) Insert(data UmsIntegrationConsumeSetting) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, umsIntegrationConsumeSettingRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.DeductionPerAmount, data.MaxPercentPerOrder, data.UseUnit, data.CouponStatus)
	return ret, err
}

func (m *UmsIntegrationConsumeSettingModel) FindOne(id int64) (*UmsIntegrationConsumeSetting, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", umsIntegrationConsumeSettingRows, m.table)
	var resp UmsIntegrationConsumeSetting
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

func (m *UmsIntegrationConsumeSettingModel) FindAll(Current int64, PageSize int64) (*[]UmsIntegrationConsumeSetting, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", umsIntegrationConsumeSettingRows, m.table)
	var resp []UmsIntegrationConsumeSetting
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

func (m *UmsIntegrationConsumeSettingModel) Count() (int64, error) {
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

func (m *UmsIntegrationConsumeSettingModel) Update(data UmsIntegrationConsumeSetting) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, umsIntegrationConsumeSettingRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.DeductionPerAmount, data.MaxPercentPerOrder, data.UseUnit, data.CouponStatus, data.Id)
	return err
}

func (m *UmsIntegrationConsumeSettingModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
