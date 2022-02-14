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
	smsCouponProductCategoryRelationFieldNames          = builderx.FieldNames(&SmsCouponProductCategoryRelation{})
	smsCouponProductCategoryRelationRows                = strings.Join(smsCouponProductCategoryRelationFieldNames, ",")
	smsCouponProductCategoryRelationRowsExpectAutoSet   = strings.Join(stringx.Remove(smsCouponProductCategoryRelationFieldNames, "id", "create_time", "update_time"), ",")
	smsCouponProductCategoryRelationRowsWithPlaceHolder = strings.Join(stringx.Remove(smsCouponProductCategoryRelationFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	SmsCouponProductCategoryRelationModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SmsCouponProductCategoryRelation struct {
		Id                  int64  `db:"id"`
		CouponId            int64  `db:"coupon_id"`
		ProductCategoryId   int64  `db:"product_category_id"`
		ProductCategoryName string `db:"product_category_name"` // 产品分类名称
		ParentCategoryName  string `db:"parent_category_name"`  // 父分类名称
	}
)

func NewSmsCouponProductCategoryRelationModel(conn sqlx.SqlConn) *SmsCouponProductCategoryRelationModel {
	return &SmsCouponProductCategoryRelationModel{
		conn:  conn,
		table: "sms_coupon_product_category_relation",
	}
}

func (m *SmsCouponProductCategoryRelationModel) Insert(data SmsCouponProductCategoryRelation) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, smsCouponProductCategoryRelationRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.CouponId, data.ProductCategoryId, data.ProductCategoryName, data.ParentCategoryName)
	return ret, err
}

func (m *SmsCouponProductCategoryRelationModel) FindOne(id int64) (*SmsCouponProductCategoryRelation, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", smsCouponProductCategoryRelationRows, m.table)
	var resp SmsCouponProductCategoryRelation
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

func (m *SmsCouponProductCategoryRelationModel) FindAll(Current int64, PageSize int64) (*[]SmsCouponProductCategoryRelation, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", smsCouponProductCategoryRelationRows, m.table)
	var resp []SmsCouponProductCategoryRelation
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

func (m *SmsCouponProductCategoryRelationModel) Count() (int64, error) {
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

func (m *SmsCouponProductCategoryRelationModel) Update(data SmsCouponProductCategoryRelation) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, smsCouponProductCategoryRelationRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.CouponId, data.ProductCategoryId, data.ProductCategoryName, data.ParentCategoryName, data.Id)
	return err
}

func (m *SmsCouponProductCategoryRelationModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
