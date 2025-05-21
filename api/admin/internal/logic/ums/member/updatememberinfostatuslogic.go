package member

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateMemberInfoStatusLogic 更新会员信息状态状态
/*
Author: LiuFeiHua
Date: 2025/05/21 14:18:26
*/
type UpdateMemberInfoStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMemberInfoStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberInfoStatusLogic {
	return &UpdateMemberInfoStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateMemberInfoStatus 更新会员信息状态
func (l *UpdateMemberInfoStatusLogic) UpdateMemberInfoStatus(req *types.UpdateMemberInfoStatusReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.MemberInfoService.UpdateMemberInfoStatus(l.ctx, &umsclient.UpdateMemberInfoStatusReq{
		Ids:       req.Ids, // 主键ID
		IsEnabled: req.IsEnabled,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新会员信息状态失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "更新会员信息状态成功",
	}, nil
}
