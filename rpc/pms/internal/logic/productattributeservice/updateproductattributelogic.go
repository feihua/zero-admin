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
	"gorm.io/gorm"
	"time"
)

// UpdateProductAttributeLogic 更新商品属性
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:37
*/
type UpdateProductAttributeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductAttributeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductAttributeLogic {
	return &UpdateProductAttributeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateProductAttribute 更新商品属性
func (l *UpdateProductAttributeLogic) UpdateProductAttribute(in *pmsclient.UpdateProductAttributeReq) (*pmsclient.UpdateProductAttributeResp, error) {
	attribute := query.PmsProductAttribute
	q := attribute.WithContext(l.ctx)

	// 1.根据商品属性id查询商品属性是否已存在
	detail, err := q.Where(attribute.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "商品属性不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("商品属性不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询商品属性异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询商品属性异常")
	}

	now := time.Now()
	item := &model.PmsProductAttribute{
		ID:           in.Id,             // 主键id
		GroupID:      in.GroupId,        // 属性分组ID
		Name:         in.Name,           // 属性名称
		InputType:    in.InputType,      // 输入类型：1-手动输入，2-单选，3-多选
		ValueType:    in.ValueType,      // 值类型：1-文本，2-数字，3-日期
		InputList:    in.InputList,      // 可选值列表，用逗号分隔
		Unit:         in.Unit,           // 单位
		IsRequired:   in.IsRequired,     // 是否必填
		IsSearchable: in.IsSearchable,   // 是否支持搜索
		IsShow:       in.IsShow,         // 是否显示
		Sort:         in.Sort,           // 排序
		Status:       in.Status,         // 状态：0->禁用；1->启用
		CreateBy:     detail.CreateBy,   // 创建人ID
		CreateTime:   detail.CreateTime, // 创建时间
		UpdateBy:     &in.UpdateBy,      // 更新人ID
		UpdateTime:   &now,              // 更新时间
	}

	// 2.商品属性存在时,则直接更新商品属性
	err = l.svcCtx.DB.Model(&model.PmsProductAttribute{}).WithContext(l.ctx).Where(attribute.ID.Eq(in.Id)).Save(item).Error

	if err != nil {
		logc.Errorf(l.ctx, "更新商品属性失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("更新商品属性失败")
	}

	return &pmsclient.UpdateProductAttributeResp{}, nil
}
