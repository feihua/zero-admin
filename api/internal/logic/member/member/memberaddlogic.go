package logic

import (
	"context"
	"encoding/json"

	"zero-admin/api/internal/common/errorx"
	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"
	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberAddLogic {
	return &MemberAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberAddLogic) MemberAdd(req *types.AddMemberReq) (*types.AddMemberResp, error) {
	_, err := l.svcCtx.Ums.MemberAdd(l.ctx, &umsclient.MemberAddReq{
		Username: req.Username,
		Password: req.Password,
		Phone:    req.Mobile,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("会员添加失败,参数: %s,响应：%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("会员添加失败")
	}

	return &types.AddMemberResp{
		Code:    "000000",
		Message: "添加会员信息成功",
	}, nil
}
