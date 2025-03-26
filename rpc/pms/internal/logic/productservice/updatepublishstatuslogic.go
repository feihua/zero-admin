package productservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdatePublishStatusLogic
/*
Author: LiuFeiHua
Date: 2024/5/15 16:09
*/
type UpdatePublishStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePublishStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePublishStatusLogic {
	return &UpdatePublishStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdatePublishStatus 批量上下架商品
func (l *UpdatePublishStatusLogic) UpdatePublishStatus(in *pmsclient.UpdateProductStatusReq) (*pmsclient.UpdateProductStatusResp, error) {
	q := query.PmsProduct
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.PublishStatus, in.Status)

	if err != nil {
		logc.Errorf(l.ctx, "批量上下架商品失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("批量上下架商品失败")
	}

	return &pmsclient.UpdateProductStatusResp{}, nil
}
