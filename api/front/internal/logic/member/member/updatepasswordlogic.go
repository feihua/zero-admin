package member

import (
	"context"
	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePasswordLogic {
	return &UpdatePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePasswordLogic) UpdatePassword(req *types.UpdatePasswordReq) (resp *types.UpdatePasswordResp, err error) {
	//_, _ = l.svcCtx.MemberService.MemberUpdatePassword(l.ctx, &umsclient.MemberUpdatePasswordReq{
	//	Id:       l.ctx.Value("memberId").(int64),
	//	Password: req.Password,
	//})

	return &types.UpdatePasswordResp{
		Code:    0,
		Message: "修改密码成功",
	}, nil
}
