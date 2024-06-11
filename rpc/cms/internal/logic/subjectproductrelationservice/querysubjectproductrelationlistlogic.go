package subjectproductrelationservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// QuerySubjectProductRelationListLogic 查询专题商品关系表列表
/*
Author: LiuFeiHua
Date: 2024/6/11 16:42
*/
type QuerySubjectProductRelationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQuerySubjectProductRelationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySubjectProductRelationListLogic {
	return &QuerySubjectProductRelationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QuerySubjectProductRelationList 查询专题商品关系表列表
func (l *QuerySubjectProductRelationListLogic) QuerySubjectProductRelationList(in *cmsclient.QuerySubjectProductRelationListReq) (*cmsclient.QuerySubjectProductRelationListResp, error) {
	var ids []int64
	q := query.CmsSubjectProductRelation
	err := q.WithContext(l.ctx).Select(q.SubjectID).Where(q.ProductID.Eq(in.ProductId)).Scan(&ids)

	if err != nil {
		logc.Errorf(l.ctx, "查询关联专题列表信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, err
	}

	return &cmsclient.QuerySubjectProductRelationListResp{
		SubjectIds: ids,
	}, nil
}
