package productspecvalueservicelogic

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

// UpdateProductSpecValueLogic 更新商品规格值
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:37
*/
type UpdateProductSpecValueLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductSpecValueLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductSpecValueLogic {
	return &UpdateProductSpecValueLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateProductSpecValue 更新商品规格值
func (l *UpdateProductSpecValueLogic) UpdateProductSpecValue(in *pmsclient.UpdateProductSpecValueReq) (*pmsclient.UpdateProductSpecValueResp, error) {
	specValue := query.PmsProductSpecValue
	q := specValue.WithContext(l.ctx)

	// 1.根据商品规格值id查询商品规格值是否已存在
	detail, err := q.Where(specValue.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "商品规格值不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("商品规格值不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询商品规格值异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询商品规格值异常")
	}

	now := time.Now()
	item := &model.PmsProductSpecValue{
		ID:         in.Id,             //
		SpecID:     in.SpecId,         // 规格ID
		Value:      in.Value,          // 规格值
		Sort:       in.Sort,           // 排序
		Status:     in.Status,         // 状态：0->禁用；1->启用
		CreateBy:   detail.CreateBy,   // 创建人ID
		CreateTime: detail.CreateTime, // 创建时间
		UpdateBy:   &in.UpdateBy,      // 更新人ID
		UpdateTime: &now,              // 更新时间
	}

	// 2.商品规格值存在时,则直接更新商品规格值
	err = l.svcCtx.DB.Model(&model.PmsProductSpecValue{}).WithContext(l.ctx).Where(specValue.ID.Eq(in.Id)).Save(item).Error

	if err != nil {
		logc.Errorf(l.ctx, "更新商品规格值失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("更新商品规格值失败")
	}

	return &pmsclient.UpdateProductSpecValueResp{}, nil
}
