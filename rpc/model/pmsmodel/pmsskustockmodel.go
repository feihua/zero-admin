package pmsmodel

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ PmsSkuStockModel = (*customPmsSkuStockModel)(nil)

type (
	// PmsSkuStockModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsSkuStockModel.
	PmsSkuStockModel interface {
		pmsSkuStockModel
		Count(ctx context.Context) (int64, error)
		FindAll(ctx context.Context, productId int64) (*[]PmsSkuStock, error)
		DeleteByIds(ctx context.Context, ids []int64) error
		DeleteByProductId(ctx context.Context, productId int64) error
		ReleaseSkuStockLock(ctx context.Context, ProductSkuId, ProductQuantity int64) error
		LockSkuStockLock(ctx context.Context, ProductSkuId, ProductQuantity int64) error
	}

	customPmsSkuStockModel struct {
		*defaultPmsSkuStockModel
	}
)

// NewPmsSkuStockModel returns a model for the database table.
func NewPmsSkuStockModel(conn sqlx.SqlConn) PmsSkuStockModel {
	return &customPmsSkuStockModel{
		defaultPmsSkuStockModel: newPmsSkuStockModel(conn),
	}
}

func (m *customPmsSkuStockModel) FindAll(ctx context.Context, productId int64) (*[]PmsSkuStock, error) {

	query := fmt.Sprintf("select %s from %s where product_id =?", pmsSkuStockRows, m.table)
	var resp []PmsSkuStock
	err := m.conn.QueryRows(&resp, query, productId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customPmsSkuStockModel) Count(ctx context.Context) (int64, error) {
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

func (m *customPmsSkuStockModel) DeleteByIds(ctx context.Context, ids []int64) error {
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

func (m *customPmsSkuStockModel) DeleteByProductId(ctx context.Context, productId int64) error {
	query := fmt.Sprintf("delete from %s where `product_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, productId)
	return err
}

// ReleaseSkuStockLock 释放库存,暂时没有添加事务
func (m *defaultPmsSkuStockModel) ReleaseSkuStockLock(ctx context.Context, ProductSkuId, ProductQuantity int64) error {
	query := fmt.Sprintf("update %s set stock=stock+?,lock_stock=lock_stock-? where `id` = ? and lock_stock>=?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, ProductQuantity, ProductQuantity, ProductSkuId, ProductQuantity)
	return err
}

// LockSkuStockLock 锁库存,暂时没有添加事务
func (m *defaultPmsSkuStockModel) LockSkuStockLock(ctx context.Context, ProductSkuId, ProductQuantity int64) error {
	query := fmt.Sprintf("update %s set stock=stock-?,lock_stock=lock_stock+? where `id` = ? and stock>=?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, ProductQuantity, ProductQuantity, ProductSkuId, ProductQuantity)
	return err
}
