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

// QueryPostDetailLogic 查询岗位详情
/*
Author: LiuFeiHua
Date: 2024/5/29 17:26
*/
type QueryPostDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryPostDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryPostDetailLogic {
	return &QueryPostDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryPostDetail 查询岗位详情
func (l *QueryPostDetailLogic) QueryPostDetail(req *types.QueryPostDetailReq) (resp *types.QueryPostDetailResp, err error) {
	detail, err := l.svcCtx.PostService.QueryPostDetail(l.ctx, &sysclient.QueryPostDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询岗位详情,参数: %+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	post := types.QueryPostDetailData{
		Id:         detail.Id,         // 岗位id
		PostCode:   detail.PostCode,   // 岗位编码
		PostName:   detail.PostName,   // 岗位名称
		Sort:       detail.Sort,       // 显示顺序
		Status:     detail.Status,     // 岗位状态（0：停用，1:正常）
		Remark:     detail.Remark,     // 备注
		CreateBy:   detail.CreateBy,   // 创建者
		CreateTime: detail.CreateTime, // 创建时间
		UpdateBy:   detail.UpdateBy,   // 更新者
		UpdateTime: detail.UpdateTime, // 更新时间
	}

	return &types.QueryPostDetailResp{
		Code:    "000000",
		Message: "查询岗位详情成功",
		Data:    post,
	}, nil
}
