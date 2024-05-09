package productvertifyrecordservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"time"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// ProductVertifyRecordUpdateLogic 商品审核
/*
Author: LiuFeiHua
Date: 2024/5/8 11:09
*/
type ProductVertifyRecordUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductVertifyRecordUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductVertifyRecordUpdateLogic {
	return &ProductVertifyRecordUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ProductVertifyRecordUpdate 更新商品审核
func (l *ProductVertifyRecordUpdateLogic) ProductVertifyRecordUpdate(in *pmsclient.ProductVertifyRecordUpdateReq) (*pmsclient.ProductVertifyRecordUpdateResp, error) {
	q := query.PmsProductVertifyRecord
	_, err := q.WithContext(l.ctx).Updates(&model.PmsProductVertifyRecord{
		ID:         in.Id,
		ProductID:  in.ProductId,
		CreateTime: time.Now(),
		VertifyMan: in.VertifyMan,
		Status:     in.Status,
		Detail:     in.Detail,
	})

	if err != nil {
		return nil, err
	}

	return &pmsclient.ProductVertifyRecordUpdateResp{}, nil
}
