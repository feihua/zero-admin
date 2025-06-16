package productattributegroupservicelogic

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

// QueryProductAttributeGroupDetailLogic 查询商品属性分组详情
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:37
*/
type QueryProductAttributeGroupDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryProductAttributeGroupDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductAttributeGroupDetailLogic {
	return &QueryProductAttributeGroupDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryProductAttributeGroupDetail 查询商品属性分组详情
func (l *QueryProductAttributeGroupDetailLogic) QueryProductAttributeGroupDetail(in *pmsclient.QueryProductAttributeGroupDetailReq) (*pmsclient.QueryProductAttributeGroupDetailResp, error) {
	item, err := query.PmsProductAttributeGroup.WithContext(l.ctx).Where(query.PmsProductAttributeGroup.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "商品属性分组不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("商品属性分组不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询商品属性分组异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询商品属性分组异常")
	}

	data := &pmsclient.QueryProductAttributeGroupDetailResp{
		Id:         item.ID,                                          // 主键id
		CategoryId: item.CategoryID,                                  // 分类ID
		Name:       item.Name,                                        // 分组名称
		Sort:       item.Sort,                                        // 排序
		Status:     item.Status,                                      // 状态：0->禁用；1->启用
		CreateBy:   item.CreateBy,                                    // 创建人ID
		CreateTime: time_util.TimeToStr(item.CreateTime),             // 创建时间
		UpdateBy:   pointerprocess.DefaltData(item.UpdateBy).(int64), // 更新人ID
		UpdateTime: time_util.TimeToString(item.UpdateTime),          // 更新时间
	}

	return data, nil
}
