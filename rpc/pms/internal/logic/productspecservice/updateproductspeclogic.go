package productspecservicelogic

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

// UpdateProductSpecLogic 更新商品规格
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:37
*/
type UpdateProductSpecLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductSpecLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductSpecLogic {
	return &UpdateProductSpecLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateProductSpec 更新商品规格
func (l *UpdateProductSpecLogic) UpdateProductSpec(in *pmsclient.UpdateProductSpecReq) (*pmsclient.UpdateProductSpecResp, error) {
	spec := query.PmsProductSpec
	q := spec.WithContext(l.ctx)

	// 1.根据商品规格id查询商品规格是否已存在
	detail, err := q.Where(spec.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "商品规格不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("商品规格不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询商品规格异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询商品规格异常")
	}

	now := time.Now()
	item := &model.PmsProductSpec{
		ID:         in.Id,             //
		CategoryID: in.CategoryId,     // 分类ID
		Name:       in.Name,           // 规格名称
		Sort:       in.Sort,           // 排序
		Status:     in.Status,         // 状态：0->禁用；1->启用
		CreateBy:   detail.CreateBy,   // 创建人ID
		CreateTime: detail.CreateTime, // 创建时间
		UpdateBy:   &in.UpdateBy,      // 更新人ID
		UpdateTime: &now,              // 更新时间
	}

	// 2.商品规格存在时,则直接更新商品规格
	err = l.svcCtx.DB.Model(&model.PmsProductSpec{}).WithContext(l.ctx).Where(spec.ID.Eq(in.Id)).Save(item).Error

	if err != nil {
		logc.Errorf(l.ctx, "更新商品规格失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("更新商品规格失败")
	}

	return &pmsclient.UpdateProductSpecResp{}, nil
}
