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

// UpdateDeptLogic 更新部门信息
/*
Author: LiuFeiHua
Date: 2023/12/18 17:17
*/
type UpdateDeptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateDeptLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateDeptLogic {
	return UpdateDeptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateDept 更新部门信息
func (l *UpdateDeptLogic) UpdateDept(req *types.UpdateDeptReq) (*types.UpdateDeptResp, error) {

	_, err := l.svcCtx.DeptService.UpdateDept(l.ctx, &sysclient.UpdateDeptReq{
		Id:         req.Id,         // 编号
		DeptName:   req.DeptName,   // 部门名称
		DeptStatus: req.DeptStatus, // 部门状态
		DeptSort:   req.DeptSort,   // 部门排序
		ParentId:   req.ParentId,   // 上级机构ID，一级机构为0
		Leader:     req.Leader,     // 负责人
		Phone:      req.Phone,      // 电话号码
		Email:      req.Email,      // 邮箱
		Remark:     req.Remark,     // 备注信息
		ParentIds:  req.ParentIds,  // 上级机构IDs，一级机构为0
		UpdateBy:   l.ctx.Value("userName").(string),
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新部门信息失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.UpdateDeptResp{
		Code:    "000000",
		Message: "更新部门成功",
	}, nil
}
