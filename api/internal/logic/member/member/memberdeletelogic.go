package logic

import (
	"context"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberDeleteLogic {
	return MemberDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberDeleteLogic) MemberDelete(req types.DeleteMemberReq) (*types.DeleteMemberResp, error) {
	_, err := l.svcCtx.Ums.MemberDelete(l.ctx, &umsclient.MemberDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据Id: %d,删除会员信息异常:%s", req.Id, err.Error())
		return nil, errorx.NewDefaultError("删除会员信息失败")
	}
	return &types.DeleteMemberResp{
		Code:    "000000",
		Message: "删除会员信息成功",
	}, nil
}
