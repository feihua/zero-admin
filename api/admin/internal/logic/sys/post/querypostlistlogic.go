package post

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"
	"strings"

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
		PostCode:   strings.TrimSpace(req.PostCode),
		PostName:   strings.TrimSpace(req.PostName),
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
			Id:         post.Id,         // 岗位id
			PostName:   post.PostName,   // 岗位名称
			PostCode:   post.PostCode,   // 岗位编码
			PostStatus: post.PostStatus, // 岗位状态
			PostSort:   post.PostSort,   // 岗位排序
			Remark:     post.Remark,     // 备注信息
			IsDeleted:  post.IsDeleted,  // 是否删除  0：否  1：是
			CreateBy:   post.CreateBy,   // 创建者
			CreateTime: post.CreateTime, // 创建时间
			UpdateBy:   post.UpdateBy,   // 更新者
			UpdateTime: post.UpdateTime, // 更新时间
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
