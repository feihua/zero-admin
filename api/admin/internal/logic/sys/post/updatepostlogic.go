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

// UpdatePostLogic 更新岗位信息
/*
Author: LiuFeiHua
Date: 2023/12/18 17:19
*/
type UpdatePostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdatePostLogic {
	return UpdatePostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdatePost 更新岗位信息
func (l *UpdatePostLogic) UpdatePost(req *types.UpdatePostReq) (*types.UpdatePostResp, error) {
	_, err := l.svcCtx.PostService.UpdatePost(l.ctx, &sysclient.UpdatePostReq{
		Id:         req.Id,
		PostCode:   req.PostCode,
		PostName:   req.PostName,
		PostSort:   req.PostSort,
		PostStatus: req.PostStatus,
		Remark:     req.Remark,
		UpdateBy:   l.ctx.Value("userName").(string),
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新岗位信息失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.UpdatePostResp{
		Code:    "000000",
		Message: "更新岗位信息成功",
	}, nil
}
