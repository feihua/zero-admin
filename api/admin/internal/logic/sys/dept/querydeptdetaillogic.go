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
		Id:         detail.Id,         // 部门id
		ParentId:   detail.ParentId,   // 上级部门id
		Ancestors:  detail.Ancestors,  // 祖级列表
		DeptName:   detail.DeptName,   // 部门名称
		Sort:       detail.Sort,       // 显示顺序
		Leader:     detail.Leader,     // 负责人
		Phone:      detail.Phone,      // 联系电话
		Email:      detail.Email,      // 邮箱
		Status:     detail.Status,     // 部门状态（0：停用，1:正常）
		DelFlag:    detail.DelFlag,    // 删除标志（0代表删除 1代表存在）
		Remark:     detail.Remark,     // 备注信息
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
