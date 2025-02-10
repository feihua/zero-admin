package user

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
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
		PageNum:    req.Current,
		PageSize:   req.PageSize,
		DeptId:     req.DeptId,
		Email:      strings.TrimSpace(req.Email),
		Mobile:     strings.TrimSpace(req.Mobile),
		NickName:   strings.TrimSpace(req.NickName),
		UserName:   strings.TrimSpace(req.UserName),
		UserStatus: req.UserStatus,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询用户列表,参数: %+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryUserListData

	for _, detail := range result.List {
		list = append(list, &types.QueryUserListData{
			Id:           detail.Id,           // 编号
			UserName:     detail.UserName,     // 用户名
			NickName:     detail.NickName,     // 昵称
			Avatar:       detail.Avatar,       // 头像
			Email:        detail.Email,        // 邮箱
			Mobile:       detail.Mobile,       // 手机号
			UserStatus:   detail.UserStatus,   // 帐号状态（0正常 1停用）
			DeptId:       detail.DeptId,       // 部门id
			Remark:       detail.Remark,       // 备注信息
			IsDeleted:    detail.IsDeleted,    // 是否删除  0：否  1：是
			LoginTime:    detail.LoginTime,    // 登录时间
			LoginIp:      detail.LoginIp,      // 登录ip
			LoginOs:      detail.LoginOs,      // 登录os
			LoginBrowser: detail.LoginBrowser, // 登录浏览器
			CreateBy:     detail.CreateBy,     // 创建者
			CreateTime:   detail.CreateTime,   // 创建时间
			UpdateBy:     detail.UpdateBy,     // 更新者
			UpdateTime:   detail.UpdateTime,   // 更新时间
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
