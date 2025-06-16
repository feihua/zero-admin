package productattributeservicelogic

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

// QueryProductAttributeDetailLogic 查询商品属性详情
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:37
*/
type QueryProductAttributeDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryProductAttributeDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductAttributeDetailLogic {
	return &QueryProductAttributeDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryProductAttributeDetail 查询商品属性详情
func (l *QueryProductAttributeDetailLogic) QueryProductAttributeDetail(in *pmsclient.QueryProductAttributeDetailReq) (*pmsclient.QueryProductAttributeDetailResp, error) {
	item, err := query.PmsProductAttribute.WithContext(l.ctx).Where(query.PmsProductAttribute.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "商品属性不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("商品属性不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询商品属性异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询商品属性异常")
	}

	data := &pmsclient.QueryProductAttributeDetailResp{
		Id:           item.ID,                                          // 主键id
		GroupId:      item.GroupID,                                     // 属性分组ID
		Name:         item.Name,                                        // 属性名称
		InputType:    item.InputType,                                   // 输入类型：1-手动输入，2-单选，3-多选
		ValueType:    item.ValueType,                                   // 值类型：1-文本，2-数字，3-日期
		InputList:    item.InputList,                                   // 可选值列表，用逗号分隔
		Unit:         item.Unit,                                        // 单位
		IsRequired:   item.IsRequired,                                  // 是否必填
		IsSearchable: item.IsSearchable,                                // 是否支持搜索
		IsShow:       item.IsShow,                                      // 是否显示
		Sort:         item.Sort,                                        // 排序
		Status:       item.Status,                                      // 状态：0->禁用；1->启用
		CreateBy:     item.CreateBy,                                    // 创建人ID
		CreateTime:   time_util.TimeToStr(item.CreateTime),             // 创建时间
		UpdateBy:     pointerprocess.DefaltData(item.UpdateBy).(int64), // 更新人ID
		UpdateTime:   time_util.TimeToString(item.UpdateTime),          // 更新时间
	}

	return data, nil
}
