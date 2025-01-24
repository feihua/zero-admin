package memberproductcategoryrelationservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberProductCategoryRelationListLogic 查询会员与产品分类关系表（用户喜欢的分类）列表
/*
Author: LiuFeiHua
Date: 2024/6/11 14:11
*/
type QueryMemberProductCategoryRelationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberProductCategoryRelationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberProductCategoryRelationListLogic {
	return &QueryMemberProductCategoryRelationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryMemberProductCategoryRelationList 查询会员与产品分类关系表（用户喜欢的分类）列表
func (l *QueryMemberProductCategoryRelationListLogic) QueryMemberProductCategoryRelationList(in *umsclient.QueryMemberProductCategoryRelationListReq) (*umsclient.QueryMemberProductCategoryRelationListResp, error) {
	q := query.UmsMemberProductCategoryRelation
	result, err := q.WithContext(l.ctx).Where(q.MemberID.Eq(in.MemberId)).Find()

	if err != nil {
		logc.Errorf(l.ctx, "查询会员与产品关糸列表信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*umsclient.MemberProductCategoryRelationListData
	for _, item := range result {

		list = append(list, &umsclient.MemberProductCategoryRelationListData{
			Id:                item.ID,                //
			MemberId:          item.MemberID,          // 会员id
			ProductCategoryId: item.ProductCategoryID, // 商品分类id
		})
	}

	return &umsclient.QueryMemberProductCategoryRelationListResp{
		Total: 0,
		List:  list,
	}, nil

}
