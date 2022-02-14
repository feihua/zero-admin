package logic

import (
	"context"
	"go-zero-admin/api/internal/common/errorx"
	"go-zero-admin/rpc/ums/umsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberLoginLogDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberLoginLogDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberLoginLogDeleteLogic {
	return MemberLoginLogDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberLoginLogDeleteLogic) MemberLoginLogDelete(req types.DeleteMemberLoginLogReq) (*types.DeleteMemberLoginLogResp, error) {
	_, err := l.svcCtx.Ums.MemberLoginLogDelete(l.ctx, &umsclient.MemberLoginLogDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据Id: %d,删除会员登录记录异常:%s", req.Id, err.Error())
		return nil, errorx.NewDefaultError("删除会员登录记录失败")
	}
	return &types.DeleteMemberLoginLogResp{
		Code:    "000000",
		Message: "删除员登录记录成功",
	}, nil
}
