package member_consume_setting

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

// DeleteMemberConsumeSettingLogic 删除积分消费设置
/*
Author: LiuFeiHua
Date: 2025/05/23 11:32:02
*/
type DeleteMemberConsumeSettingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMemberConsumeSettingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMemberConsumeSettingLogic {
	return &DeleteMemberConsumeSettingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteMemberConsumeSetting 删除积分消费设置
func (l *DeleteMemberConsumeSettingLogic) DeleteMemberConsumeSetting(req *types.DeleteMemberConsumeSettingReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.MemberConsumeSettingService.DeleteMemberConsumeSetting(l.ctx, &umsclient.DeleteMemberConsumeSettingReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "删除积分消费设置失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "删除积分消费设置成功",
	}, nil
}
