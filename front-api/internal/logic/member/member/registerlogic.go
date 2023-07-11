package member

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	memberAddResp, err := l.svcCtx.MemberService.MemberAdd(l.ctx, &umsclient.MemberAddReq{
		Username: req.Username,
		Password: req.Password,
		Phone:    req.Mobile,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("会员登录失败,参数: %s,响应：%s", reqStr, err.Error())
		return &types.RegisterResp{
			Code:    1,
			Message: "注册失败",
		}, nil
	}
	return &types.RegisterResp{
		Code:    0,
		Message: "注册成功",
		Data:    memberAddResp.Token,
	}, nil
}
