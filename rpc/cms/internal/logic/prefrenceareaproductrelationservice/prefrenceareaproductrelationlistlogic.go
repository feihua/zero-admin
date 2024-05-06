package prefrenceareaproductrelationservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// PrefrenceAreaProductRelationListLogic
/*
Author: LiuFeiHua
Date: 2024/5/6 10:22
*/
type PrefrenceAreaProductRelationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPrefrenceAreaProductRelationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PrefrenceAreaProductRelationListLogic {
	return &PrefrenceAreaProductRelationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// PrefrenceAreaProductRelationList 优选商品关联查询
func (l *PrefrenceAreaProductRelationListLogic) PrefrenceAreaProductRelationList(in *cmsclient.PrefrenceAreaProductRelationListReq) (*cmsclient.PrefrenceAreaProductRelationListResp, error) {

	var ids []int64
	productRelation := query.CmsPrefrenceAreaProductRelation
	err := productRelation.WithContext(l.ctx).Select(productRelation.PrefrenceAreaID).Where(productRelation.ProductID.Eq(in.ProductId)).Scan(&ids)

	if err != nil {
		logc.Errorf(l.ctx, "查询关联优选列表信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, err
	}

	return &cmsclient.PrefrenceAreaProductRelationListResp{
		PrefrenceAreaId: ids,
	}, nil
}
