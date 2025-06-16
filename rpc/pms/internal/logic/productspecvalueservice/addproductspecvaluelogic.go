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
)

// AddProductSpecValueLogic 添加商品规格值
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:37
*/
type AddProductSpecValueLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddProductSpecValueLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductSpecValueLogic {
	return &AddProductSpecValueLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddProductSpecValue 添加商品规格值
func (l *AddProductSpecValueLogic) AddProductSpecValue(in *pmsclient.AddProductSpecValueReq) (*pmsclient.AddProductSpecValueResp, error) {
	q := query.PmsProductSpecValue

	item := &model.PmsProductSpecValue{
		SpecID:   in.SpecId,   // 规格ID
		Value:    in.Value,    // 规格值
		Sort:     in.Sort,     // 排序
		Status:   in.Status,   // 状态：0->禁用；1->启用
		CreateBy: in.CreateBy, // 创建人ID
	}

	err := q.WithContext(l.ctx).Create(item)
	if err != nil {
		logc.Errorf(l.ctx, "添加商品规格值失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("添加商品规格值失败")
	}

	return &pmsclient.AddProductSpecValueResp{}, nil
}
