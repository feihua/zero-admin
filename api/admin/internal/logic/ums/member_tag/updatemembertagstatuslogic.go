package member_tag

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateMemberTagStatusLogic 更新用户标签状态状态
/*
Author: LiuFeiHua
Date: 2025/05/22 10:44:59
*/
type UpdateMemberTagStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMemberTagStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberTagStatusLogic {
	return &UpdateMemberTagStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateMemberTagStatus 更新用户标签状态
func (l *UpdateMemberTagStatusLogic) UpdateMemberTagStatus(req *types.UpdateMemberTagStatusReq) (resp *types.BaseResp, err error) {
	userId, err := common.GetUserId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.MemberTagService.UpdateMemberTagStatus(l.ctx, &umsclient.UpdateMemberTagStatusReq{
		Ids:      req.Ids,    // 主键ID
		Status:   req.Status, // 状态：0-禁用，1-启用
		UpdateBy: userId,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新用户标签状态失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "更新用户标签状态成功",
	}, nil
}
