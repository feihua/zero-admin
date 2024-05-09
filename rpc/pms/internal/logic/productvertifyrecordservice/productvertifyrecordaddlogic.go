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

// ProductVertifyRecordAddLogic 商品审核
/*
Author: LiuFeiHua
Date: 2024/5/8 11:08
*/
type ProductVertifyRecordAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductVertifyRecordAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductVertifyRecordAddLogic {
	return &ProductVertifyRecordAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ProductVertifyRecordAdd 添加商品审核
func (l *ProductVertifyRecordAddLogic) ProductVertifyRecordAdd(in *pmsclient.ProductVertifyRecordAddReq) (*pmsclient.ProductVertifyRecordAddResp, error) {
	err := query.PmsProductVertifyRecord.WithContext(l.ctx).Create(&model.PmsProductVertifyRecord{
		ProductID:  in.ProductId,
		CreateTime: time.Now(),
		VertifyMan: in.VertifyMan,
		Status:     in.Status,
		Detail:     in.Detail,
	})

	if err != nil {
		return nil, err
	}

	return &pmsclient.ProductVertifyRecordAddResp{}, nil
}
