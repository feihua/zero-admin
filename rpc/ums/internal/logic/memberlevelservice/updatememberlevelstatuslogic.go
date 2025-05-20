package memberlevelservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateMemberLevelStatusLogic 更新会员等级状态
/*
Author: LiuFeiHua
Date: 2025/5/20 15:18
*/
type UpdateMemberLevelStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMemberLevelStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberLevelStatusLogic {
	return &UpdateMemberLevelStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateMemberLevelStatus 更新会员等级状态
func (l *UpdateMemberLevelStatusLogic) UpdateMemberLevelStatus(in *umsclient.UpdateMemberLevelStatusReq) (*umsclient.UpdateMemberLevelStatusResp, error) {
	q := query.UmsMemberLevel

	_, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).Update(q.IsEnabled, in.IsEnabled)

	if err != nil {
		logc.Errorf(l.ctx, "更新会员等级状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新会员等级状态失败")
	}

	return &umsclient.UpdateMemberLevelStatusResp{}, nil
}
