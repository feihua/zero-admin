package smsmodel

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
	"zero-admin/rpc/sms/smsclient"
)

var _ SmsFlashPromotionProductRelationModel = (*customSmsFlashPromotionProductRelationModel)(nil)

type (
	// SmsFlashPromotionProductRelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSmsFlashPromotionProductRelationModel.
	SmsFlashPromotionProductRelationModel interface {
		smsFlashPromotionProductRelationModel
		Count(ctx context.Context, in *smsclient.FlashPromotionProductRelationListReq) (int64, error)
		FindAll(ctx context.Context, in *smsclient.FlashPromotionProductRelationListReq) (*[]SmsFlashPromotionProductRelation, error)
		DeleteByIds(ctx context.Context, ids []int64) error
	}

	customSmsFlashPromotionProductRelationModel struct {
		*defaultSmsFlashPromotionProductRelationModel
	}
)

// NewSmsFlashPromotionProductRelationModel returns a model for the database table.
func NewSmsFlashPromotionProductRelationModel(conn sqlx.SqlConn) SmsFlashPromotionProductRelationModel {
	return &customSmsFlashPromotionProductRelationModel{
		defaultSmsFlashPromotionProductRelationModel: newSmsFlashPromotionProductRelationModel(conn),
	}
}

func (m *customSmsFlashPromotionProductRelationModel) FindAll(ctx context.Context, in *smsclient.FlashPromotionProductRelationListReq) (*[]SmsFlashPromotionProductRelation, error) {

	where := fmt.Sprintf(" AND flash_promotion_id = %d", in.FlashPromotionId)
	where = where + fmt.Sprintf(" AND flash_promotion_session_id = %d", in.FlashPromotionSessionId)

	query := fmt.Sprintf("select %s from %s where % s limit ?,?", smsFlashPromotionProductRelationRows, m.table, where)
	var resp []SmsFlashPromotionProductRelation
	err := m.conn.QueryRows(&resp, query, (in.Current-1)*in.PageSize, in.PageSize)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customSmsFlashPromotionProductRelationModel) Count(ctx context.Context, in *smsclient.FlashPromotionProductRelationListReq) (int64, error) {
	where := fmt.Sprintf(" AND flash_promotion_id = %d", in.FlashPromotionId)
	where = where + fmt.Sprintf(" AND flash_promotion_session_id = %d", in.FlashPromotionSessionId)
	query := fmt.Sprintf("select count(*) as count from %s where %s", m.table, where)

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

func (m *customSmsFlashPromotionProductRelationModel) DeleteByIds(ctx context.Context, ids []int64) error {
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
