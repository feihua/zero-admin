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

// AddMemberTagLogic 添加用户标签
/*
Author: LiuFeiHua
Date: 2025/05/22 10:44:59
*/
type AddMemberTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddMemberTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMemberTagLogic {
	return &AddMemberTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddMemberTag 添加用户标签
func (l *AddMemberTagLogic) AddMemberTag(req *types.AddMemberTagReq) (resp *types.BaseResp, err error) {
	userId, err := common.GetUserId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.MemberTagService.AddMemberTag(l.ctx, &umsclient.AddMemberTagReq{
		TagName:           req.TagName,                    // 标签名称
		Description:       req.Description,                // 标签描述
		FinishOrderCount:  req.FinishOrderCount,           // 自动打标签完成订单数量
		FinishOrderAmount: float32(req.FinishOrderAmount), // 自动打标签完成订单金额
		Status:            req.Status,                     // 状态：0-禁用，1-启用
		CreateBy:          userId,                         // 创建人ID
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加用户标签失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "添加用户标签成功",
	}, nil
}
