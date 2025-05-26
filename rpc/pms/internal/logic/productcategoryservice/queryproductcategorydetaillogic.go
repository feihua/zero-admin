package productcategoryservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/pointerprocess"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

// QueryProductCategoryDetailLogic 查询产品分类详情
/*
Author: LiuFeiHua
Date: 2025/05/26 10:33:54
*/
type QueryProductCategoryDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryProductCategoryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductCategoryDetailLogic {
	return &QueryProductCategoryDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryProductCategoryDetail 查询产品分类详情
func (l *QueryProductCategoryDetailLogic) QueryProductCategoryDetail(in *pmsclient.QueryProductCategoryDetailReq) (*pmsclient.QueryProductCategoryDetailResp, error) {
	item, err := query.PmsProductCategory.WithContext(l.ctx).Where(query.PmsProductCategory.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "产品分类不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("产品分类不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询产品分类异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询产品分类异常")
	}

	data := &pmsclient.QueryProductCategoryDetailResp{
		Id:           item.ID,                                          //
		ParentId:     item.ParentID,                                    // 上级分类的编号：0表示一级分类
		Name:         item.Name,                                        // 商品分类名称
		Level:        item.Level,                                       // 分类级别：0->1级；1->2级
		ProductCount: item.ProductCount,                                // 商品数量
		ProductUnit:  item.ProductUnit,                                 // 商品单位
		NavStatus:    item.NavStatus,                                   // 是否显示在导航栏：0->不显示；1->显示
		Sort:         item.Sort,                                        // 排序
		Icon:         item.Icon,                                        // 图标
		Keywords:     item.Keywords,                                    // 关键字
		Description:  item.Description,                                 // 描述
		IsEnabled:    item.IsEnabled,                                   // 是否启用
		CreateBy:     item.CreateBy,                                    // 创建人ID
		CreateTime:   time_util.TimeToStr(item.CreateTime),             // 创建时间
		UpdateBy:     pointerprocess.DefaltData(item.UpdateBy).(int64), // 更新人ID
		UpdateTime:   time_util.TimeToString(item.UpdateTime),          // 更新时间
	}

	return data, nil
}
