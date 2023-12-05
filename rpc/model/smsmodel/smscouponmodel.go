package smsmodel

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
	"zero-admin/rpc/sms/smsclient"
)

var _ SmsCouponModel = (*customSmsCouponModel)(nil)

type (
	// SmsCouponModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSmsCouponModel.
	SmsCouponModel interface {
		smsCouponModel
		Count(ctx context.Context) (int64, error)
		FindAll(ctx context.Context, Current int64, PageSize int64, in *smsclient.CouponListReq) (*[]SmsCoupon, error)
		DeleteByIds(ctx context.Context, ids []int64) error
		FindAllByIds(ctx context.Context, ids []int64) (*[]SmsCoupon, error)
		CouponFindByProductIdAndProductCategoryId(ctx context.Context, productId, categoryId int64) (*[]SmsCoupon, error)
	}

	customSmsCouponModel struct {
		*defaultSmsCouponModel
	}
)

// NewSmsCouponModel returns a model for the database table.
func NewSmsCouponModel(conn sqlx.SqlConn) SmsCouponModel {
	return &customSmsCouponModel{
		defaultSmsCouponModel: newSmsCouponModel(conn),
	}
}

func (m *customSmsCouponModel) FindAll(ctx context.Context, Current int64, PageSize int64, in *smsclient.CouponListReq) (*[]SmsCoupon, error) {

	where := "1=1"
	if in.Type != 4 {
		where = where + fmt.Sprintf(" AND type = %d", in.Type)
	}
	if in.Platform != 3 {
		where = where + fmt.Sprintf(" AND platform = %d", in.Platform)
	}
	if in.UseType != 3 {
		where = where + fmt.Sprintf(" AND use_type = %d", in.UseType)
	}
	if len(in.Name) > 0 {
		where = where + fmt.Sprintf(" AND name like '%%%s%%'", in.Name)
	}

	if len(in.StartTime) > 0 {
		where = where + fmt.Sprintf(" AND start_time >= '%s'", in.StartTime)
	}
	if len(in.EndTime) > 0 {
		where = where + fmt.Sprintf(" AND end_time <= '%s'", in.EndTime)
	}

	query := fmt.Sprintf("select %s from %s where %s limit ?,?", smsCouponRows, m.table, where)

	var resp []SmsCoupon
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

func (m *customSmsCouponModel) Count(ctx context.Context) (int64, error) {
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

func (m *customSmsCouponModel) DeleteByIds(ctx context.Context, ids []int64) error {
	// 拼接占位符 "?"
	placeholders := make([]string, len(ids))
	for i := range ids {
		placeholders[i] = "?"
	}

	// 构建删除语句
	query := fmt.Sprintf("DELETE FROM %s WHERE id IN (%s)", m.table, strings.Join(placeholders, ","))

	// 构建参数列表
	args := make([]interface{}, len(ids))
	for i, id := range ids {
		args[i] = id
	}

	// 执行删除语句
	_, err := m.conn.ExecCtx(ctx, query, args...)
	return err
}

func (m *customSmsCouponModel) FindAllByIds(ctx context.Context, ids []int64) (*[]SmsCoupon, error) {
	query := fmt.Sprintf("select %s from %s where `id` in (?)", smsCouponRows, m.table)
	var resp []SmsCoupon
	err := m.conn.QueryRows(&resp, query, strings.Replace(strings.Trim(fmt.Sprint(ids), "[]"), " ", ",", -1))
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customSmsCouponModel) CouponFindByProductIdAndProductCategoryId(ctx context.Context, productId, categoryId int64) (*[]SmsCoupon, error) {
	query := `SELECT *
FROM sms_coupon
WHERE use_type = 0
  AND start_time < now()
  AND end_time > now()
UNION
(SELECT c.*
 FROM sms_coupon_product_category_relation cpc
          LEFT JOIN sms_coupon c ON cpc.coupon_id = c.id
 WHERE c.use_type = 1
   AND c.start_time < now()
   AND c.end_time > now()
   AND cpc.product_category_id = ?)
UNION
(SELECT c.*
 FROM sms_coupon_product_relation cp
          LEFT JOIN sms_coupon c ON cp.coupon_id = c.id
 WHERE c.use_type = 2
   AND c.start_time < now()
   AND c.end_time > now()
   AND cp.product_id = ?)`
	var resp []SmsCoupon
	err := m.conn.QueryRows(&resp, query, categoryId, productId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
