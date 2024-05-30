package dept

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

// QueryDeptDetailLogic 查询部门详细信息
/*
Author: LiuFeiHua
Date: 2024/5/29 16:41
*/
type QueryDeptDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryDeptDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryDeptDetailLogic {
	return &QueryDeptDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryDeptDetail 查询部门详细信息
func (l *QueryDeptDetailLogic) QueryDeptDetail(req *types.QueryDeptDetailReq) (resp *types.QueryDeptDetailResp, err error) {
	detail, err := l.svcCtx.DeptService.QueryDeptDetail(l.ctx, &sysclient.QueryDeptDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询部门详细信息,参数: %+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	dept := types.QueryDeptDetailData{
		CreateBy:   detail.CreateBy,
		CreateTime: detail.CreateTime,
		DeptName:   detail.DeptName,
		DeptSort:   detail.DeptSort,
		DeptStatus: detail.DeptStatus,
		Email:      detail.Email,
		Id:         detail.Id,
		Leader:     detail.Leader,
		ParentId:   detail.ParentId,
		ParentIds:  detail.ParentIds,
		Phone:      detail.Phone,
		Remark:     detail.Remark,
		UpdateBy:   detail.UpdateBy,
		UpdateTime: detail.UpdateTime,
	}

	return &types.QueryDeptDetailResp{
		Code:    "000000",
		Message: "查询部门详细信息成功",
		Data:    dept,
	}, nil
}
