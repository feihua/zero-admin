package user

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/common/res"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddUserLogic 新增用户
/*
Author: LiuFeiHua
Date: 2023/12/18 13:57
*/
type AddUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddUserLogic {
	return AddUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddUser 新增用户
func (l *AddUserLogic) AddUser(req *types.AddUserReq) (*types.BaseResp, error) {
	userAddReq := sysclient.AddUserReq{
		Mobile:   req.Mobile,   // 手机号码
		UserName: req.UserName, // 用户账号
		NickName: req.NickName, // 用户昵称
		UserType: req.UserType, // 用户类型（00系统用户）
		Avatar:   req.Avatar,   // 头像路径
		Email:    req.Email,    // 用户邮箱
		Password: req.Password, // 密码
		Status:   req.Status,   // 状态(1:正常，0:禁用)
		DeptId:   req.DeptId,   // 部门ID
		Remark:   req.Remark,   // 备注
		CreateBy: l.ctx.Value("userName").(string),
		PostIds:  req.PostIds, // 岗位id
		RoleIds:  req.RoleIds, // 角色id
	}
	_, err := l.svcCtx.UserService.AddUser(l.ctx, &userAddReq)

	if err != nil {
		logc.Errorf(l.ctx, "添加用户信息失败,参数:%+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return res.Success()
}
