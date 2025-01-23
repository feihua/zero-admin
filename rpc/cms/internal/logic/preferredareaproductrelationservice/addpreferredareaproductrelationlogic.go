package preferredareaproductrelationservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/cms/gen/model"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddPreferredAreaProductRelationLogic 添加优选专区和产品关系表
/*
Author: LiuFeiHua
Date: 2024/6/11 16:40
*/
type AddPreferredAreaProductRelationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPreferredAreaProductRelationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPreferredAreaProductRelationLogic {
	return &AddPreferredAreaProductRelationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddPreferredAreaProductRelation 添加优选专区和产品关系表
func (l *AddPreferredAreaProductRelationLogic) AddPreferredAreaProductRelation(in *cmsclient.AddPreferredAreaProductRelationReq) (*cmsclient.AddPreferredAreaProductRelationResp, error) {
	// 1.先删除优选商品的关联
	q := query.CmsPreferredAreaProductRelation
	_, err := q.WithContext(l.ctx).Where(q.ProductID.Eq(in.ProductId)).Delete()
	if err != nil {
		return nil, err
	}

	// 2.重新添加优选商品的关联
	productRelations := make([]*model.CmsPreferredAreaProductRelation, 0)
	for _, id := range in.PreferredAreaId {
		productRelations = append(productRelations, &model.CmsPreferredAreaProductRelation{
			PreferredAreaID: id,
			ProductID:       in.ProductId,
		})
	}

	err = q.WithContext(l.ctx).CreateInBatches(productRelations, len(productRelations))
	if err != nil {
		return nil, err
	}

	return &cmsclient.AddPreferredAreaProductRelationResp{}, nil
}
