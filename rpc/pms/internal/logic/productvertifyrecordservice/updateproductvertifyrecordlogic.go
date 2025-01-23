package productvertifyrecordservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"time"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateProductVertifyRecordLogic 更新商品审核记录
/*
Author: LiuFeiHua
Date: 2024/6/12 17:14
*/
type UpdateProductVertifyRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductVertifyRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductVertifyRecordLogic {
	return &UpdateProductVertifyRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateProductVertifyRecord 更新商品审核记录
func (l *UpdateProductVertifyRecordLogic) UpdateProductVertifyRecord(in *pmsclient.UpdateProductVertifyRecordReq) (*pmsclient.UpdateProductVertifyRecordResp, error) {
	q := query.PmsProductVertifyRecord
	_, err := q.WithContext(l.ctx).Updates(&model.PmsProductVertifyRecord{
		ID:         in.Id,
		ProductID:  in.ProductId,
		CreateTime: time.Now(),
		ReviewMan:  in.VertifyMan,
		Status:     in.Status,
		Detail:     in.Detail,
	})

	if err != nil {
		return nil, err
	}

	return &pmsclient.UpdateProductVertifyRecordResp{}, nil
}
