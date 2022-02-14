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
	umsMemberRuleSettingFieldNames          = builderx.FieldNames(&UmsMemberRuleSetting{})
	umsMemberRuleSettingRows                = strings.Join(umsMemberRuleSettingFieldNames, ",")
	umsMemberRuleSettingRowsExpectAutoSet   = strings.Join(stringx.Remove(umsMemberRuleSettingFieldNames, "id", "create_time", "update_time"), ",")
	umsMemberRuleSettingRowsWithPlaceHolder = strings.Join(stringx.Remove(umsMemberRuleSettingFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	UmsMemberRuleSettingModel struct {
		conn  sqlx.SqlConn
		table string
	}

	UmsMemberRuleSetting struct {
		Id                int64   `db:"id"`
		ContinueSignDay   int64   `db:"continue_sign_day"`   // 连续签到天数
		ContinueSignPoint int64   `db:"continue_sign_point"` // 连续签到赠送数量
		ConsumePerPoint   float64 `db:"consume_per_point"`   // 每消费多少元获取1个点
		LowOrderAmount    float64 `db:"low_order_amount"`    // 最低获取点数的订单金额
		MaxPointPerOrder  int64   `db:"max_point_per_order"` // 每笔订单最高获取点数
		Type              int64   `db:"type"`                // 类型：0->积分规则；1->成长值规则
	}
)

func NewUmsMemberRuleSettingModel(conn sqlx.SqlConn) *UmsMemberRuleSettingModel {
	return &UmsMemberRuleSettingModel{
		conn:  conn,
		table: "ums_member_rule_setting",
	}
}

func (m *UmsMemberRuleSettingModel) Insert(data UmsMemberRuleSetting) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, umsMemberRuleSettingRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.ContinueSignDay, data.ContinueSignPoint, data.ConsumePerPoint, data.LowOrderAmount, data.MaxPointPerOrder, data.Type)
	return ret, err
}

func (m *UmsMemberRuleSettingModel) FindOne(id int64) (*UmsMemberRuleSetting, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", umsMemberRuleSettingRows, m.table)
	var resp UmsMemberRuleSetting
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

func (m *UmsMemberRuleSettingModel) FindAll(Current int64, PageSize int64) (*[]UmsMemberRuleSetting, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", umsMemberRuleSettingRows, m.table)
	var resp []UmsMemberRuleSetting
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

func (m *UmsMemberRuleSettingModel) Count() (int64, error) {
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

func (m *UmsMemberRuleSettingModel) Update(data UmsMemberRuleSetting) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, umsMemberRuleSettingRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.ContinueSignDay, data.ContinueSignPoint, data.ConsumePerPoint, data.LowOrderAmount, data.MaxPointPerOrder, data.Type, data.Id)
	return err
}

func (m *UmsMemberRuleSettingModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
