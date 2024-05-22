package member

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

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
	_, err := l.svcCtx.MemberService.MemberDelete(l.ctx, &umsclient.MemberDeleteReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据Id: %+v,删除会员信息异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("删除会员信息失败")
	}
	return &types.DeleteMemberResp{
		Code:    "000000",
		Message: "删除会员信息成功",
	}, nil
}
