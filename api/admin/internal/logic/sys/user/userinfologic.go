package user

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"
	"strings"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UserInfoLogic 获取用户信息
/*
Author: LiuFeiHua
Date: 2023/12/18 14:01
*/
type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserInfoLogic {
	return UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UserInfo 获取用户信息
func (l *UserInfoLogic) UserInfo() (*types.UserInfoResp, error) {
	userId, _ := l.ctx.Value("userId").(json.Number).Int64()

	resp, err := l.svcCtx.UserService.UserInfo(l.ctx, &sysclient.InfoReq{
		UserId: userId,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据userId: %d,查询用户信息异常:%s", userId, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var MenuTree []*types.ListMenuTree

	// 组装antd ui中的菜单
	for _, item := range resp.MenuListTree {
		MenuTree = append(MenuTree, &types.ListMenuTree{
			Id:       item.Id,
			Path:     item.Path,
			Name:     item.Name,
			ParentId: item.ParentId,
			Icon:     item.Icon,
			VueIcon:  item.VueIcon,
		})
	}

	// 组装element ui中的菜单
	var MenuTreeVue []*types.ListMenuTreeVue

	for _, item := range resp.MenuListTree {

		if len(strings.TrimSpace(item.VuePath)) != 0 {
			MenuTreeVue = append(MenuTreeVue, &types.ListMenuTreeVue{
				Id:           item.Id,
				ParentId:     item.ParentId,
				Title:        item.Name,
				Path:         item.VuePath,
				Name:         item.Name,
				Icon:         item.VueIcon,
				VueRedirect:  item.VueRedirect,
				VueComponent: item.VueComponent,
				Meta: types.MenuTreeMeta{
					Title: item.Name,
					Icon:  item.VueIcon,
				},
			})
		}

	}

	return &types.UserInfoResp{
		Code:    "000000",
		Message: "获取个人信息成功",
		Data: types.UserInfoData{
			Avatar:      resp.Avatar,
			Name:        resp.Name,
			MenuTree:    MenuTree,
			MenuTreeVue: MenuTreeVue,
		},
	}, nil
}
