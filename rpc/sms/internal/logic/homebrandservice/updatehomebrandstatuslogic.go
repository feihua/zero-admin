package homebrandservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateHomeBrandStatusLogic 首页品牌信息
/*
Author: LiuFeiHua
Date: 2024/5/13 15:55
*/
type UpdateHomeBrandStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateHomeBrandStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHomeBrandStatusLogic {
	return &UpdateHomeBrandStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateHomeBrandStatus 批量修改推荐品牌状态
func (l *UpdateHomeBrandStatusLogic) UpdateHomeBrandStatus(in *smsclient.UpdateHomeBrandStatusReq) (*smsclient.UpdateHomeBrandStatusResp, error) {
	q := query.SmsHomeBrand
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.RecommendStatus, in.RecommendStatus)

	if err != nil {
		return nil, err
	}

	return &smsclient.UpdateHomeBrandStatusResp{}, nil
}
