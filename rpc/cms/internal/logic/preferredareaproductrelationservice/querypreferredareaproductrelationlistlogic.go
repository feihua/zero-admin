package preferredareaproductrelationservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryPreferredAreaProductRelationListLogic 查询优选专区和产品关系表列表
/*
Author: LiuFeiHua
Date: 2024/6/11 16:40
*/
type QueryPreferredAreaProductRelationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryPreferredAreaProductRelationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryPreferredAreaProductRelationListLogic {
	return &QueryPreferredAreaProductRelationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryPreferredAreaProductRelationList 查询优选专区和产品关系表列表
func (l *QueryPreferredAreaProductRelationListLogic) QueryPreferredAreaProductRelationList(in *cmsclient.QueryPreferredAreaProductRelationListReq) (*cmsclient.QueryPreferredAreaProductRelationListResp, error) {
	var ids []int64
	q := query.CmsPreferredAreaProductRelation
	err := q.WithContext(l.ctx).Select(q.PreferredAreaID).Where(q.ProductID.Eq(in.ProductId)).Scan(&ids)

	if err != nil {
		logc.Errorf(l.ctx, "查询关联优选列表信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, err
	}

	return &cmsclient.QueryPreferredAreaProductRelationListResp{
		PreferredAreaIds: ids,
	}, nil

}
