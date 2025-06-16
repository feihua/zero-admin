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
)

// AddProductSpecLogic 添加商品规格
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:37
*/
type AddProductSpecLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddProductSpecLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductSpecLogic {
	return &AddProductSpecLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddProductSpec 添加商品规格
func (l *AddProductSpecLogic) AddProductSpec(in *pmsclient.AddProductSpecReq) (*pmsclient.AddProductSpecResp, error) {
	q := query.PmsProductSpec

	item := &model.PmsProductSpec{
		CategoryID: in.CategoryId, // 分类ID
		Name:       in.Name,       // 规格名称
		Sort:       in.Sort,       // 排序
		Status:     in.Status,     // 状态：0->禁用；1->启用
		CreateBy:   in.CreateBy,   // 创建人ID
	}

	err := q.WithContext(l.ctx).Create(item)
	if err != nil {
		logc.Errorf(l.ctx, "添加商品规格失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("添加商品规格失败")
	}

	return &pmsclient.AddProductSpecResp{}, nil
}
