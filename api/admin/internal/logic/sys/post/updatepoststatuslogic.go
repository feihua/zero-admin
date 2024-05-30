package post

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdatePostStatusLogic 更新岗位状态
/*
Author: LiuFeiHua
Date: 2024/5/29 17:23
*/
type UpdatePostStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePostStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePostStatusLogic {
	return &UpdatePostStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdatePostStatus 更新岗位状态
func (l *UpdatePostStatusLogic) UpdatePostStatus(req *types.UpdatePostStatusReq) (resp *types.UpdatePostStatusResp, err error) {
	_, err = l.svcCtx.PostService.UpdatePostStatus(l.ctx, &sysclient.UpdatePostStatusReq{
		Ids:        req.PostIds,
		PostStatus: req.PostStatus,
		UpdateBy:   l.ctx.Value("userName").(string),
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新岗位状态失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.UpdatePostStatusResp{
		Code:    "000000",
		Message: "更新岗位状态成功",
	}, nil
}
