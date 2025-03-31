package post

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddPostLogic 添加岗位信息
/*
Author: LiuFeiHua
Date: 2023/12/18 17:18
*/
type AddPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddPostLogic {
	return AddPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddPost 添加岗位信息
func (l *AddPostLogic) AddPost(req *types.AddPostReq) (*types.AddPostResp, error) {
	_, err := l.svcCtx.PostService.AddPost(l.ctx, &sysclient.AddPostReq{
		PostCode: req.PostCode, // 岗位编码
		PostName: req.PostName, // 岗位名称
		Sort:     req.Sort,     // 显示顺序
		Status:   req.Status,   // 岗位状态（0：停用，1:正常）
		Remark:   req.Remark,   // 备注
		CreateBy: l.ctx.Value("userName").(string),
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加岗位信息失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.AddPostResp{
		Code:    "000000",
		Message: "添加岗位成功",
	}, nil
}
