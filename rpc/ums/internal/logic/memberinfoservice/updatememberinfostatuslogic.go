package memberinfoservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateMemberInfoStatusLogic 更新会员信息
/*
Author: LiuFeiHua
Date: 2025/05/21 14:18:26
*/
type UpdateMemberInfoStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMemberInfoStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberInfoStatusLogic {
	return &UpdateMemberInfoStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateMemberInfoStatus 更新会员信息状态
func (l *UpdateMemberInfoStatusLogic) UpdateMemberInfoStatus(in *umsclient.UpdateMemberInfoStatusReq) (*umsclient.UpdateMemberInfoStatusResp, error) {
	q := query.UmsMemberInfo

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.IsEnabled, in.IsEnabled)

	if err != nil {
		logc.Errorf(l.ctx, "更新会员信息状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新会员信息状态失败")
	}

	return &umsclient.UpdateMemberInfoStatusResp{}, nil
}
