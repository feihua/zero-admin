package menuservicelogic

import (
	"context"
	"errors"

	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryMenuResourceListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMenuResourceListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMenuResourceListLogic {
	return &QueryMenuResourceListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryMenuResourceListLogic) QueryMenuResourceList(in *sysclient.QueryMenuListReq) (*sysclient.QueryMenuListResp, error) {
	menu := query.SysMenu
	q := menu.WithContext(l.ctx)
	q = q.Where(menu.MenuType.Eq(3))
	if len(in.MenuName) > 0 {
		q = q.Where(menu.MenuName.Like("%" + in.MenuName + "%"))
	}

	if in.MenuStatus != 2 {
		q = q.Where(menu.Status.Eq(in.MenuStatus))
	}

	if in.ParentId != 0 {
		q = q.Where(menu.ParentID.Eq(in.ParentId))
	}

	result, count, err := q.Order(query.SysMenu.MenuSort).FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询菜单列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询菜单列表失败")
	}

	var list = make([]*sysclient.MenuListData, 0, len(result))

	for _, item := range result {
		list = append(list, &sysclient.MenuListData{
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
		})
	}

	return &sysclient.QueryMenuListResp{
		Total: count,
		List:  list,
	}, nil
}
