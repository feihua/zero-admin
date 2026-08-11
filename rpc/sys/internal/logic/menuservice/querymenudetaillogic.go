package menuservicelogic

import (
	"context"
	"errors"
	"strconv"

	"github.com/bytedance/sonic"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"

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
	item, err := query.SysMenu.WithContext(l.ctx).Where(query.SysMenu.ID.Eq(in.Id)).First()

	// 1.判断菜单是否存在
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, errors.New("菜单不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询菜单异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询菜单异常")
	}

	data := &sysclient.QueryMenuDetailResp{
		Id:           item.ID,                                 // 主键
		MenuName:     item.MenuName,                           // 菜单名称
		Ancestors:    item.Ancestors,                          // 祖级列表
		MenuType:     item.MenuType,                           // 菜单类型(1:目录,2:菜单,3:按钮)
		MenuUrl:      item.MenuURL,                            // 路由路径
		MenuIcon:     item.MenuIcon,                           // 菜单图标
		MenuSort:     item.MenuSort,                           // 排序
		ParentId:     item.ParentID,                           // 父id
		ApiUrl:       item.APIURL,                             // 接口url
		Visible:      item.Visible,                            // 显示状态（0:隐藏,显示:1）
		Status:       item.Status,                             // 菜单状态(1:正常，0:禁用)
		DelFlag:      item.DelFlag,                            // 删除标志（0:删除,1:存在）
		Remark:       item.Remark,                             // 备注
		VuePath:      item.VuePath,                            // vue的path
		VueComponent: item.VueComponent,                       // vue的页面
		VueIcon:      item.VueIcon,                            // vue的图标
		VueRedirect:  item.VueRedirect,                        // vue的路由重定向
		AngularIcon:  item.AngularIcon,                        // angular的图标
		ReactIcon:    item.ReactIcon,                          // antd react的图标
		CreateBy:     item.CreateBy,                           // 创建者
		CreateTime:   time_util.TimeToStr(item.CreateTime),    // 创建时间
		UpdateBy:     item.UpdateBy,                           // 更新者
		UpdateTime:   time_util.TimeToString(item.UpdateTime), // 更新时间
	}

	value, _ := sonic.Marshal(data)
	filed := strconv.FormatInt(item.ID, 10)
	_ = l.svcCtx.Redis.HsetCtx(l.ctx, key, filed, string(value))
	return data, nil
}
