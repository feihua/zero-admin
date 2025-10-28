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
	q = q.Where(menu.MenuType.Eq(2))
	if len(in.MenuName) > 0 {
		q = q.Where(menu.MenuName.Like("%" + in.MenuName + "%"))
	}

	if in.MenuStatus != 2 {
		q = q.Where(menu.MenuStatus.Eq(in.MenuStatus))
	}

	if in.ParentId != 0 {
		q = q.Where(menu.ParentID.Eq(in.ParentId))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询菜单列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询菜单列表失败")
	}

	var list = make([]*sysclient.MenuListData, 0, len(result))

	for _, detail := range result {
		list = append(list, &sysclient.MenuListData{
			Id:            detail.ID,                                 // 编号
			MenuName:      detail.MenuName,                           // 菜单名称
			ParentId:      detail.ParentID,                           // 父菜单ID，一级菜单为0
			MenuPath:      detail.MenuPath,                           // 前端路由
			MenuPerms:     detail.MenuPerms,                          // 权限标识
			MenuType:      detail.MenuType,                           // 类型 0：目录,1：菜单,2：按钮,3：外链
			MenuIcon:      detail.MenuIcon,                           // 菜单图标
			MenuSort:      detail.MenuSort,                           // 菜单排序
			CreateBy:      detail.CreateBy,                           // 创建者
			CreateTime:    time_util.TimeToStr(detail.CreateTime),    // 创建时间
			UpdateBy:      detail.UpdateBy,                           // 更新者
			UpdateTime:    time_util.TimeToString(detail.UpdateTime), // 更新时间
			MenuStatus:    detail.MenuStatus,                         // 菜单状态
			IsDeleted:     detail.IsDeleted,                          // 是否删除  0：否  1：是
			IsVisible:     detail.IsVisible,                          // 是否可见  0：否  1：是
			Remark:        detail.Remark,                             // 备注信息
			VuePath:       detail.VuePath,                            // vue系统的path
			VueComponent:  detail.VueComponent,                       // vue的页面
			VueIcon:       detail.VueIcon,                            // vue的图标
			VueRedirect:   detail.VueRedirect,                        // vue的路由重定向
			BackgroundUrl: detail.BackgroundURL,                      // 接口地址
		})
	}

	return &sysclient.QueryMenuListResp{
		Total: count,
		List:  list,
	}, nil
}
