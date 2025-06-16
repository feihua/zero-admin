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
)

// QueryProductAttributeListLogic 查询商品属性列表
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:37
*/
type QueryProductAttributeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryProductAttributeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductAttributeListLogic {
	return &QueryProductAttributeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryProductAttributeList 查询商品属性列表
func (l *QueryProductAttributeListLogic) QueryProductAttributeList(in *pmsclient.QueryProductAttributeListReq) (*pmsclient.QueryProductAttributeListResp, error) {
	productAttribute := query.PmsProductAttribute
	q := productAttribute.WithContext(l.ctx)
	if in.GroupId != 0 {
		q = q.Where(productAttribute.GroupID.Eq(in.GroupId))
	}
	if len(in.Name) > 0 {
		q = q.Where(productAttribute.Name.Like("%" + in.Name + "%"))
	}
	if in.InputType != 0 {
		q = q.Where(productAttribute.InputType.Eq(in.InputType))
	}

	if in.IsRequired != 0 {
		q = q.Where(productAttribute.IsRequired.Eq(in.IsRequired))
	}
	if in.IsSearchable != 0 {
		q = q.Where(productAttribute.IsSearchable.Eq(in.IsSearchable))
	}
	if in.IsShow != 2 {
		q = q.Where(productAttribute.IsShow.Eq(in.IsShow))
	}
	if in.Status != 2 {
		q = q.Where(productAttribute.Status.Eq(in.Status))
	}
	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询商品属性列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询商品属性列表失败")
	}

	var list []*pmsclient.ProductAttributeListData

	for _, item := range result {
		list = append(list, &pmsclient.ProductAttributeListData{
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

		})
	}

	return &pmsclient.QueryProductAttributeListResp{
		Total: count,
		List:  list,
	}, nil
}
