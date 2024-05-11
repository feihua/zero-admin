package preferredareaproductrelationservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PreferredAreaProductRelationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPreferredAreaProductRelationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PreferredAreaProductRelationListLogic {
	return &PreferredAreaProductRelationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PreferredAreaProductRelationListLogic) PreferredAreaProductRelationList(in *cmsclient.PreferredAreaProductRelationListReq) (*cmsclient.PreferredAreaProductRelationListResp, error) {
	var ids []int64
	productRelation := query.CmsPreferredAreaProductRelation
	err := productRelation.WithContext(l.ctx).Select(productRelation.PreferredAreaID).Where(productRelation.ProductID.Eq(in.ProductId)).Scan(&ids)

	if err != nil {
		logc.Errorf(l.ctx, "查询关联优选列表信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, err
	}

	return &cmsclient.PreferredAreaProductRelationListResp{
		PreferredAreaId: ids,
	}, nil
}
