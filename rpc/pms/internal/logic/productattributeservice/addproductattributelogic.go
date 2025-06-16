package productattributeservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// AddProductAttributeLogic 添加商品属性
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:37
*/
type AddProductAttributeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddProductAttributeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductAttributeLogic {
	return &AddProductAttributeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddProductAttribute 添加商品属性
func (l *AddProductAttributeLogic) AddProductAttribute(in *pmsclient.AddProductAttributeReq) (*pmsclient.AddProductAttributeResp, error) {
	q := query.PmsProductAttribute

	item := &model.PmsProductAttribute{
		GroupID:      in.GroupId,      // 属性分组ID
		Name:         in.Name,         // 属性名称
		InputType:    in.InputType,    // 输入类型：1-手动输入，2-单选，3-多选
		ValueType:    in.ValueType,    // 值类型：1-文本，2-数字，3-日期
		InputList:    in.InputList,    // 可选值列表，用逗号分隔
		Unit:         in.Unit,         // 单位
		IsRequired:   in.IsRequired,   // 是否必填
		IsSearchable: in.IsSearchable, // 是否支持搜索
		IsShow:       in.IsShow,       // 是否显示
		Sort:         in.Sort,         // 排序
		Status:       in.Status,       // 状态：0->禁用；1->启用
		CreateBy:     in.CreateBy,     // 创建人ID
	}

	err := q.WithContext(l.ctx).Create(item)
	if err != nil {
		logc.Errorf(l.ctx, "添加商品属性失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("添加商品属性失败")
	}

	return &pmsclient.AddProductAttributeResp{}, nil
}
