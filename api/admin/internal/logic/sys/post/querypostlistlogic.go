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

// QueryPostListLogic 岗位信息列表
/*
Author: LiuFeiHua
Date: 2023/12/18 17:19
*/
type QueryPostListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryPostListLogic(ctx context.Context, svcCtx *svc.ServiceContext) QueryPostListLogic {
	return QueryPostListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryPostList 岗位信息列表
func (l *QueryPostListLogic) QueryPostList(req *types.QueryPostListReq) (*types.QueryPostListResp, error) {
	result, err := l.svcCtx.PostService.QueryPostList(l.ctx, &sysclient.QueryPostListReq{
		PostCode:   req.PostCode,
		PostName:   req.PostName,
		PostStatus: req.PostStatus,
		PageNum:    req.Current,
		PageSize:   req.PageSize,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询岗位列表,参数: %+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryPostListData

	for _, post := range result.List {
		list = append(list, &types.QueryPostListData{
			CreateBy:   post.CreateBy,
			CreateTime: post.CreateTime,
			Id:         post.Id,
			PostCode:   post.PostCode,
			PostName:   post.PostName,
			PostSort:   post.PostSort,
			PostStatus: post.PostStatus,
			Remark:     post.Remark,
			UpdateBy:   post.UpdateBy,
			UpdateTime: post.UpdateTime,
		})
	}

	return &types.QueryPostListResp{
		Code:     "000000",
		Message:  "查询岗位成功",
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
	}, nil

}
