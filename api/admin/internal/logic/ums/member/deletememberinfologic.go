package member

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteMemberInfoLogic 删除会员信息
/*
Author: LiuFeiHua
Date: 2025/05/21 14:18:25
*/
type DeleteMemberInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMemberInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMemberInfoLogic {
	return &DeleteMemberInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteMemberInfo 删除会员信息
func (l *DeleteMemberInfoLogic) DeleteMemberInfo(req *types.DeleteMemberInfoReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.MemberInfoService.DeleteMemberInfo(l.ctx, &umsclient.DeleteMemberInfoReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "删除会员信息失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "删除会员信息成功",
	}, nil
}
