package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberTagDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberTagDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberTagDeleteLogic {
	return MemberTagDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberTagDeleteLogic) MemberTagDelete(req types.DeleteMemberTagReq) (*types.DeleteMemberTagResp, error) {
	_, err := l.svcCtx.MemberTagService.MemberTagDelete(l.ctx, &umsclient.MemberTagDeleteReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据Id: %+v,删除会员标签异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("删除会员标签失败")
	}
	return &types.DeleteMemberTagResp{
		Code:    "000000",
		Message: "",
	}, nil
}
