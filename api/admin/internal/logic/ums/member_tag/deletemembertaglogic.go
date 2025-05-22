package member_tag

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

// DeleteMemberTagLogic 删除用户标签
/*
Author: LiuFeiHua
Date: 2025/05/22 10:44:59
*/
type DeleteMemberTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMemberTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMemberTagLogic {
	return &DeleteMemberTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteMemberTag 删除用户标签
func (l *DeleteMemberTagLogic) DeleteMemberTag(req *types.DeleteMemberTagReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.MemberTagService.DeleteMemberTag(l.ctx, &umsclient.DeleteMemberTagReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "删除用户标签失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "删除用户标签成功",
	}, nil
}
