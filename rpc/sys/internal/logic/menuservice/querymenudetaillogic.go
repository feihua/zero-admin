package menuservicelogic

import (
	"context"
	"errors"
	"github.com/bytedance/sonic"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"
	"strconv"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMenuDetailLogic 查询菜单详情
/*
Author: LiuFeiHua
Date: 2024/5/30 11:41
*/
type QueryMenuDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMenuDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMenuDetailLogic {
	return &QueryMenuDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryMenuDetail 查询菜单详情
func (l *QueryMenuDetailLogic) QueryMenuDetail(in *sysclient.QueryMenuDetailReq) (*sysclient.QueryMenuDetailResp, error) {
	idStr := strconv.FormatInt(in.Id, 10)
	key := l.svcCtx.RedisKey + "menu"
	cachedData, _ := l.svcCtx.Redis.HgetCtx(l.ctx, key, idStr)

	var cached sysclient.QueryMenuDetailResp
	if sonic.Unmarshal([]byte(cachedData), &cached) == nil {
		return &cached, nil
	}
	menu, err := query.SysMenu.WithContext(l.ctx).Where(query.SysMenu.ID.Eq(in.Id)).First()

	// 1.判断菜单是否存在
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, errors.New("菜单不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询菜单异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询菜单异常")
	}

	data := &sysclient.QueryMenuDetailResp{
		Id:            menu.ID,                                 // 编号
		MenuName:      menu.MenuName,                           // 菜单名称
		ParentId:      menu.ParentID,                           // 父菜单ID，一级菜单为0
		MenuPath:      menu.MenuPath,                           // 前端路由
		MenuPerms:     menu.MenuPerms,                          // 权限标识
		MenuType:      menu.MenuType,                           // 类型 0：目录,1：菜单,2：按钮,3：外链
		MenuIcon:      menu.MenuIcon,                           // 菜单图标
		MenuSort:      menu.MenuSort,                           // 菜单排序
		CreateBy:      menu.CreateBy,                           // 创建者
		CreateTime:    time_util.TimeToStr(menu.CreateTime),    // 创建时间
		UpdateBy:      menu.UpdateBy,                           // 更新者
		UpdateTime:    time_util.TimeToString(menu.UpdateTime), // 更新时间
		MenuStatus:    menu.MenuStatus,                         // 菜单状态
		IsDeleted:     menu.IsDeleted,                          // 是否删除  0：否  1：是
		IsVisible:     menu.IsVisible,                          // 是否可见  0：否  1：是
		Remark:        menu.Remark,                             // 备注信息
		VuePath:       menu.VuePath,                            // vue系统的path
		VueComponent:  menu.VueComponent,                       // vue的页面
		VueIcon:       menu.VueIcon,                            // vue的图标
		VueRedirect:   menu.VueRedirect,                        // vue的路由重定向
		BackgroundUrl: menu.BackgroundURL,                      // 接口地址
	}

	value, _ := sonic.Marshal(data)
	filed := strconv.FormatInt(menu.ID, 10)
	_ = l.svcCtx.Redis.HsetCtx(l.ctx, key, filed, string(value))
	return data, nil
}
