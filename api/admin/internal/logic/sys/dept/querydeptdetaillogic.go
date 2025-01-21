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
		Id:         detail.Id,         // 编号
		DeptName:   detail.DeptName,   // 部门名称
		DeptStatus: detail.DeptStatus, // 部门状态
		DeptSort:   detail.DeptSort,   // 部门排序
		ParentId:   detail.ParentId,   // 上级机构ID，一级机构为0
		Leader:     detail.Leader,     // 负责人
		Phone:      detail.Phone,      // 电话号码
		Email:      detail.Email,      // 邮箱
		Remark:     detail.Remark,     // 备注信息
		IsDeleted:  detail.IsDeleted,  // 是否删除  0：否  1：是
		ParentIds:  detail.ParentIds,  // 上级机构IDs，一级机构为0
		CreateBy:   detail.CreateBy,   // 创建者
		CreateTime: detail.CreateTime, // 创建时间
		UpdateBy:   detail.UpdateBy,   // 更新者
		UpdateTime: detail.UpdateTime, // 更新时间
	}

	return &types.QueryDeptDetailResp{
		Code:    "000000",
		Message: "查询部门详细信息成功",
		Data:    dept,
	}, nil
}
