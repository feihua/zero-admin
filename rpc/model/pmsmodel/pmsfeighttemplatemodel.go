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
	pmsFeightTemplateFieldNames          = builderx.FieldNames(&PmsFeightTemplate{})
	pmsFeightTemplateRows                = strings.Join(pmsFeightTemplateFieldNames, ",")
	pmsFeightTemplateRowsExpectAutoSet   = strings.Join(stringx.Remove(pmsFeightTemplateFieldNames, "id", "create_time", "update_time"), ",")
	pmsFeightTemplateRowsWithPlaceHolder = strings.Join(stringx.Remove(pmsFeightTemplateFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	PmsFeightTemplateModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PmsFeightTemplate struct {
		Id             int64   `db:"id"`
		Name           string  `db:"name"`
		ChargeType     int64   `db:"charge_type"`  // 计费类型:0->按重量；1->按件数
		FirstWeight    float64 `db:"first_weight"` // 首重kg
		FirstFee       float64 `db:"first_fee"`    // 首费（元）
		ContinueWeight float64 `db:"continue_weight"`
		ContinmeFee    float64 `db:"continme_fee"`
		Dest           string  `db:"dest"` // 目的地（省、市）
	}
)

func NewPmsFeightTemplateModel(conn sqlx.SqlConn) *PmsFeightTemplateModel {
	return &PmsFeightTemplateModel{
		conn:  conn,
		table: "pms_feight_template",
	}
}

func (m *PmsFeightTemplateModel) Insert(data PmsFeightTemplate) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, pmsFeightTemplateRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Name, data.ChargeType, data.FirstWeight, data.FirstFee, data.ContinueWeight, data.ContinmeFee, data.Dest)
	return ret, err
}

func (m *PmsFeightTemplateModel) FindOne(id int64) (*PmsFeightTemplate, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", pmsFeightTemplateRows, m.table)
	var resp PmsFeightTemplate
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

func (m *PmsFeightTemplateModel) FindAll(Current int64, PageSize int64) (*[]PmsFeightTemplate, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", pmsFeightTemplateRows, m.table)
	var resp []PmsFeightTemplate
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

func (m *PmsFeightTemplateModel) Count() (int64, error) {
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

func (m *PmsFeightTemplateModel) Update(data PmsFeightTemplate) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, pmsFeightTemplateRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Name, data.ChargeType, data.FirstWeight, data.FirstFee, data.ContinueWeight, data.ContinmeFee, data.Dest, data.Id)
	return err
}

func (m *PmsFeightTemplateModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
