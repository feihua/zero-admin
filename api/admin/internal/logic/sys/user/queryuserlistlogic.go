package user

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

// QueryUserListLogic 查询用户列表信息
/*
Author: LiuFeiHua
Date: 2023/12/18 14:04
*/
type QueryUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) QueryUserListLogic {
	return QueryUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryUserList 查询用户列表信息
func (l *QueryUserListLogic) QueryUserList(req *types.QueryUserListReq) (*types.QueryUserListResp, error) {
	result, err := l.svcCtx.UserService.QueryUserList(l.ctx, &sysclient.QueryUserListReq{
		PageNum:  req.Current,
		PageSize: req.PageSize,
		Mobile:   req.Mobile,   // 手机号码
		UserName: req.UserName, // 用户账号
		NickName: req.NickName, // 用户昵称
		Email:    req.Email,    // 用户邮箱
		Status:   req.Status,   // 状态(1:正常，0:禁用)
		DeptId:   req.DeptId,   // 部门ID
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询用户列表,参数: %+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryUserListData

	for _, detail := range result.List {
		list = append(list, &types.QueryUserListData{
			Id:            detail.Id,            // 用户id
			Mobile:        detail.Mobile,        // 手机号码
			UserName:      detail.UserName,      // 用户账号
			NickName:      detail.NickName,      // 用户昵称
			UserType:      detail.UserType,      // 用户类型（00系统用户）
			Avatar:        detail.Avatar,        // 头像路径
			Email:         detail.Email,         // 用户邮箱
			Status:        detail.Status,        // 状态(1:正常，0:禁用)
			DeptId:        detail.DeptId,        // 部门ID
			LoginIp:       detail.LoginIp,       // 最后登录IP
			LoginDate:     detail.LoginDate,     // 最后登录时间
			LoginBrowser:  detail.LoginBrowser,  // 浏览器类型
			LoginOs:       detail.LoginOs,       // 操作系统
			PwdUpdateDate: detail.PwdUpdateDate, // 密码最后更新时间
			Remark:        detail.Remark,        // 备注
			DelFlag:       detail.DelFlag,       // 删除标志（0代表删除 1代表存在）
			CreateBy:      detail.CreateBy,      // 创建者
			CreateTime:    detail.CreateTime,    // 创建时间
			UpdateBy:      detail.UpdateBy,      // 更新者
			UpdateTime:    detail.UpdateTime,    // 更新时间
		})
	}

	return &types.QueryUserListResp{
		Code:     "000000",
		Message:  "查询用户列表成功",
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
	}, nil
}
