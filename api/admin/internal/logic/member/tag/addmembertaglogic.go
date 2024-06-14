package tag

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddMemberTagLogic 添加会员标签
/*
Author: LiuFeiHua
Date: 2024/5/23 9:20
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

// AddMemberTag 添加会员标签
func (l *AddMemberTagLogic) AddMemberTag(req *types.AddMemberTagReq) (resp *types.AddMemberTagResp, err error) {
	_, err = l.svcCtx.MemberTagService.AddMemberTag(l.ctx, &umsclient.AddMemberTagReq{
		TagName:           req.TagName,
		FinishOrderCount:  req.FinishOrderCount,
		FinishOrderAmount: req.FinishOrderAmount,
		Status:            req.Status,
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加会员标签信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("添加会员标签失败")
	}

	return &types.AddMemberTagResp{
		Code:    "000000",
		Message: "添加会员标签成功",
	}, nil
}
