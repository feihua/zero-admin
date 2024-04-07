package logic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberTagAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberTagAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberTagAddLogic {
	return MemberTagAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberTagAddLogic) MemberTagAdd(req types.AddMemberTagReq) (*types.AddMemberTagResp, error) {
	_, err := l.svcCtx.MemberTagService.MemberTagAdd(l.ctx, &umsclient.MemberTagAddReq{
		Name:              req.Name,
		FinishOrderCount:  req.FinishOrderCount,
		FinishOrderAmount: int64(req.FinishOrderAmount),
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加会员标签信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加会员标签失败")
	}

	return &types.AddMemberTagResp{
		Code:    "000000",
		Message: "添加会员标签成功",
	}, nil
}
