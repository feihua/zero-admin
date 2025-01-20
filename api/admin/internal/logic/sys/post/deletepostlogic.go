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

// DeletePostLogic 删除岗位信息
/*
Author: LiuFeiHua
Date: 2023/12/18 17:19
*/
type DeletePostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeletePostLogic {
	return DeletePostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeletePost 删除岗位信息
func (l *DeletePostLogic) DeletePost(req *types.DeletePostReq) (*types.DeletePostResp, error) {
	_, err := l.svcCtx.PostService.DeletePost(l.ctx, &sysclient.DeletePostReq{
		Ids: req.Ids, // 岗位ids
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据postId: %+v,删除岗位异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.DeletePostResp{
		Code:    "000000",
		Message: "删除岗位成功",
	}, nil
}
