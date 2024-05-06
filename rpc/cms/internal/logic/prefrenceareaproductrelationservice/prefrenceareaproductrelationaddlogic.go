package prefrenceareaproductrelationservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/gen/model"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// PrefrenceAreaProductRelationAddLogic
/*
Author: LiuFeiHua
Date: 2024/5/6 10:00
*/
type PrefrenceAreaProductRelationAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPrefrenceAreaProductRelationAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PrefrenceAreaProductRelationAddLogic {
	return &PrefrenceAreaProductRelationAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// PrefrenceAreaProductRelationAdd 优选商品关联
func (l *PrefrenceAreaProductRelationAddLogic) PrefrenceAreaProductRelationAdd(in *cmsclient.PrefrenceAreaProductRelationAddReq) (*cmsclient.PrefrenceAreaProductRelationAddResp, error) {
	//1.先删除优选商品的关联
	productRelation := query.CmsPrefrenceAreaProductRelation
	_, err := productRelation.WithContext(l.ctx).Where(productRelation.ProductID.Eq(in.ProductId)).Delete()
	if err != nil {
		return nil, err
	}

	//2.重新添加优选商品的关联
	productRelations := make([]*model.CmsPrefrenceAreaProductRelation, 0)
	for _, id := range in.PrefrenceAreaId {
		productRelations = append(productRelations, &model.CmsPrefrenceAreaProductRelation{
			PrefrenceAreaID: id,
			ProductID:       in.ProductId,
		})
	}

	err = productRelation.WithContext(l.ctx).CreateInBatches(productRelations, len(productRelations))
	if err != nil {
		return nil, err
	}

	return &cmsclient.PrefrenceAreaProductRelationAddResp{}, nil
}
