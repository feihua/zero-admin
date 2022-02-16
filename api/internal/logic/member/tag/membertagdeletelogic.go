package logic

import (
	"context"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

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
	_, err := l.svcCtx.Ums.MemberTagDelete(l.ctx, &umsclient.MemberTagDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据Id: %d,删除会员标签异常:%s", req.Id, err.Error())
		return nil, errorx.NewDefaultError("删除会员标签失败")
	}
	return &types.DeleteMemberTagResp{
		Code:    "000000",
		Message: "",
	}, nil
}
