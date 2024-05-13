package homebrandservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateHomeBrandSortLogic 首页品牌信息
/*
Author: LiuFeiHua
Date: 2024/5/13 15:55
*/
type UpdateHomeBrandSortLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateHomeBrandSortLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHomeBrandSortLogic {
	return &UpdateHomeBrandSortLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateHomeBrandSort 修改推荐品牌排序
func (l *UpdateHomeBrandSortLogic) UpdateHomeBrandSort(in *smsclient.UpdateHomeBrandSortReq) (*smsclient.UpdateHomeBrandSortResp, error) {
	q := query.SmsHomeBrand
	_, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).Update(q.Sort, in.Sort)

	if err != nil {
		return nil, err
	}

	return &smsclient.UpdateHomeBrandSortResp{}, nil
}
