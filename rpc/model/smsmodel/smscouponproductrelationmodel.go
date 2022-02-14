package smsmodel

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
	smsCouponProductRelationFieldNames          = builderx.FieldNames(&SmsCouponProductRelation{})
	smsCouponProductRelationRows                = strings.Join(smsCouponProductRelationFieldNames, ",")
	smsCouponProductRelationRowsExpectAutoSet   = strings.Join(stringx.Remove(smsCouponProductRelationFieldNames, "id", "create_time", "update_time"), ",")
	smsCouponProductRelationRowsWithPlaceHolder = strings.Join(stringx.Remove(smsCouponProductRelationFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	SmsCouponProductRelationModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SmsCouponProductRelation struct {
		Id          int64  `db:"id"`
		CouponId    int64  `db:"coupon_id"`
		ProductId   int64  `db:"product_id"`
		ProductName string `db:"product_name"` // 商品名称
		ProductSn   string `db:"product_sn"`   // 商品编码
	}
)

func NewSmsCouponProductRelationModel(conn sqlx.SqlConn) *SmsCouponProductRelationModel {
	return &SmsCouponProductRelationModel{
		conn:  conn,
		table: "sms_coupon_product_relation",
	}
}

func (m *SmsCouponProductRelationModel) Insert(data SmsCouponProductRelation) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, smsCouponProductRelationRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.CouponId, data.ProductId, data.ProductName, data.ProductSn)
	return ret, err
}

func (m *SmsCouponProductRelationModel) FindOne(id int64) (*SmsCouponProductRelation, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", smsCouponProductRelationRows, m.table)
	var resp SmsCouponProductRelation
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

func (m *SmsCouponProductRelationModel) FindAll(Current int64, PageSize int64) (*[]SmsCouponProductRelation, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", smsCouponProductRelationRows, m.table)
	var resp []SmsCouponProductRelation
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

func (m *SmsCouponProductRelationModel) Count() (int64, error) {
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

func (m *SmsCouponProductRelationModel) Update(data SmsCouponProductRelation) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, smsCouponProductRelationRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.CouponId, data.ProductId, data.ProductName, data.ProductSn, data.Id)
	return err
}

func (m *SmsCouponProductRelationModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
