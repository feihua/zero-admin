package member_level

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/common/res"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateMemberLevelStatusLogic 更新会员等级表状态状态
/*
Author: 刘飞华
Date: 2025/02/05 10:34:53
*/
type UpdateMemberLevelStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMemberLevelStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberLevelStatusLogic {
	return &UpdateMemberLevelStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateMemberLevelStatus 更新会员等级表状态
func (l *UpdateMemberLevelStatusLogic) UpdateMemberLevelStatus(req *types.UpdateMemberLevelStatusReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.MemberLevelService.UpdateMemberLevelStatus(l.ctx, &umsclient.UpdateMemberLevelStatusReq{
		Id:            req.Id,            //
		DefaultStatus: req.DefaultStatus, // 是否为默认等级：0->不是；1->是

	})

	if err != nil {
		logc.Errorf(l.ctx, "更新会员等级表状态失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return res.Success()
}
