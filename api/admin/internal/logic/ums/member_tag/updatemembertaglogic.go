package member_tag

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/common/res"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateMemberTagLogic 更新会员标签
/*
Author: LiuFeiHua
Date: 2024/5/23 9:22
*/
type UpdateMemberTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMemberTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberTagLogic {
	return &UpdateMemberTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateMemberTag 更新会员标签
func (l *UpdateMemberTagLogic) UpdateMemberTag(req *types.UpdateMemberTagReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.MemberTagService.UpdateMemberTag(l.ctx, &umsclient.UpdateMemberTagReq{
		Id:                req.Id,
		TagName:           req.TagName,           // 标签名称
		FinishOrderCount:  req.FinishOrderCount,  // 自动打标签完成订单数量
		Status:            req.Status,            // 状态：0->禁用；1->启用
		FinishOrderAmount: req.FinishOrderAmount, // 自动打标签完成订单金额
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新会员标签信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新会员标签失败")
	}

	return res.Success()
}
