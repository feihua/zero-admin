package pmsmodel

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
	pmsProductOperateLogFieldNames          = builderx.FieldNames(&PmsProductOperateLog{})
	pmsProductOperateLogRows                = strings.Join(pmsProductOperateLogFieldNames, ",")
	pmsProductOperateLogRowsExpectAutoSet   = strings.Join(stringx.Remove(pmsProductOperateLogFieldNames, "id", "create_time", "update_time"), ",")
	pmsProductOperateLogRowsWithPlaceHolder = strings.Join(stringx.Remove(pmsProductOperateLogFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	PmsProductOperateLogModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PmsProductOperateLog struct {
		Id               int64     `db:"id"`
		ProductId        int64     `db:"product_id"`
		PriceOld         float64   `db:"price_old"`
		PriceNew         float64   `db:"price_new"`
		SalePriceOld     float64   `db:"sale_price_old"`
		SalePriceNew     float64   `db:"sale_price_new"`
		GiftPointOld     int64     `db:"gift_point_old"` // 赠送的积分
		GiftPointNew     int64     `db:"gift_point_new"`
		UsePointLimitOld int64     `db:"use_point_limit_old"`
		UsePointLimitNew int64     `db:"use_point_limit_new"`
		OperateMan       string    `db:"operate_man"` // 操作人
		CreateTime       time.Time `db:"create_time"`
	}
)

func NewPmsProductOperateLogModel(conn sqlx.SqlConn) *PmsProductOperateLogModel {
	return &PmsProductOperateLogModel{
		conn:  conn,
		table: "pms_product_operate_log",
	}
}

func (m *PmsProductOperateLogModel) Insert(data PmsProductOperateLog) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, pmsProductOperateLogRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.ProductId, data.PriceOld, data.PriceNew, data.SalePriceOld, data.SalePriceNew, data.GiftPointOld, data.GiftPointNew, data.UsePointLimitOld, data.UsePointLimitNew, data.OperateMan)
	return ret, err
}

func (m *PmsProductOperateLogModel) FindOne(id int64) (*PmsProductOperateLog, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", pmsProductOperateLogRows, m.table)
	var resp PmsProductOperateLog
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

func (m *PmsProductOperateLogModel) FindAll(Current int64, PageSize int64) (*[]PmsProductOperateLog, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", pmsProductOperateLogRows, m.table)
	var resp []PmsProductOperateLog
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

func (m *PmsProductOperateLogModel) Count() (int64, error) {
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

func (m *PmsProductOperateLogModel) Update(data PmsProductOperateLog) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, pmsProductOperateLogRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.ProductId, data.PriceOld, data.PriceNew, data.SalePriceOld, data.SalePriceNew, data.GiftPointOld, data.GiftPointNew, data.UsePointLimitOld, data.UsePointLimitNew, data.OperateMan, data.Id)
	return err
}

func (m *PmsProductOperateLogModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
