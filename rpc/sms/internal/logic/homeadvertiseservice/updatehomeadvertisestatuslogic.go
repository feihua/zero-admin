package homeadvertiseservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateHomeAdvertiseStatusLogic 首页广告
/*
Author: LiuFeiHua
Date: 2024/5/13 17:38
*/
type UpdateHomeAdvertiseStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateHomeAdvertiseStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHomeAdvertiseStatusLogic {
	return &UpdateHomeAdvertiseStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateHomeAdvertiseStatus 修改上下线状态
func (l *UpdateHomeAdvertiseStatusLogic) UpdateHomeAdvertiseStatus(in *smsclient.UpdateHomeAdvertiseStatusReq) (*smsclient.UpdateHomeAdvertiseStatusResp, error) {
	q := query.SmsHomeAdvertise
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.Status, in.Status)
	if err != nil {
		logc.Errorf(l.ctx, "修改上下线状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("修改上下线状态失败")
	}

	return &smsclient.UpdateHomeAdvertiseStatusResp{}, nil
}
